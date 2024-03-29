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
        return RobotApi.keywordCreate(this.formdata).then(() => {
            this.router.navigate(['keyword/list']);
        });
    }

    public changeHandler() {
        const h = this.robotHandler.find((h) => h.command === this.formdata.target);
        this.formdata.level = h ? h.level : -1;
    }

    public getRobotHandlers() {
        return SundryApi.systemHandlers({}).then((data) => {
            this.robotHandler = data || [];
        });
    }

    public getWcfChatrooms() {
        return WrestApi.chatrooms().then((data) => {
            this.wcfChatrooms = data || [];
        });
    }

}
