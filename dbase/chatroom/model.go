package chatroom

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wechat-rest/dbase/tables"
)

// 创建配置

type CreateParam struct {
	Roomid     string `binding:"required"`
	Name       string
	Level      int32
	Remark     string
	JoinArgot  string
	WelcomeMsg string
}

func Create(data *CreateParam) (uint, error) {

	item := &tables.Chatroom{
		Roomid:     data.Roomid,
		Name:       data.Name,
		Level:      data.Level,
		Remark:     data.Remark,
		JoinArgot:  data.JoinArgot,
		WelcomeMsg: data.WelcomeMsg,
	}

	result := dborm.Db.Create(item)

	return item.Rd, result.Error

}

// 更新配置

type UpdateParam = CreateParam

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&tables.Chatroom{
			Roomid: data.Roomid,
		}).
		Updates(tables.Chatroom{
			Name:       data.Name,
			Level:      data.Level,
			Remark:     data.Remark,
			JoinArgot:  data.JoinArgot,
			WelcomeMsg: data.WelcomeMsg,
		})

	return result.Error

}

// 删除配置

type DeleteParam struct {
	Roomid string `binding:"required"`
}

func Delete(data *DeleteParam) error {

	var item *tables.Chatroom

	result := dborm.Db.
		Where(&tables.Chatroom{
			Roomid: data.Roomid,
		}).
		Delete(&item)

	return result.Error

}

// 获取配置

type FetchParam struct {
	Roomid string `binding:"required"`
}

func Fetch(data *FetchParam) (*tables.Chatroom, error) {

	var item *tables.Chatroom

	result := dborm.Db.
		Where(&tables.Chatroom{
			Roomid: data.Roomid,
		}).
		First(&item)

	if item == nil {
		item = &tables.Chatroom{Roomid: data.Roomid}
	}

	return item, result.Error

}

// 获取配置列表

type FetchAllParam struct {
	Level int32
}

func FetchAll(data *FetchAllParam) ([]*tables.Chatroom, error) {

	var items []*tables.Chatroom

	result := dborm.Db.
		Where(&tables.Chatroom{
			Level: data.Level,
		}).
		Find(&items)

	return items, result.Error

}

// 获取配置总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Where(&tables.Chatroom{
			Level: data.Level,
		}).
		Count(&count)

	return count, result.Error

}
