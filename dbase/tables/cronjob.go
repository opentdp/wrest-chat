package tables

// 计划任务

type Cronjob struct {
	Rd         uint   `gorm:"primaryKey" json:"rd"` // 主键
	Name       string `json:"name"`                 // 名称
	Second     string `json:"second"`               // 秒
	Minute     string `json:"minute"`               // 分
	Hour       string `json:"hour"`                 // 时
	DayofMonth string `json:"dayof_month"`          // 日
	Month      string `json:"month"`                // 月
	DayofWeek  string `json:"dayof_week"`           // 周
	Type       string `json:"type"`                 // 命令类型，[CMD, POWERSHELL, SHELL]
	Directory  string `json:"directory"`            // 工作目录
	Timeout    uint   `json:"timeout"`              // 超时时间（秒）
	Content    string `json:"content"`              // 命令内容
	EntryId    int64  `json:"entry_id"`             // 当前计划 Id
	CreatedAt  int64  `json:"created_at"`           // 创建时间戳
	UpdatedAt  int64  `json:"updated_at"`           // 最后更新时间戳
}
