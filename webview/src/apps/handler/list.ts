import { Component } from '@angular/core';

import { UserLevels } from '../../openapi/const';
import { RobotApi, RobotHandler } from '../../openapi/wrobot';


@Component({
    selector: 'page-handler-list',
    templateUrl: 'list.html'
})
export class HandlerListComponent {

    public userLevels = UserLevels;

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
