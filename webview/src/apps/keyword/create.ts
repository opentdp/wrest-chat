import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { UserLevels, SpecialRooms, KeywordGroups, BadwordLevels } from '../../openapi/const';
import { SundryApi, Handler } from '../../openapi/sundry';
import { RobotApi, KeywordCreateParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-keyword-create',
    templateUrl: 'create.html'
})
export class KeywordCreateComponent {

    public userLevels = UserLevels;
    public specialRooms = SpecialRooms;
    public keywordGroups = KeywordGroups;
    public badwordLevels = BadwordLevels;

    public robotHandler: Array<Handler> = [];
    public wcfChatrooms: Array<WcfrestContactPayload> = [];

    public formdata: KeywordCreateParam = {
        group: 'badword',
        roomid: '-',
        phrase: '',
        target: '',
        level: -1,
    };

    constructor(private router: Router) {
        this.getRobotHandlers();
        this.getWcfChatrooms();
    }

    public createKeyword() {
        if (this.formdata.level) {
            this.formdata.level = +this.formdata.level;
        }
        RobotApi.keywordCreate(this.formdata).then(() => {
            this.router.navigate(['keyword/list']);
        });
    }

    public getRobotHandlers() {
        SundryApi.handlerList({}).then((data) => {
            this.robotHandler = data || [];
        });
    }

    public getWcfChatrooms() {
        WrestApi.chatrooms().then((data) => {
            this.wcfChatrooms = data || [];
        });
    }

}
