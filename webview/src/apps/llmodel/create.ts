import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { RobotApi, TablesLLModel } from '../../openapi/wrobot';


@Component({
    selector: 'page-llmodel-create',
    templateUrl: 'create.html',
    styleUrls: ['create.scss']
})
export class LLModelCreateComponent {

    public llmodels: Array<TablesLLModel> = [];

    public formdata = {} as TablesLLModel;

    constructor(private router: Router) { }

    public createLLModel() {
        RobotApi.llmodelCreate(this.formdata).then(() => {
            this.router.navigate(['llmodel/list']);
        });
    }

}
