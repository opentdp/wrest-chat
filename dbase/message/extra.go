package message

import (
	"time"

	"github.com/opentdp/go-helper/dborm"
	"github.com/opentdp/wrest-chat/dbase/tables"
)

func Shrink(days int) {

	var item *tables.Message

	ts := time.Now().AddDate(0, 0, -days).Unix()
	result := dborm.Db.Where("created_at < ?", ts).Delete(&item)

	if result.RowsAffected > 0 {
		dborm.Db.Exec("VACUUM")
	}

}
