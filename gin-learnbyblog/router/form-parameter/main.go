package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 将web目录中的文件映射到路由中
	// router.Static("/web", "./web") 将项目目录中的 web 目录映射为 /web URL路径。 这意味着当用户访问http://localhost:8080/web/index.html时，Gin框架将在项目目录的 web 目录中查找 index.html 文件并返回至客户端浏览器
	// 注意，第二个参数 "./web" 是web目录的实际文件系统路径
	r.StaticFS("/", http.Dir("./web"))

	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})
	r.Run()
}
