package initial

import (
	"fmt"
	"go_0/routers"

	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
)

func cfgread() (a string) {

	var cfg *ini.File

	cfg, err := ini.Load("conf/app.ini") //此路径为ini文件的路径
	if err != nil {
		fmt.Println("读取ini文件失败")
	}
	section := cfg.Section("Server")         //读取名字为server的区域部分
	Port := section.Key("HttpPort").String() //选择区域内的键并指定类型
	return Port
}

func Initialall() {
	e := gin.Default()
	e.LoadHTMLGlob("views/*")
	routers.VisterRoutersInit(e)  //http://localhost:8001/vister/
	routers.DefaultRoutersInit(e) //http://localhost:8001/   http://localhost:8001/times
	routers.UploadRoutersInit(e)
	ht := cfgread()
	e.Run(ht)
}
