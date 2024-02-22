// Package util 请求响应
package util

import (
	"encoding/json"
	"log"
	"net/http"
)

// H 响应结构体
type H struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Rows  interface{} `json:"rows,omitempty"`
	Total interface{} `json:"total,omitempty"`
}

// RespFail 失败的响应
func RespFail(w http.ResponseWriter, msg string) {
	resp(w, -1, msg, nil)
}

// RespSuccess 成功的响应
func RespSuccess(w http.ResponseWriter, data interface{}, msg string) {
	resp(w, 0, msg, data)
}

// RespSuccessList 成功响应的列表
func RespSuccessList(w http.ResponseWriter, lists interface{}, total interface{}) {
	// 分页数目,
	respList(w, 0, lists, total)
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

// 响应列表
func respList(w http.ResponseWriter, code int, data interface{}, total interface{}) {
	w.Header().Set("Content-Type", "application/json")
	// 设置200状态
	w.WriteHeader(http.StatusOK)
	// 输出
	// 定义一个结构体
	// 满足某一条件的全部记录数目
	// 测试 100
	// 20
	h := H{
		Code:  code,
		Rows:  data,
		Total: total,
	}
	// 将结构体转化成JSOn字符串
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	// 输出
	w.Write(ret)
}
