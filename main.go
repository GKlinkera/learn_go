package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type At struct {
	First string
	Txt   string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "gogogo",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "shouye.html", gin.H{
			"title":   "首页",
			"context": "hello world",
		})
	})

	r.GET("/news", func(c *gin.Context) {
		ax := &At{
			First: "1",
			Txt:   "12312",
		}
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "第二页",
			"news":  ax,
		})
	})

	r.POST("/post", func(c *gin.Context) {
		c.String(200, "post return")
	})

	r.Run() //2131313
}
