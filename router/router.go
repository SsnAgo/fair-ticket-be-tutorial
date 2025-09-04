package router

import (
	"fair-ticket-be-tutorial/api"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	// 在gin中使用cors中间件
	r.Use(api.CorsMiddleware)

	// 初始化api
	projectApi := api.NewProjectApi()
	participantApi := api.NewParticipantApi()

	// 初始化路由组
	projectGroup := r.Group("/project")
	{
		projectGroup.POST("/create", projectApi.CreateProject)
		projectGroup.PUT("/start", projectApi.StartProject)
		projectGroup.PUT("/finish", projectApi.FinishProject)
		projectGroup.PUT("/lottery", projectApi.LotteryProject)
		projectGroup.GET("/info", projectApi.GetProjectInfo)
		projectGroup.GET("/list", projectApi.ListAllProjects)
	}
	participantGroup := r.Group("/participant")
	{
		participantGroup.POST("/participate", participantApi.Participate)
		participantGroup.GET("/info", participantApi.ParticipateInfo)
		participantGroup.GET("/list", participantApi.ParticipatedProjects)
		participantGroup.POST("/verify", participantApi.VerifyMerkleProof)
	}
}
