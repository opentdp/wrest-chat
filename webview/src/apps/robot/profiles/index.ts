import { Component } from '@angular/core';

import { LevelData } from 'src/openapi/const';
import { RobotApi, ProfileFetchAllParam, TablesProfile } from '../../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../../openapi/wcfrest';


@Component({
    selector: 'page-bot-profiles',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class BotProfilesComponent {

    public levels = LevelData;
    public timestamp = 0;

    public contacts: Record<string, WcfrestContactPayload> = {};
    public roomMembers: Record<string, Record<string, WcfrestContactPayload>> = {};

    public profiles: Array<TablesProfile> = [];

    constructor() {
        this.getContacts();
        this.getProfiles();
        this.timestamp = new Date().getTime();
    }

    public getContacts() {
        WrestApi.contacts().then((data) => {
            data.forEach((item) => this.contacts[item.wxid] = item);
        });
    }

    public getProfiles() {
        const rq = {} as ProfileFetchAllParam;
        RobotApi.profileList(rq).then((data) => {
            this.profiles = data || [];
            // 获取群成员列表
            const roomids = this.profiles.map((item) => item.roomid);
            this.getRoomMembers(roomids);
        });
    }

    public getRoomMembers(ids: string[]) {
        ids.forEach((id) => {
            if (this.roomMembers[id]) {
                return;
            }
            this.roomMembers[id] = {};
            WrestApi.chatroomMembers({ roomid: id }).then((items) => {
                items.forEach((item) => {
                    this.roomMembers[id][item.wxid] = item;
                });
            });
        });
    }

    public getLocalTime(ts: number) {
        return new Date(ts * 1000).toLocaleString();
    }

}
