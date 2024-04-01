package profile

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wrest-chat/dbase/tables"
)

// 创建配置

type CreateParam struct {
	Rd        uint   `json:"rd"`
	Wxid      string `json:"wxid" binding:"required"`
	Roomid    string `json:"roomid"`
	Level     int32  `json:"level"`
	Remark    string `json:"remark"`
	AiModel   string `json:"ai_model"`
	BanExpire int64  `json:"ban_expire"`
}

func Create(data *CreateParam) (uint, error) {

	item := &tables.Profile{
		Wxid:      data.Wxid,
		Roomid:    data.Roomid,
		Level:     data.Level,
		Remark:    data.Remark,
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
			Rd: data.Rd,
		}).
		Updates(tables.Profile{
			Wxid:      data.Wxid,
			Roomid:    data.Roomid,
			Level:     data.Level,
			Remark:    data.Remark,
			AiModel:   data.AiModel,
			BanExpire: data.BanExpire,
		})

	return result.Error

}

// 合并配置

type ReplaceParam = CreateParam

func Replace(data *ReplaceParam) error {

	rq := &FetchParam{Rd: data.Rd}
	if rq.Rd == 0 {
		rq.Wxid = data.Wxid
		rq.Roomid = data.Roomid
	}

	item, err := Fetch(rq)
	if err == nil && item.Rd > 0 {
		data.Rd = item.Rd
		err = Update(data)
	} else {
		_, err = Create(data)
	}

	return err

}

// 获取配置

type FetchParam struct {
	Rd     uint   `json:"rd"`
	Wxid   string `json:"wxid"`
	Roomid string `json:"roomid"`
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
			Rd:     data.Rd,
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
	Level  int32  `json:"level"`
}

func FetchAll(data *FetchAllParam) ([]*tables.Profile, error) {

	var items []*tables.Profile

	result := dborm.Db.
		Where(&tables.Profile{
			Wxid:   data.Wxid,
			Roomid: data.Roomid,
			Level:  data.Level,
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
			Level:  data.Level,
		}).
		Count(&count)

	return count, result.Error

}
