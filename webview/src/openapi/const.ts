interface OptionData {
    name: string;
}

interface AiOptionData extends OptionData {
    endpoint: string;
    keystyle: string;
    models: Record<string, string>;
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
        name: '阿里灵积（通义千问）',
        endpoint: 'https://dashscope.aliyuncs.com',
        keystyle: '密钥格式 APP-ID,AGENT-KEY,ACCESS_KEY_ID,ACCESS_KEY_SECRET',
        models: {
            'qwen-max': 'qwen-max',
            'qwen-plus': 'qwen-plus',
        },
    },
    baidu: {
        name: '百度千帆（文心)',
        endpoint: 'https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop',
        keystyle: '密钥格式 API-KEY,API-SECRET',
        models: {
            'completions_pro': 'ERNIE 4.0',
            'completions': 'ERNIE-3.5-8K',
        },
    },
    google: {
        name: 'Google Gemini',
        endpoint: 'https://generativelanguage.googleapis.com',
        keystyle: '',
        models: {
            'gemini-pro': 'gemini-pro',
        },
    },
    openai: {
        name: 'OpenAI ChatGPT',
        endpoint: 'https://api.openai.com/v1',
        keystyle: '',
        models: {
            'gpt-4': 'gpt-4',
            'gpt-4-32k': 'gpt-4-32k',
            'gpt-3.5-turbo': 'gpt-3.5-turbo',
        },
    },
    tencent: {
        name: '腾讯（混元）',
        endpoint: 'https://hunyuan.cloud.tencent.com/hyllm/v1',
        keystyle: '密钥格式 APP-ID,API-KEY,API-SECRET',
        models: {},
    },
    xunfei: {
        name: '科大讯飞（星火）',
        endpoint: 'wss://spark-api.xf-yun.com',
        keystyle: '密钥格式 APP-ID,API-KEY,API-SECRET',
        models: {
            'v3.5': '星火 v3.5',
            'v3': '星火 v3',
        },
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
