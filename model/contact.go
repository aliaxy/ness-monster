// Package model 联系人模型
package model

import "time"

// Contact 联系人
type Contact struct {
	Id      int64 `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Ownerid int64 `xorm:"bigint(20)" form:"ownerid" json:"ownerid"`
	Dstobj  int64 `xorm:"bigint(20)" form:"dstobj" json:"dstobj"`
	// 是下面定义的常量 用户加用户是0x01  用户加群0x02
	Cate     int       `xorm:"int(11)" form:"cate" json:"cate"`
	Memo     string    `xorm:"varchar(120)" form:"memo" json:"memo"` // 什么角色
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"`
}

const (
	ConcatCateUser      = 0x01 // ConcatCateUser 用户
	ConcatCateCommunity = 0x02 // ConcatCateCommunity 群组
)
