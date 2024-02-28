import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { RobotApi, KeywordCreateParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-keyword-create',
    templateUrl: 'create.html',
    styleUrls: ['create.scss']
})
export class KeywordCreateComponent {

    public chatrooms: Array<WcfrestContactPayload> = [];

    public formdata = { level: 1, roomid: '-' } as KeywordCreateParam;

    constructor(private router: Router) {
        this.getChatrooms();
    }

    public createKeyword() {
        RobotApi.keywordCreate(this.formdata).then(() => {
            this.router.navigate(['keyword/list']);
        });
    }

    public getChatrooms() {
        WrestApi.chatrooms().then((data) => {
            this.chatrooms = data || [];
        });
    }

}
