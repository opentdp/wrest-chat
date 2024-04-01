package chatroom

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wrest-chat/dbase/tables"
)

// 创建群聊

type CreateParam struct {
	Rd           uint   `json:"rd"`
	Roomid       string `json:"roomid" binding:"required"`
	Name         string `json:"name"`
	Level        int32  `json:"level"`
	Remark       string `json:"remark"`
	JoinArgot    string `json:"join_argot"`
	PatReturn    string `json:"pat_return"`
	RevokeMsg    string `json:"revoke_msg"`
	WelcomeMsg   string `json:"welcome_msg"`
	ModelContext string `json:"model_context"`
	ModelDefault string `json:"model_default"`
	ModelHistory int    `json:"model_history"`
}

func Create(data *CreateParam) (uint, error) {

	item := &tables.Chatroom{
		Roomid:       data.Roomid,
		Name:         data.Name,
		Level:        data.Level,
		Remark:       data.Remark,
		JoinArgot:    data.JoinArgot,
		PatReturn:    data.PatReturn,
		RevokeMsg:    data.RevokeMsg,
		WelcomeMsg:   data.WelcomeMsg,
		ModelContext: data.ModelContext,
		ModelDefault: data.ModelDefault,
		ModelHistory: data.ModelHistory,
	}

	result := dborm.Db.Create(item)

	return item.Rd, result.Error

}

// 更新群聊

type UpdateParam = CreateParam

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&tables.Chatroom{
			Rd: data.Rd,
		}).
		Updates(tables.Chatroom{
			Roomid:       data.Roomid,
			Name:         data.Name,
			Level:        data.Level,
			Remark:       data.Remark,
			JoinArgot:    data.JoinArgot,
			PatReturn:    data.PatReturn,
			RevokeMsg:    data.RevokeMsg,
			WelcomeMsg:   data.WelcomeMsg,
			ModelContext: data.ModelContext,
			ModelDefault: data.ModelDefault,
			ModelHistory: data.ModelHistory,
		})

	return result.Error

}

// 合并群聊

type ReplaceParam = CreateParam

func Replace(data *ReplaceParam) error {

	rq := &FetchParam{Rd: data.Rd}
	if rq.Rd == 0 {
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

// 获取群聊

type FetchParam struct {
	Rd     uint   `json:"rd"`
	Roomid string `json:"roomid"`
}

func Fetch(data *FetchParam) (*tables.Chatroom, error) {

	var item *tables.Chatroom

	result := dborm.Db.
		Where(&tables.Chatroom{
			Rd:     data.Rd,
			Roomid: data.Roomid,
		}).
		First(&item)

	if item == nil {
		item = &tables.Chatroom{Roomid: data.Roomid}
	}

	return item, result.Error

}

// 删除群聊

type DeleteParam = FetchParam

func Delete(data *DeleteParam) error {

	var item *tables.Chatroom

	result := dborm.Db.
		Where(&tables.Chatroom{
			Rd:     data.Rd,
			Roomid: data.Roomid,
		}).
		Delete(&item)

	return result.Error

}

// 获取群聊列表

type FetchAllParam struct {
	Level int32 `json:"level"`
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

// 获取群聊总数

type CountParam = FetchAllParam

func Count(data *CountParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&tables.Chatroom{}).
		Where(&tables.Chatroom{
			Level: data.Level,
		}).
		Count(&count)

	return count, result.Error

}
