package keyword

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wechat-rest/dbase/tables"
)

// 创建配置

type CreateParam struct {
	Name   string `binding:"required" json:"name"`
	Value  string `json:"value"`
	Title  string `binding:"required" json:"title"`
	Remark string `json:"remark"`
}

func Create(data *CreateParam) (uint, error) {

	item := &tables.Setting{
		Name:   data.Name,
		Value:  data.Value,
		Title:  data.Title,
		Remark: data.Remark,
	}

	result := dborm.Db.Create(item)

	return item.Rd, result.Error

}

// 更新配置

type UpdateParam struct {
	Name   string `binding:"required" json:"name"`
	Value  string `json:"value"`
	Title  string `json:"title"`
	Remark string `json:"remark"`
}

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&tables.Setting{
			Name: data.Name,
		}).
		Updates(tables.Setting{
			Value:  data.Value,
			Title:  data.Title,
			Remark: data.Remark,
		})

	return result.Error

}

// 合并配置

type MigrateParam = CreateParam

func Migrate(data *MigrateParam) error {

	item, err := Fetch(&FetchParam{
		Name: data.Name,
	})

	if err == nil && item.Rd > 0 {
		err = Update(&UpdateParam{
			Name:   data.Name,
			Value:  data.Value,
			Title:  data.Title,
			Remark: data.Remark,
		})
	} else {
		_, err = Create(data)
	}

	return err

}

// 获取配置

type FetchParam struct {
	Name string `binding:"required" json:"name"`
}

func Fetch(data *FetchParam) (*tables.Setting, error) {

	var item *tables.Setting

	result := dborm.Db.
		Where(&tables.Setting{
			Name: data.Name,
		}).
		First(&item)

	if item == nil {
		item = &tables.Setting{Name: data.Name}
	}

	return item, result.Error

}

// 删除配置

type DeleteParam = FetchParam

func Delete(data *DeleteParam) error {

	var item *tables.Setting

	result := dborm.Db.
		Where(&tables.Setting{
			Name: data.Name,
		}).
		Delete(&item)

	return result.Error

}

// 获取配置列表

type FetchAllParam struct{}

func FetchAll(data *FetchAllParam) ([]*tables.Setting, error) {

	var items []*tables.Setting

	result := dborm.Db.
		Find(&items)

	return items, result.Error

}

// 获取配置总数

type CountParam = FetchAllParam

func Count(data *CountParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&tables.Setting{}).
		Count(&count)

	return count, result.Error

}
