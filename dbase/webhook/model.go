package webhook

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/opentdp/go-helper/dborm"
	"github.com/opentdp/wrest-chat/dbase/tables"
	"io"
)

// 创建webhook

type CreateWebHookParam struct {
	TargetId string `binding:"required" json:"target_id"`
	Remark   string `json:"remark"`
}

func Create(data *CreateWebHookParam) (uint, string, error) {
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

// 查询webhook

type FetchWebHookParam struct {
	Rd uint `json:"rd"`
}

func Fetch(data *FetchWebHookParam) (*tables.Webhook, error) {
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

// 删除webhook

type DeleteWebHookParam = FetchWebHookParam

func Delete(data *DeleteWebHookParam) error {
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

// 获取全部webhook

func FetchAll() ([]*tables.Webhook, error) {
	var items []*tables.Webhook

	result := dborm.Db.Find(&items)

	return items, result.Error
}

// 通过token查询webhook

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

// 生成GUID

func generateGUID() string {
	// 创建一个16字节的切片用于存储随机数据
	byteGUID := make([]byte, 16)
	// 使用加密的随机数源生成随机数填充切片
	_, err := io.ReadFull(rand.Reader, byteGUID)
	if err != nil {
		panic("无法生成GUID")
	}

	// 设置版本号和变体以符合GUID版本4的标准
	byteGUID[8] = byteGUID[8]&^0xc0 | 0x80
	byteGUID[6] = byteGUID[6]&^0xf0 | 0x40

	// 将16字节的字节切片转换为32字节的字节切片
	guidBytes := append(byteGUID[0:4], byteGUID[4:6]...)
	guidBytes = append(guidBytes, byteGUID[6:8]...)
	guidBytes = append(guidBytes, byteGUID[8:10]...)
	guidBytes = append(guidBytes, byteGUID[10:16]...)

	// 使用十六进制编码32字节的字节切片
	return hex.EncodeToString(guidBytes)
}
