package message

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wechat-rest/dbase/tables"
)

// 创建消息

type CreateParam struct {
	Id      uint64 `binding:"required"`
	IsSelf  bool
	IsGroup bool
	Type    uint32
	Ts      uint32
	Roomid  string
	Content string
	Sender  string
	Sign    string
	Thumb   string
	Extra   string
	Xml     string
}

func Create(data *CreateParam) (uint, error) {

	item := &tables.Message{
		Id:      data.Id,
		IsSelf:  data.IsSelf,
		IsGroup: data.IsGroup,
		Type:    data.Type,
		Ts:      data.Ts,
		Roomid:  data.Roomid,
		Content: data.Content,
		Sender:  data.Sender,
		Sign:    data.Sign,
		Thumb:   data.Thumb,
		Extra:   data.Extra,
		Xml:     data.Xml,
	}

	result := dborm.Db.Create(item)

	return item.Rd, result.Error

}

// 更新消息

type UpdateParam = CreateParam

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&tables.Message{
			Id: data.Id,
		}).
		Updates(tables.Message{
			IsSelf:  data.IsSelf,
			IsGroup: data.IsGroup,
			Type:    data.Type,
			Ts:      data.Ts,
			Roomid:  data.Roomid,
			Content: data.Content,
			Sender:  data.Sender,
			Sign:    data.Sign,
			Thumb:   data.Thumb,
			Extra:   data.Extra,
			Xml:     data.Xml,
		})

	return result.Error

}

// 删除消息

type DeleteParam struct {
	Id uint64 `binding:"required"`
}

func Delete(data *DeleteParam) error {

	var item *tables.Message

	result := dborm.Db.
		Where(&tables.Message{
			Id: data.Id,
		}).
		Delete(&item)

	return result.Error

}

// 获取消息

type FetchParam struct {
	Id uint64 `binding:"required"`
}

func Fetch(data *FetchParam) (*tables.Message, error) {

	var item *tables.Message

	result := dborm.Db.
		Where(&tables.Message{
			Id: data.Id,
		}).
		First(&item)

	if item == nil {
		item = &tables.Message{Id: data.Id}
	}

	return item, result.Error

}

// 获取消息列表

type FetchAllParam struct {
	Sender string
	Roomid string
}

func FetchAll(data *FetchAllParam) ([]*tables.Message, error) {

	var items []*tables.Message

	result := dborm.Db.
		Where(&tables.Message{
			Sender: data.Sender,
			Roomid: data.Roomid,
		}).
		Find(&items)

	return items, result.Error

}

// 获取消息总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Where(&tables.Message{
			Sender: data.Sender,
			Roomid: data.Roomid,
		}).
		Count(&count)

	return count, result.Error

}
