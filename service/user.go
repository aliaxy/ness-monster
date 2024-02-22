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
func (s *UserService) Register(mobile, plainpwd, nickname, avatar, sex string) (user *model.User, err error) {
	user = new(model.User)
	_, err = DBEngin.Where("mobile = ?", mobile).Get(user)
	if err != nil {
		return
	}

	if user.ID > 0 {
		return &model.User{}, errors.New("该手机号已经注册")
	}

	user.Mobile = mobile
	user.Avatar = avatar
	user.Nickname = nickname
	user.Sex = sex
	user.Salt = fmt.Sprintf("%06d", rand.Int31n(100000))
	user.Passwd = util.MakePasswd(plainpwd, user.Salt)
	// 创建时间，用来统计用户的注册量
	user.Createat = time.Now()
	user.Token = fmt.Sprintf("%08d", rand.Int31n(10000000))

	// 插入InsertOne
	_, err = DBEngin.Insert(user)

	return
}

// Login 登录函数
func (s *UserService) Login(mobile, plainpwd string) (user *model.User, err error) {
	return &model.User{}, nil
}
