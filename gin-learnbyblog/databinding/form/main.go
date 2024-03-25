package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"user" xml:"password" binding:"required`
}

func main() {
	r := gin.Default()
	r.StaticFS("/", http.Dir("./web"))
	r.POST("/loginForm", func(c *gin.Context) {
		var form Login
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if form.User != "root" || form.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run(":8080")
}
