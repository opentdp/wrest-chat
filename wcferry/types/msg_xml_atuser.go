package types

// 群里At用户消息
// @msg.type 1

type MsgXmlAtUser struct {
	AtUserList  string `xml:"atuserlist"`
	Silence     int32  `xml:"silence"`
	MemberCount int32  `xml:"membercount"`
	Signature   string `xml:"signature"`
	TmpNode     struct {
		PublisherID string `xml:",chardata"`
	} `xml:"tmp_node"`
}
