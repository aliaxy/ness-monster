// Package main IM 入口
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func userLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认是不会解析的
	mobile := r.PostForm.Get("mobile")
	password := r.PostForm.Get("password")

	loginOk := false
	if mobile == "123456789" && password == "123456" {
		loginOk = true
	}

	if loginOk {
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		Resp(w, 0, "", data)
	} else {
		Resp(w, -1, "密码不正确", nil)
	}
}

type H struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Resp(w http.ResponseWriter, code int, msg string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	result, err := json.Marshal(h)
	if err != nil {
		log.Panicln(err.Error())
	}

	w.Write(result)
}

func main() {
	// 绑定请求和处理函数
	http.HandleFunc("/user/login", userLogin)

	// 启动 HTTP 服务
	http.ListenAndServe(":8080", nil)
}
