interface LevelData {
    name: string;
}

export const RoomLevels: Record<number, LevelData> = {
    0: { name: '未注册' },
    1: { name: '黑名单' },
    2: { name: '已注册' },
};

export const UserLevels: Record<number, LevelData> = {
    0: { name: '未注册' },
    1: { name: '黑名单' },
    2: { name: '已注册' },
    7: { name: '管理员' },
    9: { name: '创始人' }
};

export const KeywordLevels: Record<number, LevelData> = {
    1: { name: '一般违规' },
    2: { name: '较为严重' },
    3: { name: '非常严重' },
};

export const FieldTypes: Record<string, LevelData> = {
    bool: { name: '布尔' },
    lmodel: { name: 'AI 模型' },
    number: { name: '数字' },
    string: { name: '字符串' },
    text: { name: '大段文字' },
};
