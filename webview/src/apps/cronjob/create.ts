import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { CronjobTypes } from '../../openapi/const';
import { CronApi, CronjobCreateParam } from '../../openapi/cronjob';


@Component({
    selector: 'page-cronjob-create',
    templateUrl: 'create.html'
})
export class CronjobCreateComponent {

    public cronjobTypes = CronjobTypes;

    public formdata: CronjobCreateParam = {
        name: '',
        second: '0',
        minute: '*',
        hour: '*',
        day_of_month: '*',
        month: '*',
        day_of_week: '*',
        type: 'TEXT',
        timeout: 30,
        directory: '.',
        content: '',
        deliver: '-',
    };

    constructor(private router: Router) {
    }

    public createCronjob() {
        const data = this.formdata;
        const time = data.second + data.minute + data.hour + data.day_of_month + data.month + data.day_of_week;
        if (time === '******') {
            window.postMessage({ message: '排程不可全为 *', type: 'danger' });
            return;
        }
        CronApi.cronjobCreate(this.formdata).then(() => {
            this.router.navigate(['cronjob/list']);
        });
    }

}
