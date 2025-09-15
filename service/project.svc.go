package service

import (
	"crypto/sha256"
	"errors"
	"fair-ticket-be-tutorial/contract"
	"fair-ticket-be-tutorial/db"
	"fair-ticket-be-tutorial/model"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// 设置项目状态常量
const (
	PROJECT_STATUS_NOTSTART   = 0
	PROJECT_STATUS_INPROGRESS = 1
	PROJECT_STATUS_FINISHED   = 2
)

var (
	ERR_NOT_PROJECT_OWNER = errors.New("not project owner")
)

type ProjectService struct {
	db       *gorm.DB
	contract *contract.Contract
}

func NewProjectService() *ProjectService {
	return &ProjectService{db: db.DB(), contract: contract.NewFairTicketContract()}
}

// 创建项目
func (s *ProjectService) Create(name string, description string, imageURL string, totalSupply int64, owner string) (tx *types.Transaction, err error) {
	// 根据uuid生成fingerprint fingerprint是[32]byte
	fingerprint := sha256.Sum256([]byte(uuid.New().String()))
	// 使用16进制进行解码
	fingerprintStr := common.Bytes2Hex(fingerprint[:])
	// 存储到db
	if err := s.db.Create(&model.Project{
		Name:        name,
		Description: description,
		Fingerprint: fingerprintStr,
		ImageURL:    imageURL,
		TotalSupply: int32(totalSupply),
		Status:      PROJECT_STATUS_NOTSTART,
		Owner:       owner,
	}).Error; err != nil {
		return nil, err
	}
	// 发起合约调用 去创建项目
	tx, err = s.contract.CreateProject(fingerprintStr, owner, totalSupply)
	return tx, err
}

// 开始项目
func (s *ProjectService) Start(projectOnchainID int64, owner string) (tx *types.Transaction, err error) {
	// 先查询项目是否存在
	var project model.Project
	if err := s.db.Model(&model.Project{}).Where("onchain_id = ?", projectOnchainID).First(&project).Error; err != nil {
		return nil, err
	}
	// 判断是否是项目所有者
	if project.Owner != owner {
		return nil, ERR_NOT_PROJECT_OWNER
	}
	// 更新项目状态
	if err := s.db.Model(&model.Project{}).Where("onchain_id = ?", projectOnchainID).Update("status", PROJECT_STATUS_INPROGRESS).Error; err != nil {
		return nil, err
	}
	// 发起合约调用 去开始项目
	tx, err = s.contract.StartProject(projectOnchainID)
	return tx, err
}

// 结束项目
func (s *ProjectService) Finish(projectOnchainID int64, owner string) (tx *types.Transaction, err error) {
	// 先查询项目是否存在
	var project model.Project
	if err := s.db.Model(&model.Project{}).Where("onchain_id = ?", projectOnchainID).First(&project).Error; err != nil {
		return nil, err
	}
	// 判断是否是项目所有者
	if project.Owner != owner {
		return nil, ERR_NOT_PROJECT_OWNER
	}
	// 更新项目状态
	if err := s.db.Model(&model.Project{}).Where("onchain_id = ?", projectOnchainID).Update("status", PROJECT_STATUS_FINISHED).Error; err != nil {
		return nil, err
	}
	// 发起合约调用 去结束项目
	tx, err = s.contract.FinishProject(projectOnchainID)
	return tx, err
}

// 获取单个项目
func (s *ProjectService) GetOne(id int64) (project *model.Project, err error) {
	if err := s.db.Model(&model.Project{}).Where("id = ?", id).First(&project).Error; err != nil {
		return nil, err
	}
	return project, nil
}

// 根据onchain_id获取单个项目
func (s *ProjectService) GetOneByOnchainID(onchainID int64) (project *model.Project, err error) {
	if err := s.db.Model(&model.Project{}).Where("onchain_id = ?", onchainID).First(&project).Error; err != nil {
		return nil, err
	}
	return project, nil
}

// 获取所有项目
func (s *ProjectService) GetAll() (projects []*model.Project, err error) {
	if err := s.db.Model(&model.Project{}).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

// 抽奖
func (s *ProjectService) Lottery(projectOnchainID int64, owner string) (tx *types.Transaction, err error) {
	var project model.Project
	if err := s.db.Model(&model.Project{}).Where("onchain_id = ?", projectOnchainID).First(&project).Error; err != nil {
		return nil, err
	}
	// 判断是否是项目所有者
	if project.Owner != owner {
		return nil, ERR_NOT_PROJECT_OWNER
	}
	// 发起合约调用 去发起抽奖
	tx, err = s.contract.Lottery(projectOnchainID)
	return tx, err
}
