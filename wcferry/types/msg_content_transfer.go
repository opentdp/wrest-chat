package types

// 收到转账消息
// @type 49
// @field Msg.Content

type TransferReceiveMsg struct {
	Msg struct {
		AppMsg struct {
			Action    string `json:"action"`
			AppID     string `json:"appid"`
			Content   string `json:"content"`
			Des       string `json:"des"`
			ExtInfo   string `json:"extinfo"`
			LowURL    string `json:"lowurl"`
			SDKVer    string `json:"sdkver"`
			ThumbURL  string `json:"thumburl"`
			Title     string `json:"title"`
			Type      string `json:"type"` // 1:转账 3:已收款
			URL       string `json:"url"`
			WCPayInfo struct {
				BeginTransferTime string `json:"begintransfertime"`
				EffectiveDate     string `json:"effectivedate"`
				FeeDesc           string `json:"feedesc"`
				InvalidTime       string `json:"invalidtime"`
				PayMemo           string `json:"pay_memo"`
				PayerUsername     string `json:"payer_username"`
				PaySubtype        string `json:"paysubtype"`
				ReceiverUsername  string `json:"receiver_username"`
				TranscationID     string `json:"transcationid"`
				TransferID        string `json:"transferid"`
			} `json:"wcpayinfo"`
		} `json:"appmsg"`
	} `json:"msg"`
}
