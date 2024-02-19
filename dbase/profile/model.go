package profile

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wechat-rest/dbase/tables"
)

// 创建配置

type CreateParam struct {
	Rd      uint
	Wxid    string `binding:"required"`
	Roomid  string
	Level   int32
	AiArgot string
	AiModel string
}

func Create(data *CreateParam) (uint, error) {

	item := &tables.Profile{
		Wxid:    data.Wxid,
		Roomid:  data.Roomid,
		Level:   data.Level,
		AiArgot: data.AiArgot,
		AiModel: data.AiModel,
	}

	result := dborm.Db.Create(item)

	return item.Rd, result.Error

}

// 更新配置

type UpdateParam struct {
	Rd      uint
	Level   int32
	AiArgot string
	AiModel string
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&tables.Profile{
			Rd: data.Rd,
		}).
		Updates(tables.Profile{
			Level:   data.Level,
			AiArgot: data.AiArgot,
			AiModel: data.AiModel,
		})

	return result.Error

}

// 删除配置

type DeleteParam struct {
	Rd uint
}

func Delete(data *DeleteParam) error {

	var item *tables.Profile

	result := dborm.Db.
		Where(&tables.Profile{
			Rd: data.Rd,
		}).
		Delete(&item)

	return result.Error

}

// 获取配置

type FetchParam struct {
	Rd     uint
	Wxid   string
	Roomid string
}

func Fetch(data *FetchParam) (*tables.Profile, error) {

	var item *tables.Profile

	result := dborm.Db.
		Where(&tables.Profile{
			Rd:     data.Rd,
			Wxid:   data.Wxid,
			Roomid: data.Roomid,
		}).
		First(&item)

	return item, result.Error

}

// 获取配置列表

type FetchAllParam struct {
	Wxid   string
	Roomid string
}

func FetchAll(data *FetchAllParam) ([]*tables.Profile, error) {

	var items []*tables.Profile

	result := dborm.Db.
		Where(&tables.Profile{
			Wxid:   data.Wxid,
			Roomid: data.Roomid,
		}).
		Find(&items)

	return items, result.Error

}

// 获取配置总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Where(&tables.Profile{
			Wxid:   data.Wxid,
			Roomid: data.Roomid,
		}).
		Count(&count)

	return count, result.Error

}
