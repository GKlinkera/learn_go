package function

import (
	"fmt"
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

	file1, err := c.FormFile("picture")

	if err == nil {
		dst := path.Join("./static/upload", file1.Filename)
		c.SaveUploadedFile(file1, dst)

	}

	file2, err := c.FormFile("picture2")

	if err == nil {
		dst := path.Join("./static/upload", file2.Filename)
		c.SaveUploadedFile(file2, dst)
		c.HTML(http.StatusOK, "upload.html", nil)
	}

}

func Counttime(c *gin.Context) {
	start := time.Now().UnixNano()
	c.Next()
	end := time.Now().UnixNano()
	fmt.Println(end - start)

}
