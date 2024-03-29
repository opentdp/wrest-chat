import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

import { UserLevels, SpecialRooms, KeywordGroups, BadwordLevels } from '../../openapi/const';
import { RobotApi, KeywordUpdateParam } from '../../openapi/wrobot';
import { SundryApi, Handler } from '../../openapi/sundry';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-keyword-create',
    templateUrl: 'update.html'
})
export class KeywordUpdateComponent implements OnInit {

    public userLevels = UserLevels;
    public specialRooms = SpecialRooms;
    public keywordGroups = KeywordGroups;
    public badwordLevels = BadwordLevels;

    public robotHandler: Array<Handler> = [];
    public wcfChatrooms: Array<WcfrestContactPayload> = [];

    public formdata: KeywordUpdateParam = {} as KeywordUpdateParam;

    constructor(
        private router: Router,
        private route: ActivatedRoute
    ) {
        this.getRobotHandlers();
        this.getWcfChatrooms();
    }

    public ngOnInit() {
        const rd = this.route.snapshot.paramMap.get('rd');
        rd && this.getKeyWord(+rd);
    }

    public getKeyWord(rd: number) {
        return RobotApi.keywordDetail({ rd }).then((data) => {
            this.formdata = data;
        });
    }

    public updateKeyWord() {
        if (this.formdata.level) {
            this.formdata.level = +this.formdata.level;
        }
        return RobotApi.keywordUpdate(this.formdata).then(() => {
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
