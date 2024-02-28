import { Component, OnDestroy } from '@angular/core';


@Component({
    selector: 'page-wcferry-receiver',
    templateUrl: 'receiver.html',
    styleUrls: ['receiver.scss']
})
export class WcferryReceiverComponent implements OnDestroy {

    public ws!: WebSocket;
    public messages: Array<string> = [];

    constructor() {
        this.startSocket();
    }

    ngOnDestroy(): void {
        this.ws.close();
        this.messages = [];
    }

    public async startSocket() {
        const url = location.origin.replace(/^http/, 'ws');
        this.ws = new WebSocket(url + '/api/socket_receiver');
        this.ws.onopen = () => {
            this.messages.push('WebSocket is connected.');
        };
        this.ws.onmessage = event => {
            this.messages.push(event.data);
        };
        this.ws.onerror = (error) => {
            this.messages.push('WebSocket Error:' + error);
        };
        this.ws.onclose = () => {
            this.messages.push('WebSocket is closed now.');
        };
    }

}
