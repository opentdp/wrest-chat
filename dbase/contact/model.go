package contact

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wrest-chat/dbase/tables"
)

// 创建联系人

type CreateParam struct {
	Wxid     string `json:"wxid" binding:"required"`
	Code     string `json:"code"`
	Remark   string `json:"remark"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	Gender   int32  `json:"gender"`
}

func Create(data *CreateParam) (uint, error) {

	item := &tables.Contact{
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

// 更新联系人

type UpdateParam = CreateParam

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&tables.Contact{
			Wxid: data.Wxid,
		}).
		Updates(tables.Contact{
			Code:    data.Code,
			Remark:  data.Remark,
			Name:    data.Name,
			Country: data.Country,
			City:    data.City,
			Gender:  data.Gender,
		})

	return result.Error

}

// 合并联系人

type ReplaceParam = CreateParam

func Replace(data *ReplaceParam) error {

	rq := &FetchParam{
		Wxid: data.Wxid,
	}

	item, err := Fetch(rq)
	if err == nil && item.Rd > 0 {
		err = Update(data)
	} else {
		_, err = Create(data)
	}

	return err

}

// 获取联系人

type FetchParam struct {
	Wxid string `json:"wxid" binding:"required"`
}

func Fetch(data *FetchParam) (*tables.Contact, error) {

	var item *tables.Contact

	result := dborm.Db.
		Where(&tables.Contact{
			Wxid: data.Wxid,
		}).
		First(&item)

	if item == nil {
		item = &tables.Contact{Wxid: data.Wxid}
	}

	return item, result.Error

}

// 删除联系人

type DeleteParam = FetchParam

func Delete(data *DeleteParam) error {

	var item *tables.Contact

	result := dborm.Db.
		Where(&tables.Contact{
			Wxid: data.Wxid,
		}).
		Delete(&item)

	return result.Error

}

// 获取联系人列表

type FetchAllParam struct {
	Gender int32 `json:"gender"`
}

func FetchAll(data *FetchAllParam) ([]*tables.Contact, error) {

	var items []*tables.Contact

	result := dborm.Db.
		Where(&tables.Contact{
			Gender: data.Gender,
		}).
		Find(&items)

	return items, result.Error

}

// 获取联系人总数

type CountParam = FetchAllParam

func Count(data *CountParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&tables.Contact{}).
		Where(&tables.Contact{
			Gender: data.Gender,
		}).
		Count(&count)

	return count, result.Error

}
