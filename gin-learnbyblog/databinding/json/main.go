package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	// binding:"required" 修饰的字段，若接收为空值则报错，是必须字段
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	// 创建路由
	r := gin.Default()
	// 默认使用了两个中间件Logger(), Recovery()
	// JSON绑定
	r.POST("loginJSON", func(c *gin.Context) {
		// 声明接收的变量
		var json Login
		// 将reque的body中的数据，自动按照json格式解析到结构体
		if err := c.ShouldBindJSON(&json); err != nil {
			// 返回错误信息
			// gin.H 封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 判断用户名密码是否正确
		if json.User != "root1" || json.Password != "admin1" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run(":8080")
}

// curl localhost:8080/loginJSON -H 'content-type:application/json' -d '{"user":"root1","password":"admin1"}' -X POST
