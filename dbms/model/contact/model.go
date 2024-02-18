package contact

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wechat-rest/dbms/model"
)

// 创建配置

type CreateParam struct {
	Rd       uint
	Wxid     string `binding:"required"`
	Code     string
	Remark   string
	Name     string
	Country  string
	Province string
	City     string
	Gender   int32
}

func Create(data *CreateParam) (uint, error) {

	item := &model.Contact{
		Wxid:    data.Wxid,
		Code:    data.Code,
		Remark:  data.Remark,
		Name:    data.Name,
		Country: data.Country,
		City:    data.City,
		Gender:  data.Gender,
	}

	result := dborm.Db.Create(item)

	return item.Rd, result.Error

}

// 更新配置

type UpdateParam struct {
	Rd       uint
	Code     string
	Remark   string
	Name     string
	Country  string
	Province string
	City     string
	Gender   int32
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&model.Contact{
			Rd: data.Rd,
		}).
		Updates(model.Contact{
			Code:    data.Code,
			Remark:  data.Remark,
			Name:    data.Name,
			Country: data.Country,
			City:    data.City,
			Gender:  data.Gender,
		})

	return result.Error

}

// 删除配置

type DeleteParam struct {
	Rd uint
}

func Delete(data *DeleteParam) error {

	var item *model.Contact

	result := dborm.Db.
		Where(&model.Contact{
			Rd: data.Rd,
		}).
		Delete(&item)

	return result.Error

}

// 获取配置

type FetchParam struct {
	Rd   uint
	Wxid string
}

func Fetch(data *FetchParam) (*model.Contact, error) {

	var item *model.Contact

	result := dborm.Db.
		Where(&model.Contact{
			Rd:   data.Rd,
			Wxid: data.Wxid,
		}).
		First(&item)

	return item, result.Error

}

// 获取配置列表

type FetchAllParam struct {
	Gender int32
}

func FetchAll(data *FetchAllParam) ([]*model.Contact, error) {

	var items []*model.Contact

	result := dborm.Db.
		Where(&model.Contact{
			Gender: data.Gender,
		}).
		Find(&items)

	return items, result.Error

}

// 获取配置总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Where(&model.Contact{
			Gender: data.Gender,
		}).
		Count(&count)

	return count, result.Error

}