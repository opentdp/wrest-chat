import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { RobotApi, LlmodelCreateParam } from '../../openapi/wrobot';


@Component({
    selector: 'page-llmodel-create',
    templateUrl: 'create.html',
    styleUrls: ['create.scss']
})
export class LLModelCreateComponent {

    public formdata: LlmodelCreateParam = {
        mid: '',
        family: 'Gemini',
        provider: 'google',
        model: 'gemini-pro',
        secret: '',
    };

    constructor(private router: Router) { }

    public createLLModel() {
        RobotApi.llmodelCreate(this.formdata).then(() => {
            this.router.navigate(['llmodel/list']);
        });
    }

}
