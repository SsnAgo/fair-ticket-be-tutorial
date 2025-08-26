package service

import (
	"errors"
	"fair-ticket-be-tutorial/contract"
	"fair-ticket-be-tutorial/db"
	"fair-ticket-be-tutorial/model"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
)

type ParticipantService struct {
	db       *gorm.DB
	contract *contract.Contract
}

func NewParticipantService() *ParticipantService {
	return &ParticipantService{db: db.DB(), contract: contract.NewFairTicketContract()}
}

// 参与者参与项目 往数据库里面新增信息 并调用合约
func (s *ParticipantService) Participate(name string, address string, luckyNum int64, projectID int64, projectOnchainID int64) (tx *types.Transaction, err error) {
	if err := s.db.Create(&model.Participant{
		Name:             name,
		Address:          address,
		LuckyNum:         luckyNum,
		ProjectID:        projectID,
		ProjectOnchainID: projectOnchainID,
	}).Error; err != nil {
		return nil, err
	}
	tx, err = s.contract.Participate(projectOnchainID, address, luckyNum)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// 更新中奖的参与者，设置中奖状态和merkle proof
func (s *ParticipantService) MarkWin(projectOnchainId int64, address string, proof []string) error {
	proofStr := strings.Join(proof, ",")
	return s.db.Model(&model.Participant{}).Where("project_onchain_id = ? AND address = ?", projectOnchainId, address).Updates(map[string]interface{}{
		"win":          true,
		"merkle_proof": proofStr,
	}).Error
}

// 通过projectOnchainId和address确定一个参与者
func (s *ParticipantService) GetOne(projectOnchainId int64, address string) (*model.Participant, error) {
	var participant model.Participant
	err := s.db.Model(&model.Participant{}).Where("project_onchain_id = ? AND address = ?", projectOnchainId, address).First(&participant).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &participant, nil
}

// 获取项目所有参与者
func (s *ParticipantService) GetAll(projectOnchainId int64) ([]*model.Participant, error) {
	var participants []*model.Participant
	err := s.db.Model(&model.Participant{}).Where("project_onchain_id = ?", projectOnchainId).Find(&participants).Error
	if err != nil {
		return nil, err
	}
	return participants, nil
}

// 调用合约验证merkle proof
func (s *ParticipantService) VerifyMerkleProof(projectOnchainId int64, address string, proof []string) (bool, error) {
	return s.contract.VerifyMerkleProof(projectOnchainId, address, proof)
}

// 获取某个地址的参与记录
func (s *ParticipantService) GetAllByAddress(address string) ([]*model.Participant, error) {
	var participants []*model.Participant
	err := s.db.Model(&model.Participant{}).Where("address = ?", address).Find(&participants).Error
	if err != nil {
		return nil, err
	}
	return participants, nil
}
