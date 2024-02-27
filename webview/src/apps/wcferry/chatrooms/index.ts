import { Component } from '@angular/core';

import { WrestApi, WcfrestContactPayload } from '../../../openapi/wcfrest';


@Component({
    selector: 'page-wcf-chatrooms',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class WcfChatroomsComponent {

    public chatrooms: Array<WcfrestContactPayload> = [];
    public avatars: Record<string, string> = {};

    public selected!: WcfrestContactPayload;
    public members: Array<WcfrestContactPayload> = [];

    constructor() {
        this.getChatrooms();
    }

    public getChatrooms() {
        WrestApi.chatrooms().then((chatrooms) => {
            this.chatrooms = chatrooms;
            this.getAvatars();
        });
    }

    public getAvatars() {
        const ids = this.chatrooms.map((item) => item.wxid);
        WrestApi.avatars({ wxids: [...new Set(ids)] }).then((data) => {
            data && data.forEach((item) => {
                this.avatars[item.usr_name] = item.small_head_img_url;
            });
        });
    }

    public getChatroom(room: WcfrestContactPayload) {
        this.selected = room;
        WrestApi.chatroomMembers({ roomid: room.wxid }).then((members) => {
            this.members = members;
        });
    }

}
