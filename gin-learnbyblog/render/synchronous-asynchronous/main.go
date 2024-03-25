package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 1.异步
	r.GET("/long_async", func(c *gin.Context) {
		// 创建一个副本
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行:" + copyContext.Request.URL.Path)
		}()
	})
	// 同步
	r.GET("/login_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行:" + c.Request.URL.Path)
	})

	r.Run(":8080")
}
