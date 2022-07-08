package initial

import (
	"go_0/routers"

	"github.com/gin-gonic/gin"
)

func Initialall() {
	e := gin.Default()
	e.LoadHTMLGlob("views/*")
	routers.VisterRoutersInit(e)  //http://localhost:8001/vister/
	routers.DefaultRoutersInit(e) //http://localhost:8001/   http://localhost:8001/times
	routers.UploadRoutersInit(e)
	e.Run(":8001")
}
