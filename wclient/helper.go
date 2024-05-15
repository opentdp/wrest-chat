package wclient

import (
	"fmt"
	"strconv"
	"strings"
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

// 获取群管理信息
func ChatRoomInfo(roomid string) map[string]string {

	info := map[string]string{}

	// 获取群主
	sql1 := fmt.Sprintf(`SELECT * FROM ChatRoom WHERE ChatRoomName = '%s';`, roomid)
	res1 := wc.CmdClient.DbSqlQuery("MicroMsg.db", sql1)
	if len(res1) > 0 {
		info["owner"] = res1[0]["Reserved2"].(string)
	}

	// 获取群公告
	sql2 := fmt.Sprintf(`SELECT * FROM ChatRoomInfo WHERE ChatRoomName = '%s';`, roomid)
	res2 := wc.CmdClient.DbSqlQuery("MicroMsg.db", sql2)
	if len(res2) > 0 {
		info["announcement"] = res1[0]["Announcement"].(string)
	}

	return info

}

// 查找消息数据库
// return string 消息数据库
func FindMsgDb() string {

	maxIndex, maxDbName := -1, ""

	dbList := wc.CmdClient.GetDbNames()
	for _, dbName := range dbList {
		if strings.HasPrefix(dbName, "MSG") && strings.HasSuffix(dbName, ".db") {
			msgdb := strings.TrimSuffix(dbName, ".db")
			parts := strings.Split(msgdb, "MSG")
			if len(parts) == 2 {
				x, err := strconv.Atoi(parts[1])
				if err == nil && x > maxIndex {
					maxIndex, maxDbName = x, dbName
				}
			}
		}
	}

	return maxDbName

}
