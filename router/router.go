package router

import (
	"fair-ticket-be-tutorial/api"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	// 在gin中使用cors中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	projectApi := api.NewProjectApi()
	participantApi := api.NewParticipantApi()
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
