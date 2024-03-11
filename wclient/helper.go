package wclient

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

func TalkTop10(roomid string, day int64) []*TopItem {

	ts := TodayUnix() - 86400*day

	sql := `
		SELECT sender, COUNT(sender) AS record_count
		FROM message
		WHERE ? <= ts AND ts <= ? AND roomid = ?
		GROUP BY sender
		ORDER BY record_count DESC
		LIMIT 10
	`

	var result []*TopItem
	dborm.Db.Raw(sql, ts, ts+86400, roomid).Scan(&result)

	return result

}

func ImageTop10(roomid string, day int64) []*TopItem {

	ts := TodayUnix() - 86400*day

	sql := `
		SELECT sender, COUNT(sender) AS record_count
		FROM message
		WHERE ? <= ts AND ts <= ? AND roomid = ? AND type IN (3,47)
		GROUP BY sender
		ORDER BY record_count DESC
		LIMIT 10
	`

	var result []*TopItem
	dborm.Db.Raw(sql, ts, ts+86400, roomid).Scan(&result)

	return result

}
