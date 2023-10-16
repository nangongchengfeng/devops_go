package main

import (
	"devops_go/config"
	"devops_go/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller.Router.InitApiRouter(r)
	r.Run(config.ListenAddr)
}
