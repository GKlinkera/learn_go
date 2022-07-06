package routers

import (
	"go_0/function"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(e *gin.Engine) {
	defaultrRouters := e.Group("/default")
	{
		defaultrRouters.GET("/", function.Shouye)
	}
}
