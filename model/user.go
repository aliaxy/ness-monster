// Package model 用户模型
package model

import "time"

const (
	SexMan     = "M" // SexMan 男性
	SexWoman   = "W" // SexWoman 女性
	SexUnknown = "U" // SexUnknown 未知
)

// User 用户
type User struct {
	// 用户 id
	ID int64 `xorm:"'id' pk autoincr bigint(20)" form:"id" json:"id"`
	// 手机号 唯一的
	Mobile string `xorm:"varchar(20)" form:"mobile" json:"mobile"`
	// 用户密码 f(plainpwd+salt),MD5
	Passwd string `xorm:"varchar(40)" form:"passwd" json:"-"`
	// 头像
	Avatar string `xorm:"varchar(150)" form:"avatar" json:"avatar"`
	// 性别
	Sex string `xorm:"varchar(2)" form:"sex" json:"sex"`
	// 昵称
	Nickname string `xorm:"varchar(20)" form:"nickname" json:"nickname"`
	// 随机数
	Salt string `xorm:"varchar(10)" form:"salt" json:"-"`
	// 是否在线
	Online int `xorm:"int(10)" form:"online" json:"online"`
	// 前端用户登录鉴权 chat?id=1&token=x
	Token string `xorm:"varchar(40)" form:"token" json:"token"`
	Memo  string `xorm:"varchar(140)" form:"memo" json:"memo"`
	// 统计每天用户增量时间
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"`
}
