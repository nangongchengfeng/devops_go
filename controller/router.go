package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
)

// 实例化router结构体，可使用该对象点出首字母大写的方法（包外调用）
var Router router

type router struct{}

// This function initializes the router for the API
func (r *router) InitApiRouter(router *gin.Engine) {
	// Create a GET request for the "/testapi" route
	router.GET("/testapi", func(ctx *gin.Context) {
		// Return a JSON response with a status code of 200 and the message "testapi success!"
		logger.Info("testapi success!", router)
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "testapi success!",
			"data": nil,
		})
	})
}
