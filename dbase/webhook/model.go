package webhook

import (
	"errors"

	"github.com/opentdp/go-helper/dborm"
	"github.com/opentdp/wrest-chat/dbase/tables"
)

// 创建 webhook

type CreateWebhookParam struct {
	TargetId string `json:"target_id" binding:"required"`
	Remark   string `json:"remark"`
}

func Create(data *CreateWebhookParam) (string, error) {

	token := generateGUID()
	item := &tables.Webhook{
		TargetId: data.TargetId,
		Token:    token,
		Remark:   data.Remark,
	}

	result := dborm.Db.Create(item)

	return token, result.Error

}

// 获取 webhook

type FetchWebhookParam struct {
	Rd       uint   `json:"rd"`
	TargetId string `json:"target_id"`
	Token    string `json:"token"`
}

func Fetch(data *FetchWebhookParam) (*tables.Webhook, error) {

	var item *tables.Webhook

	result := dborm.Db.
		Where(&tables.Webhook{
			Rd:       data.Rd,
			TargetId: data.TargetId,
			Token:    data.Token,
		}).
		First(&item)

	if item == nil {
		return nil, errors.New("未找到记录")
	}

	return item, result.Error

}

// 删除 webhook

type DeleteWebhookParam = FetchWebhookParam

func Delete(data *DeleteWebhookParam) error {

	var item *tables.Webhook

	result := dborm.Db.
		Where(&tables.Webhook{
			Rd:       data.Rd,
			TargetId: data.TargetId,
			Token:    data.Token,
		}).
		Delete(&item)

	return result.Error

}

// 获取 webhook 列表

func FetchAll() ([]*tables.Webhook, error) {

	var items []*tables.Webhook

	result := dborm.Db.Find(&items)

	return items, result.Error

}

// 获取 webhook 总数

func Count() (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&tables.LLModel{}).
		Count(&count)

	return count, result.Error

}
