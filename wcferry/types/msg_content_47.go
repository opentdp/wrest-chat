package types

type MsgContent47 struct {
	Emoji Emoji `xml:"emoji"`
}

type Emoji struct {
	FromUsername      string `xml:"fromusername,attr"`
	ToUsername        string `xml:"tousername,attr"`
	Type              int    `xml:"type,attr"`
	IDBuffer          string `xml:"idbuffer,attr"`
	MD5               string `xml:"md5,attr"`
	Len               int    `xml:"len,attr"`
	ProductID         string `xml:"productid,attr"`
	AndroidMD5        string `xml:"androidmd5,attr"`
	AndroidLen        int    `xml:"androidlen,attr"`
	S60v3MD5          string `xml:"s60v3md5,attr"`
	S60v3Len          int    `xml:"s60v3len,attr"`
	S60v5MD5          string `xml:"s60v5md5,attr"`
	S60v5Len          int    `xml:"s60v5len,attr"`
	CDNURL            string `xml:"cdnurl,attr"`
	DesignerID        string `xml:"designerid,attr"`
	ThumbURL          string `xml:"thumburl,attr"`
	EncryptURL        string `xml:"encrypturl,attr"`
	AESKey            string `xml:"aeskey,attr"`
	ExternURL         string `xml:"externurl,attr"`
	ExternMD5         string `xml:"externmd5,attr"`
	Width             int    `xml:"width,attr"`
	Height            int    `xml:"height,attr"`
	TPURL             string `xml:"tpurl,attr"`
	TPAuthKey         string `xml:"tpauthkey,attr"`
	AttachedText      string `xml:"attachedtext,attr"`
	AttachedTextColor string `xml:"attachedtextcolor,attr"`
	LensID            string `xml:"lensid,attr"`
	EmojiAttr         string `xml:"emojiattr,attr"`
	LinkID            string `xml:"linkid,attr"`
	Desc              string `xml:"desc,attr"`
}
