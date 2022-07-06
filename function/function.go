package function

import (
	"net/http"

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

func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func DoLogin(c *gin.Context) {
	u := &Infouser{}
	if err := c.ShouldBind(u); err == nil {
		c.JSON(http.StatusOK, u)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}

}
