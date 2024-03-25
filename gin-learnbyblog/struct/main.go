package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// 使用 LoadHTMLGlob 函数加载目录中所有的 .html 文件
	r.LoadHTMLGlob("./tem/*.html")
	r.GET("/", GetBook)
	r.Run()
}
