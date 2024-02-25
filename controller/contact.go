// Package controller 联系人控制层
package controller

import (
	"fmt"
	"net/http"

	"ness_monster/arg"
	"ness_monster/model"
	"ness_monster/service"
	"ness_monster/util"
)

var contactService service.ContactService

// LoadFriend 载入好友
func LoadFriend(w http.ResponseWriter, req *http.Request) {
	arg := new(arg.ContactArg)

	// 如果这个用的上,那么可以直接
	util.Bind(req, arg)

	users := contactService.SearchFriend(arg.Userid)
	util.RespSuccessList(w, users, len(users))
}

// LoadCommunity 载入群
func LoadCommunity(w http.ResponseWriter, req *http.Request) {
	arg := new(arg.ContactArg)

	// 如果这个用的上,那么可以直接
	util.Bind(req, arg)
	comunitys := contactService.SearchComunity(arg.Userid)
	util.RespSuccessList(w, comunitys, len(comunitys))
}

// JoinCommunity 加入群
func JoinCommunity(w http.ResponseWriter, req *http.Request) {
	arg := new(arg.ContactArg)

	// 如果这个用的上,那么可以直接
	util.Bind(req, arg)

	fmt.Println(arg.Userid)

	err := contactService.JoinCommunity(arg.Userid, arg.Dstid)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		AddGroupID(arg.Userid, arg.Dstid)
		util.RespSuccess(w, nil, "")
	}
}

// Addfriend 添加好友
func Addfriend(w http.ResponseWriter, req *http.Request) {
	arg := new(arg.ContactArg)
	// 对象绑定
	util.Bind(req, arg)

	err := contactService.AddFriend(arg.Userid, arg.Dstid)

	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespSuccess(w, nil, "好友添加成功")
	}
}

// Createcommunity 创建群
func Createcommunity(w http.ResponseWriter, req *http.Request) {
	arg := new(model.Community)
	util.Bind(req, arg)
	com, err := contactService.CreateCommunity(arg)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespSuccess(w, com, "")
	}
}
