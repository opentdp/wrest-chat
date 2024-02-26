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

type UpdateParam struct {
	Rd     uint `binding:"required"`
	Roomid string
	Phrase string
	Level  int32
}

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

// 合并关键词

type MigrateParam = CreateParam

func Migrate(data *MigrateParam) error {

	item, err := Fetch(&FetchParam{
		Roomid: data.Roomid,
		Phrase: data.Phrase,
	})

	if err == nil && item.Rd > 0 {
		err = Update(&UpdateParam{
			Rd:     item.Rd,
			Roomid: data.Roomid,
			Phrase: data.Phrase,
			Level:  data.Level,
		})
	} else {
		_, err = Create(data)
	}

	return err

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

// 删除关键词

type DeleteParam = FetchParam

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

type CountParam = FetchAllParam

func Count(data *CountParam) (int64, error) {

	var count int64

	result := dborm.Db.
		Model(&tables.Keyword{}).
		Where(&tables.Keyword{
			Roomid: data.Roomid,
			Level:  data.Level,
		}).
		Count(&count)

	return count, result.Error

}
