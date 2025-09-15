package service

import (
	"errors"
	"fair-ticket-be-tutorial/bindings"
	"fair-ticket-be-tutorial/contract"
	"fair-ticket-be-tutorial/cryptoo"
	"fair-ticket-be-tutorial/db"
	"fair-ticket-be-tutorial/model"
	"log"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gorm.io/gorm"
)

var (
	ERR_USER_INCONSISTENCY = errors.New("user inconsistency")
)

type EventService struct {
	db                 *gorm.DB
	contract           *contract.Contract
	participantService *ParticipantService
}

func NewEventService() *EventService {
	return &EventService{
		db:                 db.DB(),
		contract:           contract.NewFairTicketContract(),
		participantService: NewParticipantService(),
	}
}

// 监听智能合约发出的项目创建事件
func (e *EventService) HandleProjectCreated() {
	event, sink, err := e.contract.WatchProjectCreated()
	if err != nil {
		log.Println("watch project created error", err)
		return
	}
	defer func() {
		event.Unsubscribe()
		close(sink)
	}()
	for {
		select {
		case event := <-sink:
			pid := event.ProjectId.Int64()
			fp := common.Bytes2Hex(event.Fingerprint[:])
			// 去数据库中更新 onchain_id 这样链上和链下就可通过 fingerprint和 onchain_id关联起来了
			if err := e.db.Model(&model.Project{}).Where("fingerprint = ?", fp).Update("onchain_id", pid).Error; err != nil {
				log.Println("update project onchain_id error", err)
				return
			}
		case err := <-event.Err():
			log.Println("watch project created error", err)
			return
		}
	}
}

// 监听智能合约发出的项目开始事件
func (e *EventService) HandleProjectStarted() {
	event, sink, err := e.contract.WatchProjectStarted()
	if err != nil {
		log.Println("watch project started error", err)
		return
	}
	defer func() {
		event.Unsubscribe()
		close(sink)
	}()

	for {
		select {
		case event := <-sink:
			pid := event.ProjectId.Int64()
			// 去数据库中更新状态
			if err := e.db.Model(&model.Project{}).Where("onchain_id = ?", pid).Update("status", PROJECT_STATUS_INPROGRESS).Error; err != nil {
				log.Println("update project status error", err)
				return
			}
		case err := <-event.Err():
			log.Println("watch project started error", err)
			return
		}
	}
}

// 监听智能合约发出的项目结束事件
func (e *EventService) HandleProjectFinished() {
	event, sink, err := e.contract.WatchProjectFinished()
	if err != nil {
		log.Println("watch project finished error", err)
		return
	}
	defer func() {
		event.Unsubscribe()
		close(sink)
	}()

	for {
		select {
		case event := <-sink:
			pid := event.ProjectId.Int64()
			// 去数据库中更新状态
			if err := e.db.Model(&model.Project{}).Where("onchain_id = ?", pid).Update("status", PROJECT_STATUS_FINISHED).Error; err != nil {
				log.Println("update project status error", err)
				return
			}
		case err := <-event.Err():
			log.Println("watch project finished error", err)
			return
		}
	}
}

