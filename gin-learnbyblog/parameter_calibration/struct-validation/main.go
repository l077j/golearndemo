package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	r := gin.Default()
	r.GET("/51mh", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, fmt.Sprint(err))
			return
		}
		// %# 是 fmt 包中提供的一种格式占位符，它在格式化输出中用于输出具有 Go 语法表示方式的值。该占位符包含在格式化字符串中，通常以 %v 占位符的形式出现
		// %#v 会向标准输出流输出一个类似于 Go 语言中对该变量类型的字面常量的字符串，这样可以方便地查看和调试程序中的数据类型。对于数字类型，例如整数和浮点数，它输出时会给出具体的进位方式或指数表示。
		c.String(200, fmt.Sprintf("%#v", person))
	})

	r.Run()
}
