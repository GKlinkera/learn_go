package function

import (
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

func Shouye(c *gin.Context) {
	tq := &Tianqi{
		Rain:  "下雨",
		Cloud: "多云",
	}
	c.String(200, "首页")
	c.JSON(http.StatusOK, tq)
}

func Times(c *gin.Context) {
	ti := &Timing{
		Year: "2022年",
	}
	tt := time.Now()
	c.String(200, "时间", tt)

	c.JSON(http.StatusOK, ti)
}

func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func DoLogin(c *gin.Context) {

	u := &Infouser{}
	if err := c.ShouldBind(u); err == nil {
		c.HTML(http.StatusOK, "dologin.html", u)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}

}

func Upload(c *gin.Context) {

	file, err := c.FormFile("picture")
	dst := path.Join("./static/upload", file.Filename)
	if err == nil {
		c.SaveUploadedFile(file, dst)
		c.HTML(http.StatusOK, "upload.html", nil)
	}

}
