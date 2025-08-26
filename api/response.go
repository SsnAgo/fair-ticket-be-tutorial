package api

import (
	"fair-ticket-be-tutorial/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 0
	FAILED  = 7
)

type Result struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func ApiSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Result{
		Code: SUCCESS,
		Data: data,
	})
}

func ApiFailed(c *gin.Context, err error) {
	c.JSON(http.StatusOK, Result{
		Code:    FAILED,
		Data:    nil,
		Message: err.Error(),
	})
}

/*
Project
*/
type CreateProjectResponse struct {
	TxHash string `json:"txHash"`
}

type ProjectChangeStatusResponse struct {
	TxHash string `json:"txHash"`
}

/*
Participant
*/
type ParticipateResponse struct {
	TxHash string `json:"txHash"`
}

type ParticipateInfoResponse struct {
	Project     *model.Project     `json:"project"`
	Participant *model.Participant `json:"participant"`
}
