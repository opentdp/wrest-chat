import { Component } from '@angular/core';

import { RobotApi, TablesLLModel, LlmodelFetchAllParam } from '../../openapi/wrobot';


@Component({
    selector: 'page-llmodel-list',
    templateUrl: 'list.html',
    styleUrls: ['list.scss']
})
export class LLModelListComponent {

    public llmodels: Array<TablesLLModel> = [];

    constructor() {
        this.getLLModels();
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
