package main

import (
	"devops_go/config"
	"devops_go/controller"
	"devops_go/service"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
)

func main() {
	// 初始化 k8s
	service.K8s.Init()

	// Create a router instance using the gin package
	r := gin.Default()
	// Initialize the router with the controller's router
	controller.Router.InitApiRouter(r)
	// Run the router on the specified address
	logger.Info("服务启动成功")
	r.Run(config.ListenAddr)

}
