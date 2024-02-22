// Package controller 用户控制层
package controller

import (
	"fmt"
	"math/rand"
	"net/http"

	"ness_monster/model"
	"ness_monster/service"
	"ness_monster/util"
)

var userService service.UserService

// UserLogin 用户登录
func UserLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认是不会解析的
	mobile := r.PostForm.Get("mobile")
	passwd := r.PostForm.Get("passwd")

	loginOk := false
	if mobile == "123456789" && passwd == "123456" {
		loginOk = true
	}

	if loginOk {
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		util.RespSuccess(w, data, "")
	} else {
		util.RespFail(w, "密码不正确")
	}
}

// UserRegister 用户注册
func UserRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认是不会解析的
	mobile := r.PostForm.Get("mobile")
	fmt.Println(mobile)
	passwd := r.PostForm.Get("passwd")
	nickname := fmt.Sprintf("user%06d", rand.Int31())
	avator := ""
	sex := model.SexUnknown

	user, err := userService.Register(mobile, passwd, nickname, avator, sex)
	if err != nil {
		util.RespSuccess(w, user, "")
	} else {
		util.RespFail(w, "密码不正确")
	}
}
