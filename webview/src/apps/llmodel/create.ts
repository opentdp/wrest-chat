import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { AiModels, UserLevels } from '../../openapi/const';
import { RobotApi, LlmodelCreateParam } from '../../openapi/wrobot';


@Component({
    selector: 'page-llmodel-create',
    templateUrl: 'create.html'
})
export class LLModelCreateComponent {

    public aiModels = AiModels;
    public userLevels = UserLevels;

    public formdata: LlmodelCreateParam = {
        mid: '',
        level: -1,
        family: '',
        provider: 'google',
        model: 'gemini-pro',
        secret: '',
    };

    constructor(private router: Router) { }

    public createLLModel() {
        if (this.formdata.level) {
            this.formdata.level = +this.formdata.level;
        }
        return RobotApi.llmodelCreate(this.formdata).then(() => {
            this.router.navigate(['llmodel/list']);
        });
    }

}
