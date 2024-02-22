// Package main IM 入口
package main

import (
	"net/http"
)

func main() {
	// 启动 HTTP 服务
	http.ListenAndServe(":8080", nil)
}
