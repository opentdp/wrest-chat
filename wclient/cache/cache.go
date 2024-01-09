package cache

type HistoryItem struct {
	Role    string
	Content string
}

var Models = make(map[string]string)
var History = make(map[string][]HistoryItem)

var Handlers = make(map[string]func(id, msg string) string)
