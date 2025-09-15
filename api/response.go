package api

import (
	"fair-ticket-be-tutorial/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
封装Api返回结果的函数
*/

// 成功和失败的code常量
const (
	SUCCESS = 0
	FAILED  = 7
)

// 统一返回结果
type Result struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

// 成功返回结果
func ApiSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Result{
		Code: SUCCESS,
		Data: data,
	})
}

// 失败返回结果
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
