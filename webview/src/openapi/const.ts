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