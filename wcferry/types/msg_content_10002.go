package types

// 撤回消息提示
// @field Msg.Content

type RevokeContent struct {
	Type      string `xml:"type,attr"`
	RevokeMsg struct {
		Session    string `xml:"session"`
		MsgID      string `xml:"msgid"`
		NewMsgID   string `xml:"newmsgid"`
		ReplaceMsg struct {
			Text string `xml:",cdata"`
		} `xml:"replacemsg"`
		AnnouncementID struct {
			Text string `xml:",cdata"`
		} `xml:"announcement_id"`
	} `xml:"revokemsg"`
}
