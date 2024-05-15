import { Component } from '@angular/core';

import { UserLevels } from '../../openapi/const';
import { RobotApi, TablesLLModel } from '../../openapi/wrobot';


@Component({
    selector: 'page-llmodel-list',
    templateUrl: 'list.html'
})
export class LLModelListComponent {

    public userLevels = UserLevels;

    public llmodels: Array<TablesLLModel> = [];

    constructor() {
        this.getLLModels();
    }

    public getLLModels() {
        return RobotApi.llmodelList({}).then((data) => {
            this.llmodels = data || [];
        });
    }

    public deleteLLModel(item: TablesLLModel) {
        return RobotApi.llmodelDelete({ rd: item.rd }).then(() => {
            this.getLLModels();
        });
    }

}
