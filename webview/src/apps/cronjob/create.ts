import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { CronjobTypes } from '../../openapi/const';
import { RobotApi, CronjobCreateParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-cronjob-create',
    templateUrl: 'create.html'
})
export class CronjobCreateComponent {

    public cronjobTypes = CronjobTypes;

    public wcfFriends: Array<WcfrestContactPayload> = [];
    public wcfChatrooms: Array<WcfrestContactPayload> = [];
    public wcfRoomMembers: Record<string, Array<WcfrestContactPayload>> = {};

    public conacts: Array<WcfrestContactPayload> = [];
    public conactsFilter = '';

    public deliver = ['-', '-', '-'];

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
        this.getWcfFriends();
        this.getWcfChatrooms();
    }

    public createCronjob() {
        const data = this.formdata;
        const time = data.second + data.minute + data.hour + data.day_of_month + data.month + data.day_of_week;
        if (time === '******') {
            window.postMessage({ message: '排程不可全为 *', type: 'danger' });
            return Promise.resolve();
        }
        this.formdata.deliver = Object.values(this.deliver).join(',');
        return RobotApi.cronjobCreate(this.formdata).then(() => {
            this.router.navigate(['cronjob/list']);
        });
    }

    public changeDeliver() {
        this.deliver[1] = '-';
        this.deliver[2] = '-';
        this.changeConacts();
    }

    public changeConacts() {
        const id = this.deliver[1] || '-';
        return this.getWcfRoomMembers(this.deliver[1]).then(() => {
            this.conacts = id == '-' ? this.wcfFriends : this.wcfRoomMembers[id] || [];
        });
    }

    public getWcfFriends() {
        return WrestApi.friends().then((data) => {
            this.wcfFriends = data || [];
        });
    }

    public getWcfChatrooms() {
        return WrestApi.chatrooms().then((data) => {
            this.wcfChatrooms = data || [];
        });
    }

    public getWcfRoomMembers(id: string) {
        if (this.wcfRoomMembers[id]) {
            return Promise.resolve(); //已获取
        }
        this.wcfRoomMembers[id] = []; //初始化
        return WrestApi.chatroomMembers({ roomid: id }).then((data) => {
            this.wcfRoomMembers[id] = data || [];
        });
    }

}
