import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { RoomLevels } from 'src/openapi/const';
import { RobotApi, ChatroomCreateParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-chatroom-create',
    templateUrl: 'create.html'
})
export class ChatroomCreateComponent {

    public roomLevels = RoomLevels;

    public wcfChatrooms: Array<WcfrestContactPayload> = [];

    public formdata: ChatroomCreateParam = {
        roomid: '',
    };

    constructor(private router: Router) {
        this.getWcfChatrooms();
    }

    public createChatroom() {
        if (this.formdata.level) {
            this.formdata.level = +this.formdata.level;
        }
        RobotApi.chatroomCreate(this.formdata).then(() => {
            this.router.navigate(['chatroom/list']);
        });
    }

    public getWcfChatrooms() {
        WrestApi.chatrooms().then((data) => {
            this.wcfChatrooms = data || [];
        });
    }

}
