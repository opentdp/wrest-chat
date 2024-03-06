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

// 引用 type=57

type ReferMsg struct {
	Svrid string `xml:"svrid"`
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
