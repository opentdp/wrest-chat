package cronjob

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wechat-rest/dbase/tables"
)

// 创建计划

type CreateParam struct {
	Rd         uint   `json:"rd"`
	Name       string `binding:"required" json:"name"`
	Second     string `binding:"required" json:"second"`
	Minute     string `binding:"required" json:"minute"`
	Hour       string `binding:"required" json:"hour"`
	DayofMonth string `binding:"required" json:"dayof_month"`
	Month      string `binding:"required" json:"month"`
	DayofWeek  string `binding:"required" json:"dayof_week"`
	Type       string `binding:"required" json:"type"`
	Directory  string `binding:"required" json:"directory"`
	Timeout    uint   `binding:"required" json:"timeout"`
	Content    string `binding:"required" json:"content"`
	EntryId    int64  `json:"entry_id"`
}

func Create(data *CreateParam) (uint, error) {

	item := &tables.Cronjob{
		Name:       data.Name,
		Second:     data.Second,
		Minute:     data.Minute,
		Hour:       data.Hour,
		DayofMonth: data.DayofMonth,
		Month:      data.Month,
		DayofWeek:  data.DayofWeek,
		Type:       data.Type,
		Directory:  data.Directory,
		Timeout:    data.Timeout,
		Content:    data.Content,
		EntryId:    data.EntryId,
	}

	result := dborm.Db.Create(item)

	return item.Rd, result.Error

}

// 更新计划

type UpdateParam = CreateParam

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&tables.Cronjob{
			Rd: data.Rd,
		}).
		Updates(tables.Cronjob{
			Name:       data.Name,
			Second:     data.Second,
			Minute:     data.Minute,
			Hour:       data.Hour,
			DayofMonth: data.DayofMonth,
			Month:      data.Month,
			DayofWeek:  data.DayofWeek,
			Type:       data.Type,
			Directory:  data.Directory,
			Timeout:    data.Timeout,
			Content:    data.Content,
			EntryId:    data.EntryId,
		})

	return result.Error

}

// 合并联系人

type ReplaceParam = CreateParam

func Replace(data *ReplaceParam) error {

	item, err := Fetch(&FetchParam{
		Rd: data.Rd,
	})

	if err == nil && item.Rd > 0 {
		err = Update(data)
	} else {
		_, err = Create(data)
	}

	return err

}

// 获取计划

type FetchParam struct {
	Rd      uint  `json:"rd"`
	EntryId int64 `json:"entry_id"`
}

func Fetch(data *FetchParam) (*tables.Cronjob, error) {

	var item *tables.Cronjob

	result := dborm.Db.
		Where(&tables.Cronjob{
			Rd:      data.Rd,
			EntryId: data.EntryId,
		}).
		First(&item)

	return item, result.Error

}

// 删除计划

type DeleteParam = FetchParam

func Delete(data *DeleteParam) error {

	result := dborm.Db.
		Where(&tables.Cronjob{
			Rd: data.Rd,
		}).
		Delete(&tables.Cronjob{})

	return result.Error

}

// 获取计划列表

type FetchAllParam struct {
	Type string `json:"type"`
}

func FetchAll(data *FetchAllParam) ([]*tables.Cronjob, error) {

	var items []*tables.Cronjob

	result := dborm.Db.
		Where(&tables.Cronjob{
			Type: data.Type,
		}).
		Find(&items)

	return items, result.Error

}

// 获取计划总数

type CountParam = FetchAllParam

func Count(data *CountParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&tables.Cronjob{}).
		Where(&tables.Cronjob{
			Type: data.Type,
		}).
		Count(&count)

	return count, result.Error

}
