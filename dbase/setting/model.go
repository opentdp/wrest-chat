package setting

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wrest-chat/dbase/tables"
)

// 创建配置

type CreateParam struct {
	Rd     uint   `json:"rd"`
	Name   string `json:"name" binding:"required"`
	Type   string `json:"type"`
	Group  string `json:"group"`
	Value  string `json:"value"`
	Title  string `json:"title"`
	Remark string `json:"remark"`
}

func Create(data *CreateParam) (uint, error) {

	item := &tables.Setting{
		Name:   data.Name,
		Type:   data.Type,
		Group:  data.Group,
		Value:  data.Value,
		Title:  data.Title,
		Remark: data.Remark,
	}

	result := dborm.Db.Create(item)

	return item.Rd, result.Error

}

// 更新配置

type UpdateParam = CreateParam

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&tables.Setting{
			Rd: data.Rd,
		}).
		Updates(tables.Setting{
			Name:   data.Name,
			Type:   data.Type,
			Group:  data.Group,
			Value:  data.Value,
			Title:  data.Title,
			Remark: data.Remark,
		})

	return result.Error

}

// 合并配置

type ReplaceParam = CreateParam

func Replace(data *ReplaceParam) error {

	rq := &FetchParam{Rd: data.Rd}
	if rq.Rd == 0 {
		rq.Name = data.Name
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

// 获取配置

type FetchParam struct {
	Rd   uint   `json:"rd"`
	Name string `json:"name"`
}

func Fetch(data *FetchParam) (*tables.Setting, error) {

	var item *tables.Setting

	result := dborm.Db.
		Where(&tables.Setting{
			Rd:   data.Rd,
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
			Rd:   data.Rd,
			Name: data.Name,
		}).
		Delete(&item)

	return result.Error

}

// 获取配置列表

type FetchAllParam struct {
	Group string `json:"group"`
}

func FetchAll(data *FetchAllParam) ([]*tables.Setting, error) {

	var items []*tables.Setting

	result := dborm.Db.
		Where(&tables.Setting{
			Group: data.Group,
		}).
		Find(&items)

	return items, result.Error

}

// 获取配置总数

type CountParam = FetchAllParam

func Count(data *CountParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&tables.Setting{}).
		Where(&tables.Setting{
			Group: data.Group,
		}).
		Count(&count)

	return count, result.Error

}
