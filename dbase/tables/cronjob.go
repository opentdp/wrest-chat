package tables

// 计划任务

type Cronjob struct {
	Rd         uint   `json:"rd" gorm:"primaryKey"` // 主键
	Name       string `json:"name"`                 // 名称
	Second     string `json:"second"`               // 秒
	Minute     string `json:"minute"`               // 分
	Hour       string `json:"hour"`                 // 时
	DayOfMonth string `json:"day_of_month"`         // 日
	Month      string `json:"month"`                // 月
	DayOfWeek  string `json:"day_of_week"`          // 周
	Type       string `json:"type"`                 // 命令类型 [CMD, POWERSHELL, SHELL]
	Timeout    uint   `json:"timeout"`              // 超时时间（秒）
	Directory  string `json:"directory"`            // 工作目录
	Content    string `json:"content"`              // 命令内容
	Deliver    string `json:"deliver"`              // 执行结果交付方式
	EntryId    int64  `json:"entry_id"`             // 任务运行时 Id
	CreatedAt  int64  `json:"created_at"`           // 创建时间戳
	UpdatedAt  int64  `json:"updated_at"`           // 最后更新时间戳
}
