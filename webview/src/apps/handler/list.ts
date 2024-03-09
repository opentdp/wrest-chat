import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { RobotApi, RobotHandler } from '../../openapi/wrobot';


@Component({
    selector: 'page-handler-list',
    templateUrl: 'list.html'
})
export class HandlerListComponent {

    public robotHandler: Array<RobotHandler> = [];


    constructor(private router: Router) {
        this.getRobotHandlers();
    }

    public getRobotHandlers() {
        RobotApi.robotHandlers().then((data) => {
            this.robotHandler = data || [];
        });
    }

}
