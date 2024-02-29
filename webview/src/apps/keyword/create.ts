import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { KeywordLevels } from 'src/openapi/const';
import { RobotApi, KeywordCreateParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-keyword-create',
    templateUrl: 'create.html',
    styleUrls: ['create.scss']
})
export class KeywordCreateComponent {

    public keywordLevels = KeywordLevels;

    public wcfChatrooms: Array<WcfrestContactPayload> = [];

    public formdata: KeywordCreateParam = { roomid: '-', level: 1 };

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
