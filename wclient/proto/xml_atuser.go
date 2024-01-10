package proto

type AtMsgSource struct {
	AtUserList  string  `xml:"atuserlist"`
	Silence     int32   `xml:"silence"`
	MemberCount int32   `xml:"membercount"`
	Signature   string  `xml:"signature"`
	TmpNode     TmpNode `xml:"tmp_node"`
}

type TmpNode struct {
	PublisherID string `xml:",chardata"`
}
