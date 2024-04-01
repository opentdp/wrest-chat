package message

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wrest-chat/dbase/tables"
)

// 创建消息

type CreateParam struct {
	Id      uint64 `json:"id" binding:"required"`
	IsSelf  bool   `json:"is_self"`
	IsGroup bool   `json:"is_group"`
	Type    uint32 `json:"type"`
	Ts      uint32 `json:"ts"`
	Roomid  string `json:"roomid"`
	Content string `json:"content"`
	Sender  string `json:"sender"`
	Sign    string `json:"sign"`
	Thumb   string `json:"thumb"`
	Extra   string `json:"extra"`
	Xml     string `json:"xml"`
	Remark  string `json:"remark"`
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
		Remark:  data.Remark,
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
			Remark:  data.Remark,
		})

	return result.Error

}

// 合并消息

type ReplaceParam = CreateParam

func Replace(data *ReplaceParam) error {

	rq := &FetchParam{
		Id: data.Id,
	}

	item, err := Fetch(rq)
	if err == nil && item.Rd > 0 {
		err = Update(data)
	} else {
		_, err = Create(data)
	}

	return err

}

// 获取消息

type FetchParam struct {
	Id uint64 `json:"id" binding:"required"`
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

// 删除消息

type DeleteParam = FetchParam

func Delete(data *DeleteParam) error {

	var item *tables.Message

	result := dborm.Db.
		Where(&tables.Message{
			Id: data.Id,
		}).
		Delete(&item)

	return result.Error

}

// 获取消息列表

type FetchAllParam struct {
	Sender string `json:"sender"`
	Roomid string `json:"roomid"`
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

type CountParam = FetchAllParam

func Count(data *CountParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&tables.Message{}).
		Where(&tables.Message{
			Sender: data.Sender,
			Roomid: data.Roomid,
		}).
		Count(&count)

	return count, result.Error

}
