package main

import (
	"./server"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := server.SetupRouter()
	r.Run(":8080")
}