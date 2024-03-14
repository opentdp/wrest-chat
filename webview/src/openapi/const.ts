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
    handler: { name: '指令别名' },
};

export const KeywordLevels: Record<number, OptionData> = {
    1: { name: '1' },
    2: { name: '2' },
    3: { name: '3' },
};

export const RoomLevels: Record<number, OptionData> = {
    1: { name: '待验证' },
    2: { name: '已注册' },
};

export const UserLevels: Record<number, OptionData> = {
    1: { name: '待验证' },
    2: { name: '已注册' },
    7: { name: '管理员' },
    9: { name: '创始人' }
};
