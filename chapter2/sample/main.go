package main

import (
	"log"
	"os"

	"chapter2/sample/search"
)

// init在main之前调用
func init() {
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

// main是整个程序的入口
func main() {
	// 使用特定的项做搜索
	search.Run("president")
}
