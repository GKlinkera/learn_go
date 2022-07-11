package routers

import (
	"go_0/function"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(e *gin.Engine) {
	defaultrRouters := e.Group("/")

	{

		defaultrRouters.GET("/", function.Counttime, function.Shouye)

		aaRouters := defaultrRouters.Group("/times") //路由嵌套
		{
			aaRouters.GET("/", function.Counttime, function.Times)
		}

	}
}
