import { Component } from '@angular/core';

import { RobotApi, ChatroomFetchAllParam, TablesChatroom } from '../../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../../openapi/wcfrest';


@Component({
    selector: 'page-bot-chatrooms',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class BotChatroomsComponent {

    public contacts: Record<string, WcfrestContactPayload> = {};

    public chatrooms: Array<TablesChatroom> = [];

    constructor() {
        this.getContacts();
        this.getChatrooms();
    }

    public getContacts() {
        WrestApi.contacts().then((data) => {
            data.forEach((item) => this.contacts[item.wxid] = item);
        });
    }

    public getChatrooms() {
        RobotApi.chatroomList({} as ChatroomFetchAllParam).then((data) => {
            this.chatrooms = data;
        });
    }

}
