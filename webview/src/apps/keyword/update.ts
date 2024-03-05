import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

import { KeywordLevels } from 'src/openapi/const';
import { RobotApi, KeywordUpdateParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-keyword-create',
    templateUrl: 'update.html'
})
export class KeywordUpdateComponent implements OnInit {

    public keywordLevels = KeywordLevels;

    public wcfChatrooms: Array<WcfrestContactPayload> = [];

    public formdata: KeywordUpdateParam = {} as KeywordUpdateParam;

    constructor(
        private router: Router,
        private route: ActivatedRoute
    ) {
        this.getWcfChatrooms();
    }

    public ngOnInit() {
        const rd = this.route.snapshot.paramMap.get('rd');
        rd && this.getKeyWord(+rd);
    }

    public getKeyWord(rd: number) {
        RobotApi.keywordDetail({ rd }).then((data) => {
            this.formdata = data;
        });
    }

    public updateKeyWord() {
        if (this.formdata.level) {
            this.formdata.level = +this.formdata.level;
        }
        RobotApi.keywordUpdate(this.formdata).then(() => {
            this.router.navigate(['keyword/list']);
        });
    }

    public getWcfChatrooms() {
        WrestApi.chatrooms().then((data) => {
            this.wcfChatrooms = data || [];
        });
    }

}
