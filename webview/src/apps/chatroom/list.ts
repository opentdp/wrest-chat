import { Component } from '@angular/core';

import { RoomLevels } from '../../openapi/const';
import { RobotApi, TablesChatroom } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-chatroom-list',
    templateUrl: 'list.html'
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
        return RobotApi.chatroomList({}).then((data) => {
            this.chatrooms = data || [];
        });
    }

    public deleteChatroom(item: TablesChatroom) {
        return RobotApi.chatroomDelete({ rd: item.rd }).then(() => {
            this.getChatrooms();
        });
    }

    public getWcfChatrooms() {
        return WrestApi.chatrooms().then((data) => {
            data.forEach((item) => this.wcfChatrooms[item.wxid] = item);
        });
    }

}
