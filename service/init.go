// Package service 初始化
package service

import (
	"fmt"
	"log"

	"ness_monster/model"

	"xorm.io/xorm"
)

// DBEngin xorm 引擎
var DBEngin *xorm.Engine

// 初始化数据库
func init() {
	driverName := "mysql"
	dsn := "root:211010@(127.0.0.1:13306)/ness_monster?charset=utf8"
	var err error
	DBEngin, err = xorm.NewEngine(driverName, dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	// 是否显示sql语句
	DBEngin.ShowSQL(true)
	// 数据库最大链接数，线上环境自己设置，这个直接影响数据库的性能
	DBEngin.SetMaxOpenConns(2)

	DBEngin.Sync2(new(model.User))

	fmt.Println("init data base ok")
}
