package wcf

// 执行结果
type ActionStatus struct {
	Error   error `json:"error"`
	Success bool  `json:"success"`
}

// 数据库查询参数
type DbSqlQueryRequest struct {
	Db  string `json:"db"`
	Sql string `json:"sql"`
}

// 消息转发参数
type ForwardMsgRequest struct {
	Url string `json:"url"`
}
