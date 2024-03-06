package types

// 引用消息
// @field Msg.Content

type ReferContent struct {
	Msg struct {
		AppMsg struct {
			Action   string `json:"action"`
			Des      string `json:"des"`
			ReferMsg struct {
				Svrid string `json:"svrid"`
			} `json:"refermsg"`
			Title string `json:"title"`
			Type  string `json:"type"` // 57
			URL   string `json:"url"`
		} `json:"appmsg"`
		CommentURL   string `json:"commenturl"`
		FromUsername string `json:"fromusername"`
		Scene        string `json:"scene"`
	} `json:"msg"`
}

// 聊天记录
// @field Msg.Content

type RecordContent struct {
	Msg struct {
		AppMsg struct {
			Action     string `json:"action"`
			Des        string `json:"des"`
			RecordItem string `json:"recorditem"`
			Title      string `json:"title"`
			Type       string `json:"type"` // 19
			URL        string `json:"url"`
		} `json:"appmsg"`
		CommentURL   string `json:"commenturl"`
		FromUsername string `json:"fromusername"`
		Scene        string `json:"scene"`
	} `json:"msg"`
}

// 接收转账
// @field Msg.Content

type TransferReceiveContent struct {
	Msg struct {
		AppMsg struct {
			Action    string `json:"action"`
			Des       string `json:"des"`
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
