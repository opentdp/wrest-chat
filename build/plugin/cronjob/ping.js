// 这是一个计划任务插件示例，用于测试计划任务能否正常工作
// 插件开发语言不限于 node.js，只要添加下列注释参数，并设置正确的 @Content 作为解析器即可

// @Name: Ping
// @Second: 0
// @Minute: 0
// @Hour: 0
// @DayOfMonth: *
// @Month: *
// @DayOfWeek: *
// @Timeout: 300
// @Content: node.exe
// @Deliver: wechat,xxx@room,wxid_xxxx

console.log('pong');
