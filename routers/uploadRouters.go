package routers

import (
	"go_0/function"

	"github.com/gin-gonic/gin"
)

func UploadRoutersInit(e *gin.Engine) {
	uploadRouters := e.Group("/upload")
	uploadRouters.Use(function.Counttime)
	{
		uploadRouters.POST("/", function.Upload)
	}
}
