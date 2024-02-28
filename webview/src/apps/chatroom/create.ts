import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { RobotApi, ChatroomCreateParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-chatroom-create',
    templateUrl: 'create.html',
    styleUrls: ['create.scss']
})
export class ChatroomCreateComponent {

    public chatrooms: Array<WcfrestContactPayload> = [];

    public formdata = { level: 1 } as ChatroomCreateParam;

    constructor(private router: Router) {
        this.getChatrooms();
    }

    public createChatroom() {
        RobotApi.chatroomCreate(this.formdata).then(() => {
            this.router.navigate(['chatroom/list']);
        });
    }

    public getChatrooms() {
        WrestApi.chatrooms().then((data) => {
            this.chatrooms = data || [];
        });
    }

}
