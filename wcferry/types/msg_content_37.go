package types

// 好友请求消息

type MsgContent37 struct {
	FromUserName      string `xml:"fromusername,attr"`
	EncryptUserName   string `xml:"encryptusername,attr"`
	FromNickName      string `xml:"fromnickname,attr"`
	Content           string `xml:"content,attr"`
	FullPY            string `xml:"fullpy,attr"`
	ShortPY           string `xml:"shortpy,attr"`
	ImageStatus       int32  `xml:"imagestatus,attr"`
	Scene             int32  `xml:"scene,attr"`
	Country           string `xml:"country,attr"`
	Province          string `xml:"province,attr"`
	City              string `xml:"city,attr"`
	Sign              string `xml:"sign,attr"`
	PerCard           int32  `xml:"percard,attr"`
	Sex               int32  `xml:"sex,attr"`
	Alias             string `xml:"alias,attr"`
	Weibo             string `xml:"weibo,attr"`
	AlbumFlag         int32  `xml:"albumflag,attr"`
	AlbumStyle        int32  `xml:"albumstyle,attr"`
	AlbumBgImgID      string `xml:"albumbgimgid,attr"`
	SnsFlag           int32  `xml:"snsflag,attr"`
	SnsBgImgID        string `xml:"snsbgimgid,attr"`
	SnsBgObjectID     string `xml:"snsbgobjectid,attr"`
	MHash             string `xml:"mhash,attr"`
	MFullHash         string `xml:"mfullhash,attr"`
	BigHeadImgURL     string `xml:"bigheadimgurl,attr"`
	SmallHeadImgURL   string `xml:"smallheadimgurl,attr"`
	Ticket            string `xml:"ticket,attr"`
	OpCode            int32  `xml:"opcode,attr"`
	GoogleContact     string `xml:"googlecontact,attr"`
	QRTicket          string `xml:"qrticket,attr"`
	ChatRoomUserName  string `xml:"chatroomusername,attr"`
	SourceUserName    string `xml:"sourceusername,attr"`
	SourceNickName    string `xml:"sourcenickname,attr"`
	ShareCardUserName string `xml:"sharecardusername,attr"`
	ShareCardNickName string `xml:"sharecardnickname,attr"`
	CardVersion       int32  `xml:"cardversion,attr"`
	ExtFlag           int32  `xml:"extflag,attr"`
	BrandList         struct {
		Count int32 `xml:"count,attr"`
		Ver   int32 `xml:"ver,attr"`
	} `xml:"brandlist"`
}
