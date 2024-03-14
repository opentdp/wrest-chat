import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { UserLevels, KeywordGroups, BadwordLevels } from '../../openapi/const';
import { RobotApi, KeywordCreateParam, RobotHandler } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-keyword-create',
    templateUrl: 'create.html'
})
export class KeywordCreateComponent {

    public userLevels = UserLevels;
    public keywordGroups = KeywordGroups;
    public badwordLevels = BadwordLevels;

    public robotHandler: Array<RobotHandler> = [];
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
        RobotApi.robotHandlers().then((data) => {
            this.robotHandler = data || [];
        });
    }

    public getWcfChatrooms() {
        WrestApi.chatrooms().then((data) => {
            this.wcfChatrooms = data || [];
        });
    }

}
