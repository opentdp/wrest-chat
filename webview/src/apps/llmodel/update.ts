import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

import { RobotApi, LlmodelUpdateParam } from '../../openapi/wrobot';


@Component({
    selector: 'page-llmodel-create',
    templateUrl: 'update.html'
})
export class LLModelUpdateComponent implements OnInit {

    public formdata: LlmodelUpdateParam = {} as LlmodelUpdateParam;

    constructor(
        private router: Router,
        private route: ActivatedRoute
    ) { }

    public ngOnInit() {
        const mid = this.route.snapshot.paramMap.get('mid');
        mid && this.getLLModel(mid);
    }

    public getLLModel(mid: string) {
        RobotApi.llmodelDetail({ mid }).then((data) => {
            this.formdata = data;
        });
    }

    public updateLLModel() {
        RobotApi.llmodelUpdate(this.formdata).then(() => {
            this.router.navigate(['llmodel/list']);
        });
    }

}
