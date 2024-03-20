package aichat

import (
	"github.com/opentdp/wrest-chat/dbase/setting"
)

var msgHistories = map[string][]*MsgHistory{}

type MsgHistory struct {
	Content string
	Role    string // user,assistant
}

func AddHistory(id, rid string, items ...*MsgHistory) {

	rd := id + "/" + rid

	if len(msgHistories[rd]) >= setting.ModelHistory {
		msgHistories[rd] = msgHistories[rd][len(items):]
	}

	msgHistories[rd] = append(msgHistories[rd], items...)

}

func GetHistory(id, rid string) []*MsgHistory {

	rd := id + "/" + rid

	if _, ok := msgHistories[rd]; !ok {
		ResetHistory(id, rid)
	}

	return msgHistories[rd]

}

func CountHistory(id, rid string) int {

	rd := id + "/" + rid

	if _, ok := msgHistories[rd]; !ok {
		ResetHistory(id, rid)
	}

	return len(msgHistories[rd])

}

func ResetHistory(id, rid string) {

	rd := id + "/" + rid

	msgHistories[rd] = []*MsgHistory{}

}
