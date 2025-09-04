package main

import (
	"fair-ticket-be-tutorial/listener"
	"fair-ticket-be-tutorial/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 启动监听器
	listener.Start()
	// 初始化gin
	r := gin.Default()
	// 初始化路由
	router.Router(r)
	// 启动gin
	r.Run(":8080")
}
