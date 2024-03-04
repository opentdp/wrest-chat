import { Component } from '@angular/core';
import {ActivatedRoute, Route, Router} from '@angular/router';

import {RobotApi, LlmodelCreateParam, LlmodelUpdateParam} from '../../openapi/wrobot';


@Component({
    selector: 'page-llmodel-create',
    templateUrl: 'update.html'
})
export class LLModelUpdateComponent {

    public formdata: LlmodelUpdateParam = {} as LlmodelUpdateParam

    constructor(
        private router: Router,
        private route: ActivatedRoute
    ) { }

    public ngOnInit() {
        const mid = this.route.snapshot.paramMap.get('mid');
        mid && this.getLLModel(mid);
    }

    public getLLModel(mid: string) {
        RobotApi.llmodelDetail({mid}).then((data) => {
            this.formdata = data
        });
    }

    public updateLLModel() {
        RobotApi.llmodelUpdate(this.formdata).then(() => {
            this.router.navigate(['llmodel/list']);
        });
    }

}
