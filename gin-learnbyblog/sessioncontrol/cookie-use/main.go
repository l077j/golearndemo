package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 服务端要给客户端cookie
	r.GET("cookie", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		// 获取名为 "key_cookie" 的 HTTP cookie 的值，如果不存在则返回 "Not-Set"
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			// 给客户端设置cookie
			// maxAge int, 单位为秒
			// path,cookie所在目录
			// domain string, 域名
			// secure 是否只能通过https访问
			// httpOnly bool 是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/", "localhost", false, true)
			// 获取名为 "key_cookie" 的 HTTP cookie 的值，如果不存在则返回 "Not-Set"
		}
		fmt.Printf("cookie的值是:%s\n", cookie)
	})

	r.Run(":8080")
}
