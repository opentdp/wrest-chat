package cache

var Models = make(map[string]string)
var History = make(map[string][]string)

var Handlers = make(map[string]func(id, msg string) string)
