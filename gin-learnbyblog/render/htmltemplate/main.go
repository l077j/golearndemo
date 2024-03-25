package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("tem/**/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.html", gin.H{"title": "i am test", "address": "www.51mh.com"})
	})
	// 重定向 redirect
	// r.GET("/index", func(c *gin.Context) {
	// 	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	// })
	r.Run(":8080")
}
