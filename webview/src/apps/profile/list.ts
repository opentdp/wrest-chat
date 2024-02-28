import { Component } from '@angular/core';

import { LevelData } from 'src/openapi/const';
import { RobotApi, TablesProfile, ProfileFetchAllParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-profile-list',
    templateUrl: 'list.html',
    styleUrls: ['list.scss']
})
export class ProfileListComponent {

    public levels = LevelData;
    public timestamp = 0;

    public avatars: Record<string, string> = {};
    public contacts: Record<string, WcfrestContactPayload> = {};
    public roomMembers: Record<string, Record<string, WcfrestContactPayload>> = {};

    public profiles: Array<TablesProfile> = [];

    constructor() {
        this.getContacts();
        this.getProfiles();
        this.timestamp = new Date().getTime();
    }

    public getProfiles() {
        const rq = {} as ProfileFetchAllParam;
        RobotApi.profileList(rq).then((data) => {
            this.profiles = data || [];
            // 获取群成员列表
            const ids = this.profiles.map((item) => item.roomid);
            this.getRoomMembers(ids);
        });
    }

    public deleteProfile(item: TablesProfile) {
        RobotApi.chatroomDelete({ roomid: item.roomid }).then(() => {
            this.getProfiles();
        });
    }

    public getContacts() {
        WrestApi.contacts().then((data) => {
            data.forEach((item) => this.contacts[item.wxid] = item);
        });
    }

    public getRoomMembers(ids: string[]) {
        [...new Set(ids)].forEach((id) => {
            if (id === '-' || this.roomMembers[id]) {
                return;
            }
            this.roomMembers[id] = {};
            WrestApi.chatroomMembers({ roomid: id }).then((data) => {
                data && data.forEach((item) => {
                    this.roomMembers[id][item.wxid] = item;
                });
            });
        });
    }

    public getAvatars(ids: string[]) {
        WrestApi.avatars({ wxids: [...new Set(ids)] }).then((data) => {
            data && data.forEach((item) => {
                this.avatars[item.usr_name] = item.small_head_img_url;
            });
        });
    }

    public getLocalTime(ts: number) {
        return new Date(ts * 1000).toLocaleString();
    }

}
