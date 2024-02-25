// Package service 提供联系人相关服务
package service

import (
	"errors"
	"time"

	"ness_monster/model"
)

// ContactService 联系人服务
type ContactService struct{}

// AddFriend 添加好友
func (c *ContactService) AddFriend(userID, dstID int64) (err error) {
	if userID == dstID {
		return errors.New("不能添加自己为好友啊")
	}

	// 判断是否已经加了好友
	contact := new(model.Contact)
	DBEngin.Where("ownerid = ?", userID).
		And("dstobj = ?", dstID).
		And("cate = ?", model.ConcatCateUser).
		Get(contact)

	if contact.Id > 0 {
		return errors.New("该用户已经被添加过啦")
	}

	// 事务
	session := DBEngin.NewSession()
	session.Begin()

	// 插入自己的
	_, err = session.InsertOne(model.Contact{
		Ownerid:  userID,
		Dstobj:   dstID,
		Cate:     model.ConcatCateUser,
		Createat: time.Now(),
	})

	if err != nil {
		session.Rollback()
		return
	}

	// 插入对方的
	_, err = session.InsertOne(model.Contact{
		Ownerid:  dstID,
		Dstobj:   userID,
		Cate:     model.ConcatCateUser,
		Createat: time.Now(),
	})

	if err != nil {
		session.Rollback()
		return
	}

	session.Commit()
	return
}

// SearchComunity 查找群
func (c *ContactService) SearchComunity(userID int64) []model.Community {
	conconts := make([]model.Contact, 0)
	comIds := make([]int64, 0)

	DBEngin.Where("ownerid = ? and cate = ?", userID, model.ConcatCateCommunity).Find(&conconts)
	for _, v := range conconts {
		comIds = append(comIds, v.Dstobj)
	}
	coms := make([]model.Community, 0)
	if len(comIds) == 0 {
		return coms
	}
	DBEngin.In("id", comIds).Find(&coms)
	return coms
}

// SearchComunityIds 查找群
func (c *ContactService) SearchComunityIds(userID int64) (comIds []int64) {
	// todo 获取用户全部群ID
	conconts := make([]model.Contact, 0)
	comIds = make([]int64, 0)

	DBEngin.Where("ownerid = ? and cate = ?", userID, model.ConcatCateCommunity).Find(&conconts)
	for _, v := range conconts {
		comIds = append(comIds, v.Dstobj)
	}
	return comIds
}

// JoinCommunity 加群
func (c *ContactService) JoinCommunity(userID, comID int64) error {
	cot := model.Contact{
		Ownerid: userID,
		Dstobj:  comID,
		Cate:    model.ConcatCateCommunity,
	}
	DBEngin.Get(&cot)
	if cot.Id == 0 {
		cot.Createat = time.Now()
		_, err := DBEngin.InsertOne(cot)
		return err
	}
	return nil
}

// CreateCommunity 建群
func (c *ContactService) CreateCommunity(comm *model.Community) (ret model.Community, err error) {
	if len(comm.Name) == 0 {
		err = errors.New("缺少群名称")
		return ret, err
	}
	if comm.Ownerid == 0 {
		err = errors.New("请先登录")
		return ret, err
	}
	com := model.Community{
		Ownerid: comm.Ownerid,
	}
	num, err := DBEngin.Count(&com)

	if num > 5 {
		err = errors.New("一个用户最多只能创见5个群")
		return com, err
	}

	comm.Createat = time.Now()
	session := DBEngin.NewSession()
	session.Begin()
	_, err = session.InsertOne(comm)
	if err != nil {
		session.Rollback()
		return com, err
	}
	_, err = session.InsertOne(
		model.Contact{
			Ownerid:  comm.Ownerid,
			Dstobj:   comm.Id,
			Cate:     model.ConcatCateCommunity,
			Createat: time.Now(),
		})
	if err != nil {
		session.Rollback()
	} else {
		session.Commit()
	}
	return com, err
}

// SearchFriend 查找好友
func (c *ContactService) SearchFriend(userID int64) []model.User {
	conconts := make([]model.Contact, 0)
	objIds := make([]int64, 0)
	DBEngin.Where("ownerid = ? and cate = ?", userID, model.ConcatCateUser).Find(&conconts)
	for _, v := range conconts {
		objIds = append(objIds, v.Dstobj)
	}
	coms := make([]model.User, 0)
	if len(objIds) == 0 {
		return coms
	}
	DBEngin.In("id", objIds).Find(&coms)
	return coms
}
