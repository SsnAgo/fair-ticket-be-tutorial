package contract

import (
	"fair-ticket-be-tutorial/bindings"
	"fair-ticket-be-tutorial/contract/cli"
	"math"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/pkg/errors"
)

const (
	// 注意每次部署完合约后，记得检查更新合约地址
	CONTRACT_ADDRESS = "0x5FbDB2315678afecb367f032d93F642f64180aa3"
)

var once sync.Once
var contractInstance *Contract // 单例模式

func NewFairTicketContract() *Contract {
	cli := cli.GetClient()
	once.Do(func() {
		ft, err := bindings.NewFairTicket(common.HexToAddress(CONTRACT_ADDRESS), cli.Client)
		if err != nil {
			panic(err)
		}
		contractInstance = &Contract{
			c:        cli,
			contract: ft,
		}
	})
	return contractInstance
}

var (
	ErrInvalidMerkleRoot = errors.New("invalid merkle root")
	ErrInvalidProjectId  = errors.New("invalid project id")
)

type Contract struct {
	c        *cli.EthClient
	contract *bindings.FairTicket
}

// 返回封装的客户端
func (f Contract) Client() *cli.EthClient {
	return f.c
}

// 返回原生的客户端
func (f Contract) NativeClient() *ethclient.Client {
	return f.c.Client
}

func (f Contract) CreateProject(fingerprint string, owner string, totalSupply int64) (tx *types.Transaction, err error) {
	opt := f.c.TransactorOpt()
	fingerprintBytes := common.FromHex(fingerprint)
	totalSupplyBig := new(big.Int).SetInt64(totalSupply)
	tx, err = f.contract.FairTicketTransactor.CreateProject(opt, [32]byte(fingerprintBytes), common.HexToAddress(owner), totalSupplyBig)
	return tx, errors.WithStack(err)
}

func (f Contract) StartProject(projectId int64) (tx *types.Transaction, err error) {
	opt := f.c.TransactorOpt()
	pid := new(big.Int).SetInt64(projectId)
	tx, err = f.contract.FairTicketTransactor.StartProject(opt, pid)
	return tx, errors.WithStack(err)
}

func (f Contract) FinishProject(projectId int64) (tx *types.Transaction, err error) {
	opt := f.c.TransactorOpt()
	pid := new(big.Int).SetInt64(projectId)
	tx, err = f.contract.FairTicketTransactor.FinishProject(opt, pid)
	return tx, errors.WithStack(err)
}

func (f Contract) Participate(projectId int64, pAddress string, luckyNum int64) (tx *types.Transaction, err error) {
	opt := f.c.TransactorOpt()
	pid := new(big.Int).SetInt64(projectId)
	addr := common.HexToAddress(pAddress)
	luckyNumBig := new(big.Int).SetInt64(luckyNum)
	tx, err = f.contract.FairTicketTransactor.Participate(opt, pid, addr, luckyNumBig)
	return tx, errors.WithStack(err)
}

func (f Contract) Lottery(projectId int64) (tx *types.Transaction, err error) {
	opt := f.c.TransactorOpt()
	pid := new(big.Int).SetInt64(projectId)
	tx, err = f.contract.FairTicketTransactor.Lottery(opt, pid)
	return tx, errors.WithStack(err)
}

func (f Contract) GetProjectInfo(projectId int64) (bindings.Project, error) {
	opt := f.c.CallOpt()
	pid := new(big.Int).SetInt64(projectId)
	project, err := f.contract.FairTicketCaller.GetProjectInfo(opt, pid)
	return project, errors.WithStack(err)
}

func (f Contract) getParticipantsPaginationFromContract(projectId int64, offset int, limit int) ([]bindings.Participant, error) {
	opt := f.c.CallOpt()
	pid := new(big.Int).SetInt64(projectId)
	offsetBig := big.NewInt(int64(offset))
	limitBig := big.NewInt(int64(limit))
	participants, err := f.contract.FairTicketCaller.GetProjectParticipants(opt, pid, offsetBig, limitBig)
	return participants, errors.WithStack(err)
}

