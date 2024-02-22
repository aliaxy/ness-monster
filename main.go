// Package main IM 入口
package main

import (
	"html/template"
	"log"
	"net/http"

	"ness_monster/controller"
	"ness_monster/service"

	_ "github.com/go-sql-driver/mysql"
)

var userService service.UserService

// RegisterView 注册模版页面
func RegisterView() {
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, v := range tpl.Templates() {
		tplname := v.Name()
		if tplname[0] != '/' {
			continue
		}
		http.HandleFunc(tplname, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer, tplname, nil)
		})
	}
}

func main() {
	// 绑定请求和处理函数
	http.HandleFunc("/user/login", controller.UserLogin)
	http.HandleFunc("/user/register", controller.UserRegister)

	// 提供静态资源支持
	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	// usre/login.shtml
	RegisterView()

	// 启动 HTTP 服务
	http.ListenAndServe(":8080", nil)
}