// 监听智能合约发出的幸运数发布事件
func (e *EventService) HandleMagicNumberPublished() {
	event, sink, err := e.contract.WatchMagicNumberPublished()
	if err != nil {
		log.Println("watch magic number published error", err)
		return
	}
	defer func() {
		event.Unsubscribe()
		close(sink)
	}()

	for {
		select {
		case event := <-sink:
			pid := event.ProjectId.Int64()
			magicNum := event.MagicNumber.Int64()

			// 进行参与者数据比对
			var participantsFromDB []*model.Participant
			if err := e.db.Model(&model.Participant{}).Where("project_onchain_id = ?", pid).Find(&participantsFromDB).Error; err != nil {
				log.Println("get participants from db error", err)
				return
			}
			participantsFromContract, err := e.contract.GetParticipantsFromContract(pid)
			if err != nil {
				log.Println("get participants from contract error", err)
				return
			}
			// 比对
			if err := e.userConsistencyCheck(participantsFromDB, participantsFromContract); err != nil {
				log.Println("user consistency check error", err)
				return
			}

			// 获取项目信息
			var project model.Project
			if err := e.db.Model(&model.Project{}).Where("onchain_id = ?", pid).First(&project).Error; err != nil {
				log.Println("get project error", err)
				return
			}

			// 计算中奖用户，构建merkle tree
			winners := e.calculateWinners(participantsFromContract, magicNum, int64(project.TotalSupply))
			treeRoot, proofs := e.merkleRootAndProofs(winners, magicNum)
			// 去数据库中更新 lottery_num 和 merkle root
			if err := e.db.Model(&model.Project{}).Where("onchain_id = ?", pid).Updates(map[string]interface{}{
				"lottery_num": magicNum,
				"merkle_root": common.Bytes2Hex(treeRoot),
			}).Error; err != nil {
				log.Println("update project lottery_num and merkle_root error", err)
				return
			}
			// 去链上更新项目的 MerkleRoot
			if _, err := e.contract.SetMerkleRoot(pid, treeRoot); err != nil {
				log.Println("update project merkle root error", err)
				return
			}
			// 更新参与者中奖信息
			for _, winner := range winners {
				winnerProofs := proofs[winner]
				proofStr := make([]string, len(winnerProofs))
				for i, proof := range winnerProofs {
					proofStr[i] = common.Bytes2Hex(proof)
				}
				if err := e.participantService.MarkWin(project.OnchainID, winner.Hex(), proofStr); err != nil {
					log.Println("mark win error", err)
				}
			}
		case err := <-event.Err():
			log.Println("watch project finished error", err)
			return
		}
	}
}

func (e *EventService) userConsistencyCheck(participantsFromDB []*model.Participant, participantsFromContract []bindings.Participant) (err error) {
	if len(participantsFromDB) != len(participantsFromContract) {
		return ERR_USER_INCONSISTENCY
	}
	participantsFromDBMap := make(map[string]bool)
	for _, participant := range participantsFromDB {
		participantsFromDBMap[participant.Address] = true
	}
	for _, participant := range participantsFromContract {
		if _, ok := participantsFromDBMap[participant.Addr.Hex()]; !ok {
			return ERR_USER_INCONSISTENCY
		}
	}
	return nil
}

// 计算中奖名单
// todo 这里其实可优化成堆排序构造一个x个元素的堆 数据可流式进入  这样就可以大大减少数据量很大时 内存和耗时高的问题
func (e *EventService) calculateWinners(participantsFromContract []bindings.Participant, magicNum int64, totalSupply int64) (winners []common.Address) {
	// 如果总供应量能满足所有参与者，则所有参与者中奖
	if totalSupply >= int64(len(participantsFromContract)) {
		for _, p := range participantsFromContract {
			winners = append(winners, p.Addr)
		}
		return winners
	}
	// 临时用一下这个结构体 方便进行临时存储和排序
	type ParticipantLottery struct {
		Addr common.Address
		Hash common.Hash
	}
	// 将用户地址 幸运数 和 magic number 拼接起来，计算hash，然后排序
	participantsLottery := make([]ParticipantLottery, len(participantsFromContract))
	for _, p := range participantsFromContract {
		uaddr := p.Addr.Bytes()
		ulucky := p.LuckyNum.Bytes()
		magic := big.NewInt(magicNum).Bytes()
		hash := crypto.Keccak256Hash(uaddr, ulucky, magic)
		participantsLottery = append(participantsLottery, ParticipantLottery{
			Addr: p.Addr,
			Hash: hash,
		})
	}
	// 根据hash从大到小排序
	sort.Slice(participantsLottery, func(i, j int) bool {
		return participantsLottery[i].Hash.Big().Cmp(participantsLottery[j].Hash.Big()) > 0
	})
	// 前 totalSupply 个用户中奖
	for i := 0; i < int(totalSupply); i++ {
		winners = append(winners, participantsLottery[i].Addr)
	}
	return winners
}

func (e *EventService) merkleRootAndProofs(addresses []common.Address, magicNum int64) (treeRoot []byte, proofs map[common.Address][][]byte) {

	return cryptoo.BuildMerkleTree(addresses)
}
