import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

import { UserLevels } from 'src/openapi/const';
import { RobotApi, LlmodelUpdateParam } from '../../openapi/wrobot';


@Component({
    selector: 'page-llmodel-create',
    templateUrl: 'update.html'
})
export class LLModelUpdateComponent implements OnInit {

    public userLevels = UserLevels;

    public formdata: LlmodelUpdateParam = {} as LlmodelUpdateParam;

    constructor(
        private router: Router,
        private route: ActivatedRoute
    ) { }

    public ngOnInit() {
        const rd = this.route.snapshot.paramMap.get('rd');
        rd && this.getLLModel(+rd);
    }

    public getLLModel(rd: number) {
        RobotApi.llmodelDetail({ rd }).then((data) => {
            this.formdata = data;
        });
    }

    public updateLLModel() {
        RobotApi.llmodelUpdate(this.formdata).then(() => {
            this.router.navigate(['llmodel/list']);
        });
    }

}
