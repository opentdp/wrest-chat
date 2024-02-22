package keyword

import (
	"github.com/opentdp/go-helper/dborm"

	"github.com/opentdp/wechat-rest/dbase/tables"
)

// 创建关键词

type CreateParam struct {
	Roomid string `binding:"required"`
	Phrase string `binding:"required"`
	Level  int32
}

func Create(data *CreateParam) (uint, error) {

	item := &tables.Keyword{
		Roomid: data.Roomid,
		Phrase: data.Phrase,
		Level:  data.Level,
	}

	result := dborm.Db.Create(item)

	return item.Rd, result.Error

}

// 更新关键词

type UpdateParam = CreateParam

func Update(data *UpdateParam) error {

	result := dborm.Db.
		Where(&tables.Keyword{
			Roomid: data.Roomid,
			Phrase: data.Phrase,
		}).
		Updates(tables.Keyword{
			Roomid: data.Roomid,
			Phrase: data.Phrase,
			Level:  data.Level,
		})

	return result.Error

}

// 删除关键词

type DeleteParam struct {
	Roomid string `binding:"required"`
	Phrase string `binding:"required"`
}

func Delete(data *DeleteParam) error {

	var item *tables.Keyword

	result := dborm.Db.
		Where(&tables.Keyword{
			Roomid: data.Roomid,
			Phrase: data.Phrase,
		}).
		Delete(&item)

	return result.Error

}

// 获取关键词

type FetchParam struct {
	Roomid string `binding:"required"`
	Phrase string `binding:"required"`
}

func Fetch(data *FetchParam) (*tables.Keyword, error) {

	var item *tables.Keyword

	result := dborm.Db.
		Where(&tables.Keyword{
			Roomid: data.Roomid,
			Phrase: data.Phrase,
		}).
		First(&item)

	if item == nil {
		item = &tables.Keyword{Phrase: data.Phrase}
	}

	return item, result.Error

}

// 获取关键词列表

type FetchAllParam struct {
	Roomid string
	Level  int32
}

func FetchAll(data *FetchAllParam) ([]*tables.Keyword, error) {

	var items []*tables.Keyword

	result := dborm.Db.
		Where(&tables.Keyword{
			Roomid: data.Roomid,
			Level:  data.Level,
		}).
		Find(&items)

	return items, result.Error

}

// 获取关键词总数

func Count(data *FetchAllParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Where(&tables.Keyword{
			Roomid: data.Roomid,
			Level:  data.Level,
		}).
		Count(&count)

	return count, result.Error

}
