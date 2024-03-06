package message

import (
	"time"

	"github.com/opentdp/go-helper/dborm"
)

type TopItem struct {
	Sender      string `json:"sender"`
	RecordCount int32  `json:"record_count"`
}

func TodayUnix() int64 {

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	return today.Unix()

}

func TalkTop10(roomid string) []*TopItem {

	var result []*TopItem

	sql := `
		SELECT sender, COUNT(sender) AS record_count
		FROM message
		WHERE created_at >= ? AND roomid = ?
		GROUP BY sender
		ORDER BY record_count DESC
		LIMIT 10
	`

	timestamp := TodayUnix()
	dborm.Db.Raw(sql, timestamp, roomid).Scan(&result)

	return result

}

func ImageTop10(roomid string) []*TopItem {

	var result []*TopItem

	sql := `
		SELECT sender, COUNT(sender) AS record_count
		FROM message
		WHERE created_at >= ? AND roomid = ? AND type IN (3,47)
		GROUP BY sender
		ORDER BY record_count DESC
		LIMIT 10
	`

	timestamp := TodayUnix()
	dborm.Db.Raw(sql, timestamp, roomid).Scan(&result)

	return result

}
