package main

import (
	"fmt"
	"net/http"

	"struct/dbops"

	"github.com/gin-gonic/gin"
)

// GetBook...
func GetBook(c *gin.Context) {
	books, err := dbops.GetAllbook()
	if err != nil {
		fmt.Println("获取参数失败", err)
	}
	for i, ele := range books {
		fmt.Println("comment: %d, %v \n", i, ele)
	}
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "i am test", "ce": books})
}
