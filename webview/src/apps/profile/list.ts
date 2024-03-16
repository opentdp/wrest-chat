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
        roomid: '-',
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
        RobotApi.profileList(this.formdata).then((data) => {
            this.profiles = data || [];
        });
        if (this.formdata.roomid) {
            this.getWcfRoomMembers(this.formdata.roomid);
        } else {
            Object.keys(this.wcfChatrooms).forEach((id) => {
                this.getWcfRoomMembers(id);
            });
        }
    }

    public deleteProfile(item: TablesProfile) {
        RobotApi.profileDelete({ rd: item.rd }).then(() => {
            this.getProfiles();
        });
    }

    public getWcfContacts() {
        WrestApi.contacts().then((data) => {
            data.forEach((item) => this.wcfContacts[item.wxid] = item);
        });
    }

    public getWcfChatrooms() {
        WrestApi.chatrooms().then((data) => {
            data.forEach((item) => this.wcfChatrooms[item.wxid] = item);
        });
    }

    public getWcfRoomMembers(id: string) {
        if (this.wcfRoomMembers[id]) {
            return; //已获取
        }
        WrestApi.chatroomMembers({ roomid: id }).then((data) => {
            this.wcfRoomMembers[id] = {};
            data && data.forEach((item) => {
                this.wcfRoomMembers[id][item.wxid] = item;
            });
        });
    }

    public getWcfAvatars(ids: string[]) {
        WrestApi.avatars({ wxids: [...new Set(ids)] }).then((data) => {
            data && data.forEach((item) => {
                this.wcfAvatars[item.usr_name] = item.small_head_img_url;
            });
        });
    }

    public getLocalTime(ts: number) {
        return new Date(ts * 1000).toLocaleString();
    }

}
