import { Component } from '@angular/core';

import { RobotApi, LlmodelFetchAllParam, TablesLLModel } from '../../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../../openapi/wcfrest';


@Component({
    selector: 'page-bot-llmodel',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class BotLLModelsComponent {

    public contacts: Record<string, WcfrestContactPayload> = {};

    public llmodels: Array<TablesLLModel> = [];

    constructor() {
        this.getContacts();
        this.getLLModels();
    }

    public getContacts() {
        WrestApi.contacts().then((data) => {
            data.forEach((item) => this.contacts[item.wxid] = item);
        });
    }

    public getLLModels() {
        const rq = {} as LlmodelFetchAllParam;
        RobotApi.llmodelList(rq).then((data) => {
            this.llmodels = data || [];
        });
    }

}
