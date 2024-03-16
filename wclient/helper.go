package wclient

import (
	"time"

	"github.com/opentdp/go-helper/dborm"
)

func TodayUnix() int64 {

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	return today.Unix()

}

// 群聊统计

type RoomData struct {
	Talk  int32 `json:"talk"`
	Image int32 `json:"image"`
}

func RoomCount(roomid string, day int64) *RoomData {

	var result = RoomData{}

	ts := TodayUnix() - 86400*day

	sql1 := `
		SELECT COUNT(*) AS record_count
		FROM message
		WHERE ? <= ts AND ts <= ? AND roomid = ?
	`
	dborm.Db.Raw(sql1, ts, ts+86400, roomid).Scan(&result.Talk)

	sql2 := `
		SELECT COUNT(*) AS record_count
		FROM message
		WHERE ? <= ts AND ts <= ? AND roomid = ? AND type IN (3,47)
	`
	dborm.Db.Raw(sql2, ts, ts+86400, roomid).Scan(&result.Image)

	return &result

}

type TopItem struct {
	Sender      string `json:"sender"`
	RecordCount int32  `json:"record_count"`
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
