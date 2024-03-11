import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { RoomLevels } from '../../openapi/const';
import { RobotApi,TablesLLModel, ChatroomCreateParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-chatroom-create',
    templateUrl: 'create.html'
})
export class ChatroomCreateComponent {

    public roomLevels = RoomLevels;

    public wcfChatrooms: Array<WcfrestContactPayload> = [];

    public llmodels: Array<TablesLLModel> = [];
    
    public formdata: ChatroomCreateParam = {
        level: 1,
        roomid: '',
        pat_return: 'false',
    };

    constructor(private router: Router) {
        this.getWcfChatrooms();
        this.getLLModels();
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

    public getLLModels() {
        RobotApi.llmodelList({}).then((data) => {
            this.llmodels = data || [];
        });
    }

}
