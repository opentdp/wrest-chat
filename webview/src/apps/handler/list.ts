import { Component } from '@angular/core';

import { RobotApi, RobotHandler } from '../../openapi/wrobot';


@Component({
    selector: 'page-handler-list',
    templateUrl: 'list.html'
})
export class HandlerListComponent {

    public robotHandler: Array<RobotHandler> = [];

    constructor() {
        this.getRobotHandlers();
    }

    public getRobotHandlers() {
        RobotApi.robotHandlers().then((data) => {
            this.robotHandler = data || [];
        });
    }

}