func (f Contract) GetParticipantsFromContract(projectId int64) ([]bindings.Participant, error) {
	amount, err := f.GetProjectParticipantsAmount(projectId)
	if err != nil {
		return nil, err
	}
	offset, limit := 0, 100
	loop := int(math.Ceil(float64(amount.Int64()) / 100.0))
	var pChan = make(chan []bindings.Participant, loop)
	go func() {
		for i := 0; i < loop; i++ {
			participants, err := f.getParticipantsPaginationFromContract(projectId, offset, limit)
			if err != nil {
				return
			}
			pChan <- participants
		}
		close(pChan)
	}()

	var participants []bindings.Participant
	for p := range pChan {
		participants = append(participants, p...)
	}
	return participants, nil
}

func (f Contract) GetProjectParticipantsAmount(projectId int64) (*big.Int, error) {
	opt := f.c.CallOpt()
	pid := new(big.Int).SetInt64(projectId)
	amount, err := f.contract.FairTicketCaller.GetProjectParticipantsAmount(opt, pid)
	return amount, errors.WithStack(err)
}

func (f Contract) SetMerkleRoot(projectId int64, merkleRoot []byte) (tx *types.Transaction, err error) {
	opt := f.c.TransactorOpt()
	pid := new(big.Int).SetInt64(projectId)
	tx, err = f.contract.FairTicketTransactor.SetMerkleRoot(opt, pid, [32]byte(merkleRoot))
	return tx, errors.WithStack(err)
}

func (f Contract) VerifyMerkleProof(projectId int64, address string, proof []string) (bool, error) {
	opt := f.c.CallOpt()
	opt.From = common.HexToAddress(address)
	pid := new(big.Int).SetInt64(projectId)
	merkleProof := make([][32]byte, len(proof))
	for i, p := range proof {
		merkleProof[i] = [32]byte(common.FromHex(p))
	}
	pass, err := f.contract.FairTicketCaller.VerifyMerkleProof(opt, pid, merkleProof)
	return pass, errors.WithStack(err)
}

// 监听合约的事件，合约中定义的Event 可以通过 WatchXXX 函数进行监听
// 我们定义一个事件对应的接收通道，这里都设置为缓冲区大小为100的chan。
// 随后，当有合约中的事件触发时，会向事件通道中发送事件对象，我们通过读取通道，就可以获取这些事件
func (f Contract) WatchMagicNumberPublished() (event.Subscription, chan *bindings.FairTicketMagicNumberPublished, error) {
	sink := make(chan *bindings.FairTicketMagicNumberPublished, 100)
	sub, err := f.contract.WatchMagicNumberPublished(nil, sink, nil)
	return sub, sink, errors.WithStack(err)
}

func (f Contract) WatchProjectCreated() (event.Subscription, chan *bindings.FairTicketProjectCreated, error) {
	sink := make(chan *bindings.FairTicketProjectCreated, 100)
	sub, err := f.contract.WatchProjectCreated(nil, sink, nil, nil)
	return sub, sink, errors.WithStack(err)
}

func (f Contract) WatchProjectStarted() (event.Subscription, chan *bindings.FairTicketProjectStarted, error) {
	sink := make(chan *bindings.FairTicketProjectStarted, 100)
	sub, err := f.contract.WatchProjectStarted(nil, sink, nil)
	return sub, sink, errors.WithStack(err)
}

func (f Contract) WatchProjectFinished() (event.Subscription, chan *bindings.FairTicketProjectFinished, error) {
	sink := make(chan *bindings.FairTicketProjectFinished, 100)
	sub, err := f.contract.WatchProjectFinished(nil, sink, nil)
	return sub, sink, errors.WithStack(err)
}
