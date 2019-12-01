package main

import (
	"github.com/RevelesD/GoBasicGA/server"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := server.SetupRouter()
	r.Run(":8080")
}