import { Component } from '@angular/core';
import { Router } from '@angular/router';

import { CronjobTypes } from '../../openapi/const';
import { SundryApi, CronjobCreateParam } from '../../openapi/sundry';
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

    public deliver = ['wechat', '-', '-'];

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
            return;
        }
        this.formdata.deliver = Object.values(this.deliver).join(',');
        SundryApi.cronjobCreate(this.formdata).then(() => {
            this.router.navigate(['cronjob/list']);
        });
    }

    public async changeConacts() {
        const id = this.deliver[1] || '-';
        await this.getWcfRoomMembers(this.deliver[1]);
        this.conacts = id == '-' ? this.wcfFriends : this.wcfRoomMembers[id] || [];
    }

    public getWcfFriends() {
        WrestApi.friends().then((data) => {
            this.wcfFriends = data || [];
        });
    }

    public getWcfChatrooms() {
        WrestApi.chatrooms().then((data) => {
            this.wcfChatrooms = data || [];
        });
    }

    public getWcfRoomMembers(id: string) {
        if (this.wcfRoomMembers[id]) {
            return; //已获取
        }
        return WrestApi.chatroomMembers({ roomid: id }).then((data) => {
            this.wcfRoomMembers[id] = data || [];
        });
    }

}
