package types

// 共享实时位置、文件、转账、链接、群邀请等
type MsgContent49 struct {
	AppMsg struct {
		Action     string    `xml:"action"`
		AppAttach  AppAttach `xml:"appattach"` // 6
		Des        string    `xml:"des"`
		RecordItem string    `xml:"recorditem"` // 19
		ReferMsg   ReferMsg  `xml:"refermsg"`   // 57
		Title      string    `xml:"title"`
		Type       int32     `xml:"type"`
		URL        string    `xml:"url"`
		WCPayInfo  WCPayInfo `xml:"wcpayinfo"` // 1,3
	} `xml:"appmsg"`
	CommentURL   string `xml:"commenturl"`
	FromUsername string `xml:"fromusername"`
	Scene        int32  `xml:"scene"`
}

// 接收转账 type=1:转账,3:已收款

type WCPayInfo struct {
	BeginTransferTime string `xml:"begintransfertime"`
	EffectiveDate     string `xml:"effectivedate"`
	FeeDesc           string `xml:"feedesc"`
	InvalidTime       string `xml:"invalidtime"`
	PayMemo           string `xml:"pay_memo"`
	PayerUsername     string `xml:"payer_username"`
	PaySubtype        string `xml:"paysubtype"`
	ReceiverUsername  string `xml:"receiver_username"`
	TranscationID     string `xml:"transcationid"`
	TransferID        string `xml:"transferid"`
}

// 文件 type=6

type AppAttach struct {
	AesKey            string `xml:"aeskey"`
	AttachID          string `xml:"attachid"`
	CDNAttachURL      string `xml:"cdnattachurl"`
	EmoticonMD5       string `xml:"emoticonmd5"`
	EncryVer          string `xml:"encryver"`
	FileExt           string `xml:"fileext"`
	FileKey           string `xml:"filekey"`
	FileUploadToken   string `xml:"fileuploadtoken"`
	OverwriteNewMsgID string `xml:"overwrite_newmsgid"`
	TotalLen          string `xml:"totallen"`
}

// 聊天记录 type=19
// 先 html.UnescapeString(RecordItem)，再解析

type RecordInfo struct {
	Title       string `xml:"title"`
	Description string `xml:"desc"`
	DataList    struct {
		Count     int `xml:"count,attr"`
		DataItems []struct {
			DataID         string `xml:"dataid,attr"`
			DataType       int    `xml:"datatype,attr"`
			DataSourceID   int64  `xml:"datasourceid,attr"`
			CDNEncryver    int    `xml:"cdnencryver"`
			DataDesc       string `xml:"datadesc"`
			SourceName     string `xml:"sourcename"`
			SourceTime     string `xml:"sourcetime"`
			SourceHeadURL  string `xml:"sourceheadurl"`
			FromNewMsgID   int64  `xml:"fromnewmsgid"`
			DataItemSource struct {
				MsgID        int64  `xml:"msgid"`
				CreateTime   int64  `xml:"createtime"`
				HashUsername string `xml:"hashusername"`
			} `xml:"dataitemsource"`
		} `xml:"dataitem"`
	} `xml:"datalist"`
	FavUsername   string `xml:"favusername"`
	FavCreateTime int64  `xml:"favcreatetime"`
}

// 聊天记录 type=40
// 不完整的聊天记录，可能关联了其它表记录

// 引用 type=57

type ReferMsg struct {
	Svrid uint64 `xml:"svrid"`
	Type  uint32 `xml:"type"`
}
