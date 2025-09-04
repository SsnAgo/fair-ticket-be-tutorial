package api

import (
	"errors"
	"fair-ticket-be-tutorial/cryptoo"
	"fair-ticket-be-tutorial/service"

	"github.com/gin-gonic/gin"
)

type ProjectApi struct {
	projectService     *service.ProjectService
	participantService *service.ParticipantService
}

func NewProjectApi() *ProjectApi {
	return &ProjectApi{
		projectService:     service.NewProjectService(),
		participantService: service.NewParticipantService(),
	}
}

// 创建项目
func (api *ProjectApi) CreateProject(c *gin.Context) {
	var req ProjectCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ApiFailed(c, err)
		return
	}
	if !cryptoo.VerifySignature(req.Message, req.Signature, req.Owner) {
		ApiFailed(c, errors.New("signature verification failed"))
		return
	}
	tx, err := api.projectService.Create(req.Name, req.Description, req.ImageURL, req.TotalSupply, req.Owner)
	if err != nil {
		ApiFailed(c, err)
		return
	}
	resp := CreateProjectResponse{
		TxHash: tx.Hash().String(),
	}
	ApiSuccess(c, resp)
}

func (api *ProjectApi) StartProject(c *gin.Context) {
	var req ProjectChangeStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ApiFailed(c, err)
		return
	}
	tx, err := api.projectService.Start(req.ProjectOnchainID, req.Owner)
	if err != nil {
		ApiFailed(c, err)
		return
	}
	resp := ProjectChangeStatusResponse{
		TxHash: tx.Hash().String(),
	}
	ApiSuccess(c, resp)
}

func (api *ProjectApi) FinishProject(c *gin.Context) {
	var req ProjectChangeStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ApiFailed(c, err)
		return
	}
	tx, err := api.projectService.Finish(req.ProjectOnchainID, req.Owner)
	if err != nil {
		ApiFailed(c, err)
		return
	}
	resp := ProjectChangeStatusResponse{
		TxHash: tx.Hash().String(),
	}
	ApiSuccess(c, resp)
}

func (api *ProjectApi) LotteryProject(c *gin.Context) {
	var req ProjectChangeStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ApiFailed(c, err)
		return
	}
	tx, err := api.projectService.Lottery(req.ProjectOnchainID, req.Owner)
	if err != nil {
		ApiFailed(c, err)
		return
	}
	resp := ProjectChangeStatusResponse{
		TxHash: tx.Hash().String(),
	}
	ApiSuccess(c, resp)
}

func (api *ProjectApi) GetProjectInfo(c *gin.Context) {
	var req ProjectInfoRequest
	// 这里是bindQuery 对应request中的form tag
	if err := c.ShouldBindQuery(&req); err != nil {
		ApiFailed(c, err)
		return
	}
	project, err := api.projectService.GetOne(req.ProjectID)
	if err != nil {
		ApiFailed(c, err)
		return
	}
	ApiSuccess(c, project)
}

func (api *ProjectApi) ListAllProjects(c *gin.Context) {
	projects, err := api.projectService.GetAll()
	if err != nil {
		ApiFailed(c, err)
		return
	}
	ApiSuccess(c, projects)
}
