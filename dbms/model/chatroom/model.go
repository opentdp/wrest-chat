package chatroom

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wechat-rest/dbms/model"
)

// 创建配置

type CreateParam struct {
	Rd        uint
	Roomid    string `binding:"required"`
	Name      string
	Level     int32
	Remark    string
	JoinArgot string
	Welcome   string
}

func Create(data *CreateParam) (uint, error) {

	item := &model.Chatroom{
		Roomid:    data.Roomid,
		Name:      data.Name,
		Level:     data.Level,
		Remark:    data.Remark,
		JoinArgot: data.JoinArgot,
		Welcome:   data.Welcome,
	}

	result := dborm.Db.Create(item)

	return item.Rd, result.Error

}

// 更新配置

type UpdateParam struct {
	Rd        uint
	Name      string
	Level     int32
	Remark    string
	JoinArgot string
	Welcome   string
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&model.Chatroom{
			Rd: data.Rd,
		}).
		Updates(model.Chatroom{
			Name:      data.Name,
			Level:     data.Level,
			Remark:    data.Remark,
			JoinArgot: data.JoinArgot,
			Welcome:   data.Welcome,
		})

	return result.Error

}

// 删除配置

type DeleteParam struct {
	Rd uint
}

func Delete(data *DeleteParam) error {

	var item *model.Chatroom

	result := dborm.Db.
		Where(&model.Chatroom{
			Rd: data.Rd,
		}).
		Delete(&item)

	return result.Error

}

// 获取配置

type FetchParam struct {
	Rd     uint
	Roomid string
}

func Fetch(data *FetchParam) (*model.Chatroom, error) {

	var item *model.Chatroom

	result := dborm.Db.
		Where(&model.Chatroom{
			Rd:     data.Rd,
			Roomid: data.Roomid,
		}).
		First(&item)

	return item, result.Error

}

// 获取配置列表

type FetchAllParam struct {
	Level int32
}

func FetchAll(data *FetchAllParam) ([]*model.Chatroom, error) {

	var items []*model.Chatroom

	result := dborm.Db.
		Where(&model.Chatroom{
			Level: data.Level,
		}).
		Find(&items)

	return items, result.Error

}

// 获取配置总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Where(&model.Chatroom{
			Level: data.Level,
		}).
		Count(&count)

	return count, result.Error

}
