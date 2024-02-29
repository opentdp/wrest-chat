package profile

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wechat-rest/dbase/tables"
)

// 创建配置

type CreateParam struct {
	Wxid      string `binding:"required" json:"wxid"`
	Roomid    string `json:"roomid"`
	Level     int32  `json:"level"`
	Remark    string `json:"remark"`
	AiArgot   string `json:"ai_argot"`
	AiModel   string `json:"ai_model"`
	BanExpire int64  `json:"ban_expire"`
}

func Create(data *CreateParam) (uint, error) {

	item := &tables.Profile{
		Wxid:      data.Wxid,
		Roomid:    data.Roomid,
		Level:     data.Level,
		Remark:    data.Remark,
		AiArgot:   data.AiArgot,
		AiModel:   data.AiModel,
		BanExpire: data.BanExpire,
	}

	result := dborm.Db.Create(item)

	return item.Rd, result.Error

}

// 更新配置

type UpdateParam = CreateParam

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&tables.Profile{
			Wxid:   data.Wxid,
			Roomid: data.Roomid,
		}).
		Updates(tables.Profile{
			Level:     data.Level,
			Remark:    data.Remark,
			AiArgot:   data.AiArgot,
			AiModel:   data.AiModel,
			BanExpire: data.BanExpire,
		})

	return result.Error

}

// 合并配置

type MigrateParam = CreateParam

func Migrate(data *MigrateParam) error {

	item, err := Fetch(&FetchParam{
		Wxid:   data.Wxid,
		Roomid: data.Roomid,
	})

	if err == nil && item.Rd > 0 {
		err = Update(data)
	} else {
		_, err = Create(data)
	}

	return err

}

// 获取配置

type FetchParam struct {
	Wxid   string `binding:"required" json:"wxid"`
	Roomid string `json:"roomid"`
}

func Fetch(data *FetchParam) (*tables.Profile, error) {

	var item *tables.Profile

	result := dborm.Db.
		Where(&tables.Profile{
			Wxid:   data.Wxid,
			Roomid: data.Roomid,
		}).
		First(&item)

	if item == nil {
		item = &tables.Profile{Wxid: data.Wxid, Roomid: data.Roomid}
	}

	return item, result.Error

}

// 删除配置

type DeleteParam = FetchParam

func Delete(data *DeleteParam) error {

	var item *tables.Profile

	result := dborm.Db.
		Where(&tables.Profile{
			Wxid:   data.Wxid,
			Roomid: data.Roomid,
		}).
		Delete(&item)

	return result.Error

}

// 获取配置列表

type FetchAllParam struct {
	Wxid   string `json:"wxid"`
	Roomid string `json:"roomid"`
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

type CountParam = FetchAllParam

func Count(data *CountParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&tables.Profile{}).
		Where(&tables.Profile{
			Wxid:   data.Wxid,
			Roomid: data.Roomid,
		}).
		Count(&count)

	return count, result.Error

}
