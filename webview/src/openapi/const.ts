interface OptionData {
    name: string;
}

export const FieldTypes: Record<string, OptionData> = {
    bool: { name: '布尔' },
    number: { name: '数字' },
    string: { name: '字符串' },
    text: { name: '多行文本' },
    lmodel: { name: 'AI 模型' },
};

export const CronjobTypes: Record<string, OptionData> = {
    TEXT: { name: '文本内容' },
    AI: { name: 'AI 生成文本' },
    BAT: { name: 'BAT 批处理' },
    POWERSHELL: { name: 'PowerShell 脚本' },
    EXEC: { name: '可执行程序' },
};

export const KeywordGroups: Record<string, OptionData> = {
    badword: { name: '违禁词' },
    command: { name: '外部指令' },
    handler: { name: '指令别名' },
};

export const BadwordLevels: Record<number, OptionData> = {
    '-1': { name: '未启用' },
    1: { name: '一般违规' },
    2: { name: '中度违规' },
    3: { name: '严重违规' },
};

export const RoomLevels: Record<number, OptionData> = {
    '-1': { name: '未限制' },
    1: { name: '待验证' },
    2: { name: '已注册' },
};

export const UserLevels: Record<number, OptionData> = {
    '-1': { name: '未限制' },
    1: { name: '待验证' },
    2: { name: '已注册' },
    7: { name: '管理员' },
    9: { name: '创始人' }
};

export const SpecialRooms: Record<string, OptionData> = {
    '*': { name: '全局' },
    '-': { name: '私聊' },
    '+': { name: '群聊全局' },
};
