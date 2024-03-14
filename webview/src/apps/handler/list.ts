import { Component } from '@angular/core';

import { UserLevels, SpecialRooms } from '../../openapi/const';
import { RobotApi, RobotHandler } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-handler-list',
    templateUrl: 'list.html'
})
export class HandlerListComponent {

    public userLevels = UserLevels;
    public specialRooms = SpecialRooms;

    public robotHandler: Array<RobotHandler> = [];

    public wcfChatrooms: Record<string, WcfrestContactPayload> = {};

    constructor() {
        this.getWcfChatrooms();
        this.getRobotHandlers();
    }

    public getRobotHandlers(reset?: boolean) {
        RobotApi.robotHandlers({ reset }).then((data) => {
            this.robotHandler = data || [];
        });
    }

    public getWcfChatrooms() {
        WrestApi.chatrooms().then((data) => {
            data.forEach((item) => this.wcfChatrooms[item.wxid] = item);
        });
    }

}
