import { Component } from '@angular/core';

import { WrestApi, WcfrestCommonPayload, WcfrestUserInfoPayload, WcfrestAvatarPayload } from '../../openapi/wcfrest';


@Component({
    selector: 'page-welcome',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class WelcomeComponent {

    public islogin = false;
    public loginqr!: WcfrestCommonPayload;

    public user!: WcfrestUserInfoPayload;
    public avatar!: WcfrestAvatarPayload;

    constructor() {
        this.checkLogin();
    }

    public async checkLogin() {
        this.islogin = await WrestApi.isLogin();
        if (this.islogin) {
            return this.getSelfInfo();
        }
        setTimeout(() => this.checkLogin(), 200 * 1000);
        this.loginqr = await WrestApi.loginQr();
    }

    public async getSelfInfo() {
        this.user = await WrestApi.selfInfo();
        if (this.user && this.user.wxid) {
            const avatars = await WrestApi.avatars({ wxids: [this.user.wxid] });
            if (avatars && avatars.length > 0) {
                this.avatar = avatars[0];
            }
        }
    }

}