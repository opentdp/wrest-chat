package llmodel

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wechat-rest/dbase/tables"
)

// 创建模型

type CreateParam struct {
	Mid      string `binding:"required"`
	Provider string `binding:"required"`
	Endpoint string
	Family   string `binding:"required"`
	Model    string `binding:"required"`
	Secret   string `binding:"required"`
}

func Create(data *CreateParam) (uint, error) {

	item := &tables.LLModel{
		Mid:      data.Mid,
		Provider: data.Provider,
		Endpoint: data.Endpoint,
		Family:   data.Family,
		Model:    data.Model,
		Secret:   data.Secret,
	}

	result := dborm.Db.Create(item)

	return item.Rd, result.Error

}

// 更新模型

type UpdateParam = CreateParam

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&tables.LLModel{
			Mid: data.Mid,
		}).
		Updates(tables.LLModel{
			Provider: data.Provider,
			Endpoint: data.Endpoint,
			Family:   data.Family,
			Model:    data.Model,
			Secret:   data.Secret,
		})

	return result.Error

}

// 合并模型

type MigrateParam = CreateParam

func Migrate(data *MigrateParam) error {

	item, err := Fetch(&FetchParam{
		Mid: data.Mid,
	})

	if err == nil && item.Rd > 0 {
		err = Update(data)
	} else {
		_, err = Create(data)
	}

	return err

}

// 获取模型

type FetchParam struct {
	Mid string `binding:"required"`
}

func Fetch(data *FetchParam) (*tables.LLModel, error) {

	var item *tables.LLModel

	result := dborm.Db.
		Where(&tables.LLModel{
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
			Mid: data.Mid,
		}).
		Delete(&item)

	return result.Error

}

// 获取模型列表

type FetchAllParam struct {
	Provider string
	Family   string
	Model    string
}

func FetchAll(data *FetchAllParam) ([]*tables.LLModel, error) {

	var items []*tables.LLModel

	result := dborm.Db.
		Where(&tables.LLModel{
			Provider: data.Provider,
			Family:   data.Family,
			Model:    data.Model,
		}).
		Find(&items)

	return items, result.Error

}

// 获取模型总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Where(&tables.LLModel{
			Provider: data.Provider,
			Family:   data.Family,
			Model:    data.Model,
		}).
		Count(&count)

	return count, result.Error

}
