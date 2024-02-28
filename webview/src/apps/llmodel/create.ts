import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { RobotApi, LlmodelFetchAllParam, TablesLLModel } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-llmodel-create',
    templateUrl: 'create.html',
    styleUrls: ['create.scss']
})
export class LLModelCreateComponent {

    public contacts: Record<string, WcfrestContactPayload> = {};

    public llmodels: Array<TablesLLModel> = [];

    public formdata = {} as TablesLLModel;

    constructor(private router: Router) {
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

    public createLLModel() {
        RobotApi.llmodelCreate(this.formdata).then(() => {
            this.router.navigate(['llmodel/list']);
        });
    }

}
