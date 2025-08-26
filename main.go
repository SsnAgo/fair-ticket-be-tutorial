package main

import (
	"fair-ticket-be-tutorial/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Router(r)
	r.Run(":8080")
}
