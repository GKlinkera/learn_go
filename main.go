package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
