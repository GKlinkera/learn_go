package main

//gin实现简单的登录程序
import (
	"go_0/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.LoadHTMLGlob("views/*")
	routers.VisterRoutersInit(e)  //http://localhost:8001/vister/
	routers.DefaultRoutersInit(e) //http://localhost:8001/   http://localhost:8001/times
	e.Run(":8001")
}
