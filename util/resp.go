// Package util 请求响应
package util

import (
	"encoding/json"
	"log"
	"net/http"
)

// H 响应结构体
type H struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// RespFail 失败的响应
func RespFail(w http.ResponseWriter, msg string) {
	resp(w, -1, msg, nil)
}

// RespSuccess 成功的响应
func RespSuccess(w http.ResponseWriter, data interface{}, msg string) {
	resp(w, 0, msg, data)
}

// 响应函数
func resp(w http.ResponseWriter, code int, msg string, data interface{}) {
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
