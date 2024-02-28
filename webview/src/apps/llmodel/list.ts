import { Component } from '@angular/core';

import { RobotApi, TablesLLModel, LlmodelFetchAllParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-llmodel-list',
    templateUrl: 'list.html',
    styleUrls: ['list.scss']
})
export class LLModelListComponent {

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

    public deleteLLModel(item: TablesLLModel) {
        RobotApi.llmodelDelete({ mid: item.mid }).then(() => {
            this.getLLModels();
        });
    }

}
