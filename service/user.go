// Package service 提供用户相关服务
package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"ness_monster/model"
	"ness_monster/util"
)

// UserService 用户服务
type UserService struct{}

// Register 注册函数
func (u *UserService) Register(mobile, plainPwd, nickname, avatar, sex string) (user *model.User, err error) {
	user = new(model.User)
	_, err = DBEngin.Where("mobile = ?", mobile).Get(user)
	if err != nil {
		return
	}

	if user.Id > 0 {
		return &model.User{}, errors.New("该手机号已经注册")
	}

	user.Mobile = mobile
	user.Avatar = avatar
	user.Nickname = nickname
	user.Sex = sex
	user.Salt = fmt.Sprintf("%06d", rand.Int31n(100000))
	user.Passwd = util.MakePasswd(plainPwd, user.Salt)
	// 创建时间，用来统计用户的注册量
	user.Createat = time.Now()
	user.Token = fmt.Sprintf("%08d", rand.Int31n(10000000))
	// 插入InsertOne
	_, err = DBEngin.Insert(user)

	return
}

// Login 登录函数
func (u *UserService) Login(mobile, plainPwd string) (user *model.User, err error) {
	user = new(model.User)
	// 通过手机号查询用户
	DBEngin.Where("mobile = ?", mobile).Get(user)
	if user.Id == 0 {
		return &model.User{}, errors.New("该用户不存在")
	}

	// 比对密码
	if util.ValidatePasswd(plainPwd, user.Salt, user.Passwd) == false {
		return &model.User{}, errors.New("帐号或密码错误")
	}

	// 刷新 token 安全
	str := fmt.Sprintf("%d", time.Now().Unix())
	user.Token = util.MD5Encode(str)

	// 返回数据
	_, err = DBEngin.ID(user.Id).Cols("token").Update(user)
	return
}

// Find 查找用户
func (u *UserService) Find(userID int64) (user *model.User) {
	user = new(model.User)
	DBEngin.ID(userID).Get(user)
	return
}
