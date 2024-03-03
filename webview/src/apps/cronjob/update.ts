import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import { CronApi, CronjobUpdateParam } from '../../openapi/cronjob';


@Component({
    selector: 'page-cronjob-update',
    templateUrl: 'update.html'
})
export class CronjobUpdateComponent implements OnInit {

    public formdata = {} as CronjobUpdateParam;

    constructor(
        private router: Router,
        private route: ActivatedRoute
    ) {
    }

    public ngOnInit() {
        const rd = this.route.snapshot.paramMap.get('rd');
        rd && this.getCronjob(+rd);
    }

    public getCronjob(rd: number) {
        CronApi.cronjobDetail({ rd }).then((data) => {
            this.formdata = data;
        });
    }

    public updateCronjob() {
        const data = this.formdata;
        const time = data.second + data.minute + data.hour + data.day_of_month + data.month + data.day_of_week;
        if (time === '******') {
            window.postMessage({ message: '排程不可全为 *', type: 'danger' });
            return;
        }
        CronApi.cronjobUpdate(this.formdata).then(() => {
            this.router.navigate(['cronjob/list']);
        });
    }

}
