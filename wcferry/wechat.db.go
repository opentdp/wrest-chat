package wcferry

// 头像表结构
type ContactHeadImgUrlTable struct {
	BigHeadImgUrl   string `json:"bigHeadImgUrl,omitempty"`
	HeadImgMd5      string `json:"headImgMd5,omitempty"`
	Reverse0        int    `json:"reverse0,omitempty"`
	Reverse1        any    `json:"reverse1,omitempty"`
	SmallHeadImgUrl string `json:"smallHeadImgUrl,omitempty"`
	UsrName         string `json:"usrName,omitempty"`
}
