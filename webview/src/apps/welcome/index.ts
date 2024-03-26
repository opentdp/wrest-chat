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

    public formdata = {
        token: '',
    };

    constructor() {
        this.checkLogin();
        this.refreshQrcode();
        // 获取会话存储的令牌
        this.formdata.token = sessionStorage.getItem('token') || '';
    }

    public submitForm() {
        sessionStorage.setItem('token', this.formdata.token);
        location.reload();
    }

    public async checkLogin() {
        this.islogin = await WrestApi.isLogin();
        if (this.islogin) {
            return this.getSelfInfo();
        }
        setTimeout(() => this.checkLogin(), 5 * 1000);
    }

    public async refreshQrcode() {
        this.loginqr = await WrestApi.loginQr();
        if (!this.islogin) {
            const t = this.loginqr.result ? 200 : 5;
            setTimeout(() => this.refreshQrcode(), t * 1000);
        }
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