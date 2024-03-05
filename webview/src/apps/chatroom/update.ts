import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

import { RoomLevels } from 'src/openapi/const';
import { RobotApi, ChatroomUpdateParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-chatroom-update',
    templateUrl: 'update.html'
})
export class ChatroomUpdateComponent implements OnInit {

    public roomLevels = RoomLevels;

    public wcfChatrooms: Array<WcfrestContactPayload> = [];

    public formdata = {} as ChatroomUpdateParam;

    constructor(
        private router: Router,
        private route: ActivatedRoute
    ) {
        this.getWcfChatrooms();
    }

    public ngOnInit() {
        const rd = this.route.snapshot.paramMap.get('rd');
        rd && this.getChatroom(+rd);
    }

    public getChatroom(rd: number) {
        RobotApi.chatroomDetail({ rd }).then((data) => {
            this.formdata = data;
        });
    }

    public updateChatroom() {
        if (this.formdata.level) {
            this.formdata.level = +this.formdata.level;
        }
        RobotApi.chatroomUpdate(this.formdata).then(() => {
            this.router.navigate(['chatroom/list']);
        });
    }

    public getWcfChatrooms() {
        WrestApi.chatrooms().then((data) => {
            this.wcfChatrooms = data || [];
        });
    }

}
