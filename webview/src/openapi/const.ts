interface OptionData {
    name: string;
}

interface AiOptionData extends OptionData {
    endpoint: string;
    keystyle: string;
}

export const FieldTypes: Record<string, OptionData> = {
    bool: {
        name: '布尔',
    },
    number: {
        name: '数字',
    },
    string: {
        name: '字符串',
    },
    text: {
        name: '多行文本',
    },
    lmodel: {
        name: 'AI 模型',
    },
};

export const AiModels: Record<string, AiOptionData> = {
    aliyun: {
        name: '阿里通义千问',
        endpoint: 'https://dashscope.aliyuncs.com',
        keystyle: '阿里百炼（通义千问）填写 APP-ID,AGENT-KEY,ACCESS_KEY_ID,ACCESS_KEY_SECRET',
    },
    baidu: {
        name: '百度文心',
        endpoint: 'https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop',
        keystyle: '文心一言填写 API-KEY,API-SECRET',
    },
    google: {
        name: 'Google Gemini',
        endpoint: 'https://generativelanguage.googleapis.com',
        keystyle: '',
    },
    openai: {
        name: 'OpenAI GPT',
        endpoint: 'https://api.openai.com/v1',
        keystyle: '',
    },
    tencent: {
        name: '腾讯混元',
        endpoint: 'https://hunyuan.cloud.tencent.com/hyllm/v1',
        keystyle: '腾讯混元填写 APP-ID,API-KEY,API-SECRET',
    },
    xunfei: {
        name: '科大讯飞',
        endpoint: 'wss://spark-api.xf-yun.com',
        keystyle: '科大讯飞填写 APP-ID,API-KEY,API-SECRET',
    },
};

export const CronjobTypes: Record<string, OptionData> = {
    TEXT: {
        name: '文本内容',
    },
    AI: {
        name: 'AI 生成文本',
    },
    BAT: {
        name: 'BAT 批处理',
    },
    POWERSHELL: {
        name: 'PowerShell 脚本',
    },
    EXEC: {
        name: '可执行程序',
    },
};

export const KeywordGroups: Record<string, OptionData> = {
    badword: {
        name: '违禁词',
    },
    handler: {
        name: '指令别名',
    },
    command: {
        name: '外部指令',
    },
    imagefn: {
        name: '图片处理',
    },
};

export const BadwordLevels: Record<number, OptionData> = {
    '-1': {
        name: '未启用',
    },
    1: {
        name: '一般违规',
    },
    2: {
        name: '中度违规',
    },
    3: {
        name: '严重违规',
    },
};

export const RoomLevels: Record<number, OptionData> = {
    '-1': {
        name: '未注册',
    },
    1: {
        name: '待验证',
    },
    2: {
        name: '已注册',
    },
};

export const UserLevels: Record<number, OptionData> = {
    '-1': {
        name: '未注册',
    },
    1: {
        name: '待验证',
    },
    2: {
        name: '已注册',
    },
    7: {
        name: '管理员',
    },
    9: {
        name: '创始人',
    }
};

export const SpecialRooms: Record<string, OptionData> = {
    '*': {
        name: '全局',
    },
    '-': {
        name: '私聊',
    },
    '+': {
        name: '群聊全局',
    },
};
