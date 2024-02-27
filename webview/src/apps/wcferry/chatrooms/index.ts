import { Component } from '@angular/core';

import { WrestApi, WcfrestContactPayload } from '../../../openapi/wcfrest';


@Component({
    selector: 'page-wcf-chatrooms',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class WcfChatroomsComponent {

    public chatrooms: Array<WcfrestContactPayload> = [];

    public selected!: WcfrestContactPayload;
    public members: Array<WcfrestContactPayload> = [];

    constructor() {
        this.getChatrooms();
    }

    public getChatrooms() {
        WrestApi.chatrooms().then((chatrooms) => {
            this.chatrooms = chatrooms;
        });
    }

    public getChatroom(room: WcfrestContactPayload) {
        this.selected = room;
        WrestApi.chatroomMembers({ roomid: room.wxid }).then((members) => {
            this.members = members;
        });
    }

}
