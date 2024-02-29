import { Component } from '@angular/core';

import { RoomLevels } from 'src/openapi/const';
import { RobotApi, TablesChatroom } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-chatroom-list',
    templateUrl: 'list.html',
    styleUrls: ['list.scss']
})
export class ChatroomListComponent {

    public roomLevels = RoomLevels;

    public wcfAvatars: Record<string, string> = {};
    public wcfChatrooms: Record<string, WcfrestContactPayload> = {};

    public chatrooms: Array<TablesChatroom> = [];

    constructor() {
        this.getChatrooms();
        this.getWcfChatrooms();
    }

    public getChatrooms() {
        RobotApi.chatroomList({}).then((data) => {
            this.chatrooms = data || [];
        });
    }

    public deleteChatroom(item: TablesChatroom) {
        RobotApi.chatroomDelete({ roomid: item.roomid }).then(() => {
            this.getChatrooms();
        });
    }

    public getWcfAvatars(ids: string[]) {
        WrestApi.avatars({ wxids: [...new Set(ids)] }).then((data) => {
            data && data.forEach((item) => {
                this.wcfAvatars[item.usr_name] = item.small_head_img_url;
            });
        });
    }

    public getWcfChatrooms() {
        WrestApi.chatrooms().then((data) => {
            data.forEach((item) => this.wcfChatrooms[item.wxid] = item);
        });
    }

}
