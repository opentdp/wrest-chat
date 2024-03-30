import { Component, OnDestroy, ViewChild, ElementRef, Input } from '@angular/core';

import { SundryApi, AiChatMsgHistory, AiChatUserConfig } from '../../openapi/sundry';


@Component({
    selector: 'layout-aichat',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class LayoutAichatComponent implements OnDestroy {

    @ViewChild('scrollLayout')
    private scrollLayout!: ElementRef;

    @Input()
    public avatar = '/assets/icon.png';

    public config = {} as AiChatUserConfig;
    public messages: Array<AiChatMsgHistory> = [];

    public loading = false;
    public content = '';

    public constructor() {
        this.getConfig();
    }

    public ngOnDestroy() {
        this.messages = [];
    }

    public getConfig() {
        return SundryApi.aichatConfig({ wxid: 'webui', message: '' }).then((data) => {
            this.config = data || {};
            this.messages = this.config.msg_historys || [];
        });
    }

    public sendMessage() {
        this.loading = true;
        this.scrollToBottom();
        const aiwait = { role: 'assistant', content: '正在思考...' };
        this.messages.push({ role: 'user', content: this.content }, aiwait);
        // 请求结果
        return SundryApi.aichatText({ wxid: 'webui', message: this.content }).then((data) => {
            aiwait.content = data || '未知错误';
            this.content = '';
        }).catch((err) => {
            aiwait.content = err || '未知错误';
        }).finally(() => {
            this.scrollToBottom();
            this.loading = false;
        });
    }

    public scrollToBottom() {
        setTimeout(() => {
            const el = this.scrollLayout.nativeElement;
            el.scrollTop = el.scrollHeight;
        }, 100);
    }

}
