import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { CronApi, CronjobCreateParam } from '../../openapi/cronjob';


@Component({
    selector: 'page-cronjob-create',
    templateUrl: 'create.html',
    styleUrls: ['create.scss']
})
export class CronjobCreateComponent {

    public formdata: CronjobCreateParam = {
        name: '',
        second: '*',
        minute: '*',
        hour: '*',
        day_of_month: '*',
        month: '*',
        day_of_week: '*',
        type: 'BAT',
        timeout: 30,
        directory: 'C:\\',
        content: '',
    };

    constructor(private router: Router) {
    }

    public createCronjob() {
        CronApi.cronjobCreate(this.formdata).then(() => {
            this.router.navigate(['cronjob/list']);
        });
    }

}
