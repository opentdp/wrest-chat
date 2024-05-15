import { Component } from '@angular/core';

import { UserLevels } from '../../openapi/const';
import { RobotApi, TablesProfile, ProfileFetchAllParam } from '../../openapi/wrobot';
import { WrestApi, WcfrestContactPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-profile-list',
    templateUrl: 'list.html'
})
export class ProfileListComponent {

    public userLevels = UserLevels;

    public wcfAvatars: Record<string, string> = {};
    public wcfContacts: Record<string, WcfrestContactPayload> = {};
    public wcfChatrooms: Record<string, WcfrestContactPayload> = {};
    public wcfRoomMembers: Record<string, Record<string, WcfrestContactPayload>> = {};

    public profiles: Array<TablesProfile> = [];

    public formdata: ProfileFetchAllParam = {
        roomid: '',
        level: 0,
    };

    public timestamp = 0;

    constructor() {
        this.getProfiles();
        this.getWcfContacts();
        this.getWcfChatrooms();
        this.timestamp = new Date().getTime();
    }

    public getProfiles() {
        if (this.formdata.level) {
            this.formdata.level = +this.formdata.level;
        }
        return RobotApi.profileList(this.formdata).then((data) => {
            this.profiles = data || [];
            // 获取群成员
            this.profiles.forEach((item) => {
                if (item.roomid && item.roomid.indexOf('@chatroom') > 0) {
                    this.getWcfRoomMembers(item.roomid);
                }
            });
        });
    }

    public deleteProfile(item: TablesProfile) {
        return RobotApi.profileDelete({ rd: item.rd }).then(() => {
            this.getProfiles();
        });
    }

    public getWcfContacts() {
        return WrestApi.contacts().then((data) => {
            data.forEach((item) => this.wcfContacts[item.wxid] = item);
        });
    }

    public getWcfChatrooms() {
        return WrestApi.chatrooms().then((data) => {
            data.forEach((item) => this.wcfChatrooms[item.wxid] = item);
        });
    }

    public getWcfRoomMembers(id: string) {
        if (this.wcfRoomMembers[id]) {
            return Promise.resolve(); //已获取
        }
        this.wcfRoomMembers[id] = {}; //初始化
        return WrestApi.chatroomMembers({ roomid: id }).then((data) => {
            data && data.forEach((item) => {
                this.wcfRoomMembers[id][item.wxid] = item;
            });
        });
    }

    public getLocalTime(ts: number) {
        return new Date(ts * 1000).toLocaleString();
    }

}
