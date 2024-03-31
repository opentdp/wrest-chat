package webhook

import (
	"errors"

	"github.com/opentdp/go-helper/dborm"
	"github.com/opentdp/wrest-chat/dbase/tables"
)

// 创建 webhook

type CreateWebhookParam struct {
	TargetId string `binding:"required" json:"target_id"`
	Remark   string `json:"remark"`
}

func Create(data *CreateWebhookParam) (uint, string, error) {

	if exist(data.TargetId) {
		return 0, "", errors.New("已存在webhook")
	}

	token := generateGUID()
	item := &tables.Webhook{
		TargetId: data.TargetId,
		Token:    token,
		Remark:   data.Remark,
	}

	result := dborm.Db.Create(item)

	return item.Rd, token, result.Error

}

func exist(id string) bool {

	var item *tables.Webhook

	result := dborm.Db.Where(&tables.Webhook{TargetId: id}).First(&item)

	if result.Error != nil {
		return false
	}

	return item != nil

}

// 查询 webhook

type FetchWebhookParam struct {
	Rd uint `json:"rd"`
}

func Fetch(data *FetchWebhookParam) (*tables.Webhook, error) {

	var item *tables.Webhook

	result := dborm.Db.
		Where(&tables.Webhook{
			Rd: data.Rd,
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
			Rd: data.Rd,
		}).
		Delete(&item)

	return result.Error

}

func DeleteByTargetId(targetId string) bool {

	var item *tables.Webhook

	result := dborm.Db.Where(&tables.Webhook{TargetId: targetId}).Delete(&item)

	return result.Error == nil

}

// 获取全部 webhook

func FetchAll() ([]*tables.Webhook, error) {

	var items []*tables.Webhook

	result := dborm.Db.Find(&items)

	return items, result.Error

}

// 通过 token 查询 webhook

func FindByToken(token string) (*tables.Webhook, error) {

	var item *tables.Webhook

	result := dborm.Db.
		Where(&tables.Webhook{
			Token: token,
		}).
		First(&item)

	if result.Error != nil {
		return nil, result.Error
	}

	return item, nil

}
