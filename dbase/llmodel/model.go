package llmodel

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wrest-chat/dbase/tables"
)

// 创建模型

type CreateParam struct {
	Rd       uint   `json:"rd"`
	Mid      string `json:"mid" binding:"required"`
	Level    int32  `json:"level"`
	Family   string `json:"family" binding:"required"`
	Provider string `json:"provider" binding:"required"`
	Model    string `json:"model" binding:"required"`
	Secret   string `json:"secret" binding:"required"`
	Endpoint string `json:"endpoint"`
}

func Create(data *CreateParam) (uint, error) {

	item := &tables.LLModel{
		Mid:      data.Mid,
		Level:    data.Level,
		Family:   data.Family,
		Provider: data.Provider,
		Model:    data.Model,
		Secret:   data.Secret,
		Endpoint: data.Endpoint,
	}

	result := dborm.Db.Create(item)

	return item.Rd, result.Error

}

// 更新模型

type UpdateParam = CreateParam

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&tables.LLModel{
			Rd: data.Rd,
		}).
		Updates(tables.LLModel{
			Mid:      data.Mid,
			Level:    data.Level,
			Family:   data.Family,
			Provider: data.Provider,
			Model:    data.Model,
			Secret:   data.Secret,
			Endpoint: data.Endpoint,
		})

	return result.Error

}

// 合并模型

type ReplaceParam = CreateParam

func Replace(data *ReplaceParam) error {

	rq := &FetchParam{Rd: data.Rd}
	if rq.Rd == 0 {
		rq.Mid = data.Mid
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

// 获取模型

type FetchParam struct {
	Rd  uint   `json:"rd"`
	Mid string `json:"mid"`
}

func Fetch(data *FetchParam) (*tables.LLModel, error) {

	var item *tables.LLModel

	result := dborm.Db.
		Where(&tables.LLModel{
			Rd:  data.Rd,
			Mid: data.Mid,
		}).
		First(&item)

	if item == nil {
		item = &tables.LLModel{Mid: data.Mid}
	}

	return item, result.Error

}

// 删除模型

type DeleteParam = FetchParam

func Delete(data *DeleteParam) error {

	var item *tables.LLModel

	result := dborm.Db.
		Where(&tables.LLModel{
			Rd:  data.Rd,
			Mid: data.Mid,
		}).
		Delete(&item)

	return result.Error

}

// 获取模型列表

type FetchAllParam struct {
	Level    int32  `json:"level"`
	Family   string `json:"family"`
	Provider string `json:"provider"`
	Model    string `json:"model"`
}

func FetchAll(data *FetchAllParam) ([]*tables.LLModel, error) {

	var items []*tables.LLModel

	result := dborm.Db.
		Where(&tables.LLModel{
			Level:    data.Level,
			Family:   data.Family,
			Provider: data.Provider,
			Model:    data.Model,
		}).
		Find(&items)

	return items, result.Error

}

// 获取模型总数

type CountParam = FetchAllParam

func Count(data *CountParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&tables.LLModel{}).
		Where(&tables.LLModel{
			Level:    data.Level,
			Family:   data.Family,
			Provider: data.Provider,
			Model:    data.Model,
		}).
		Count(&count)

	return count, result.Error

}
