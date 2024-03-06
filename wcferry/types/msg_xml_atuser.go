package types

// 群里At用户消息
// @type 1
// @field Msg.Xml

type AtMsgSource struct {
	AtUserList  string `xml:"atuserlist"`
	Silence     int32  `xml:"silence"`
	MemberCount int32  `xml:"membercount"`
	Signature   string `xml:"signature"`
	TmpNode     struct {
		PublisherID string `xml:",chardata"`
	} `xml:"tmp_node"`
}
