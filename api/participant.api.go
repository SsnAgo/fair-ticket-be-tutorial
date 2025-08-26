package api

import (
	"errors"
	"fair-ticket-be-tutorial/service"

	"github.com/gin-gonic/gin"
)

type ParticipantApi struct {
	participantService *service.ParticipantService
	projectService     *service.ProjectService
}

func NewParticipantApi() *ParticipantApi {
	return &ParticipantApi{
		participantService: service.NewParticipantService(),
		projectService:     service.NewProjectService(),
	}
}

func (api *ParticipantApi) Participate(c *gin.Context) {
	var req ParticipateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ApiFailed(c, err)
		return
	}
	participant, err := api.participantService.GetOne(req.ProjectOnchainID, req.Address)
	if err != nil {
		ApiFailed(c, err)
		return
	}
	if participant != nil {
		ApiFailed(c, errors.New("already participated"))
		return
	}
	tx, err := api.participantService.Participate(req.Name, req.Address, req.LuckyNum, req.ProjectID, req.ProjectOnchainID)
	if err != nil {
		ApiFailed(c, err)
		return
	}
	ApiSuccess(c, ParticipateResponse{
		TxHash: tx.Hash().Hex(),
	})
}

func (api *ParticipantApi) ParticipateInfo(c *gin.Context) {
	var req ParticipateInfoRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		ApiFailed(c, err)
		return
	}
	// 获取某个参与信息
	participant, err := api.participantService.GetOne(req.ProjectOnchainID, req.Address)
	if err != nil {
		ApiFailed(c, err)
		return
	}
	// 获取对应项目信息
	project, err := api.projectService.GetOneByOnchainID(req.ProjectOnchainID)
	if err != nil {
		ApiFailed(c, err)
		return
	}
	resp := ParticipateInfoResponse{
		Participant: participant,
		Project:     project,
	}
	ApiSuccess(c, resp)
}

func (api *ParticipantApi) ParticipatedProjects(c *gin.Context) {
	var req ParticipatedProjectsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		ApiFailed(c, err)
		return
	}
	participated, err := api.participantService.GetAllByAddress(req.Address)
	if err != nil {
		ApiFailed(c, err)
		return
	}
	var resp []ParticipateInfoResponse
	// 获取项目信息
	for _, participate := range participated {
		project, err := api.projectService.GetOneByOnchainID(participate.ProjectOnchainID)
		if err != nil {
			ApiFailed(c, err)
			return
		}
		resp = append(resp, ParticipateInfoResponse{
			Project:     project,
			Participant: participate,
		})
	}
	ApiSuccess(c, resp)
}

func (api *ParticipantApi) VerifyMerkleProof(c *gin.Context) {
	var req VerifyMerkleProofRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ApiFailed(c, err)
		return
	}
	pass, err := api.participantService.VerifyMerkleProof(req.ProjectOnchainID, req.Address, req.Proof)
	if err != nil {
		ApiFailed(c, err)
		return
	}
	ApiSuccess(c, pass)
}
