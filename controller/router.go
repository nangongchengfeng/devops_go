package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 实例化router结构体，可使用该对象点出首字母大写的方法（包外调用）
var Router router

type router struct{}

func (r *router) InitApiRouter(router *gin.Engine) {
	router.GET("/testapi", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "testapi success!",
			"data": nil,
		})
	})
}
