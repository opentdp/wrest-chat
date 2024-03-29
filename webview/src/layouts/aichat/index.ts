import { Component, OnDestroy, ViewChild, ElementRef } from '@angular/core';

import { SundryApi, AiChatMsgHistory } from '../../openapi/sundry';


@Component({
    selector: 'layout-aichat',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class LayoutAichatComponent implements OnDestroy {

    @ViewChild('scrollLayout')
    private scrollLayout!: ElementRef;

    public messages: Array<AiChatMsgHistory> = [];

    public content = '';

    public constructor() {
    }

    public ngOnDestroy() {
        this.messages = [];
    }

    public sendMessage() {
        this.scrollToBottom();
        this.messages.push({ role: 'user', content: this.content });
        SundryApi.aichatText({ wxid: 'webui', message: this.content }).then((data) => {
            this.messages.push({ role: 'assistant', content: data });
            this.scrollToBottom();
        });
    }

    public scrollToBottom() {
        setTimeout(() => {
            const el = this.scrollLayout.nativeElement;
            el.scrollTop = el.scrollHeight;
        }, 100);
    }

}
