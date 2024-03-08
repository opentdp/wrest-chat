import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { CommandList, KeywordGroups, KeywordLevels } from '../../openapi/const';
import { RobotApi, KeywordCreateParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-keyword-create',
    templateUrl: 'create.html'
})
export class KeywordCreateComponent {

    public commandList = CommandList;
    public keywordGroups = KeywordGroups;
    public keywordLevels = KeywordLevels;

    public wcfChatrooms: Array<WcfrestContactPayload> = [];

    public formdata: KeywordCreateParam = {
        group: 'badword',
        roomid: '-',
        phrase: '',
        target: '',
        level: 1,
    };

    constructor(private router: Router) {
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

    public getWcfChatrooms() {
        WrestApi.chatrooms().then((data) => {
            this.wcfChatrooms = data || [];
        });
    }

}
