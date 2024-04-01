package cronjob

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wrest-chat/dbase/tables"
)

// 创建计划

type CreateParam struct {
	Rd         uint   `json:"rd"`
	Name       string `json:"name" binding:"required"`
	Second     string `json:"second" binding:"required"`
	Minute     string `json:"minute" binding:"required"`
	Hour       string `json:"hour" binding:"required"`
	DayOfMonth string `json:"day_of_month" binding:"required"`
	Month      string `json:"month" binding:"required"`
	DayOfWeek  string `json:"day_of_week" binding:"required"`
	Type       string `json:"type" binding:"required"`
	Timeout    uint   `json:"timeout" binding:"required"`
	Directory  string `json:"directory" binding:"required"`
	Content    string `json:"content" binding:"required"`
	Deliver    string `json:"deliver"`
	EntryId    int64  `json:"entry_id"`
}

func Create(data *CreateParam) (uint, error) {

	item := &tables.Cronjob{
		Name:       data.Name,
		Second:     data.Second,
		Minute:     data.Minute,
		Hour:       data.Hour,
		DayOfMonth: data.DayOfMonth,
		Month:      data.Month,
		DayOfWeek:  data.DayOfWeek,
		Type:       data.Type,
		Timeout:    data.Timeout,
		Directory:  data.Directory,
		Content:    data.Content,
		Deliver:    data.Deliver,
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
			DayOfMonth: data.DayOfMonth,
			Month:      data.Month,
			DayOfWeek:  data.DayOfWeek,
			Type:       data.Type,
			Timeout:    data.Timeout,
			Directory:  data.Directory,
			Content:    data.Content,
			Deliver:    data.Deliver,
			EntryId:    data.EntryId,
		})

	return result.Error

}

// 合并计划

type ReplaceParam = CreateParam

func Replace(data *ReplaceParam) error {

	rq := &FetchParam{
		Rd: data.Rd,
	}

	item, err := Fetch(rq)
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
