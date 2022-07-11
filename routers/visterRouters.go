package routers

import (
	"go_0/function"

	"github.com/gin-gonic/gin"
)

func VisterRoutersInit(e *gin.Engine) {
	visterRouters := e.Group("/vister")
	visterRouters.Use(function.Counttime)
	{
		visterRouters.GET("/", function.Login)
		visterRouters.POST("/dologin", function.DoLogin) //post请求交给DoLogin处理
	}
}
