import { Component } from '@angular/core';

import { LevelData } from 'src/openapi/const';
import { RobotApi, TablesChatroom, ChatroomFetchAllParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-chatroom-list',
    templateUrl: 'list.html',
    styleUrls: ['list.scss']
})
export class ChatroomListComponent {

    public levels = LevelData;

    public avatars: Record<string, string> = {};
    public contacts: Record<string, WcfrestContactPayload> = {};

    public chatrooms: Array<TablesChatroom> = [];

    constructor() {
        this.getChatrooms();
        this.getContacts();
    }

    public getChatrooms() {
        const rq = {} as ChatroomFetchAllParam;
        RobotApi.chatroomList(rq).then((data) => {
            this.chatrooms = data || [];
        });
    }

    public deleteChatroom(item: TablesChatroom) {
        RobotApi.chatroomDelete({ roomid: item.roomid }).then(() => {
            this.getChatrooms();
        });
    }

    public getContacts() {
        WrestApi.contacts().then((data) => {
            data.forEach((item) => this.contacts[item.wxid] = item);
        });
    }

    public getAvatars(ids: string[]) {
        WrestApi.avatars({ wxids: [...new Set(ids)] }).then((data) => {
            data && data.forEach((item) => {
                this.avatars[item.usr_name] = item.small_head_img_url;
            });
        });
    }

}
