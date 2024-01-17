package types

type SysMsg struct {
	Type      string    `xml:"type,attr"`
	RevokeMsg RevokeMsg `xml:"revokemsg"`
}

type RevokeMsg struct {
	Session        string `xml:"session"`
	MsgID          string `xml:"msgid"`
	NewMsgID       string `xml:"newmsgid"`
	ReplaceMsg     CDATA  `xml:"replacemsg"`
	AnnouncementID CDATA  `xml:"announcement_id"`
}

type CDATA struct {
	Text string `xml:",cdata"`
}
