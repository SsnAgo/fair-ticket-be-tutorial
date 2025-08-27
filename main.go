package main

import (
	"fair-ticket-be-tutorial/listener"
	"fair-ticket-be-tutorial/router"

	"github.com/gin-gonic/gin"
)

func main() {
	listener.Start()
	r := gin.Default()
	router.Router(r)
	r.Run(":8080")
}
