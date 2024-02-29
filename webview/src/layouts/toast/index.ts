import { Component, HostBinding } from '@angular/core';

@Component({
    selector: 'layout-toast',
    templateUrl: 'index.html',
    styleUrls: ['index.scss']
})
export class LayoutToastComponent {

    @HostBinding()
    public class = 'toast-container position-fixed top-0 end-0 p-3';

    @HostBinding()
    public style = 'z-index: 1200';

    public items: Toast[] = [];

    constructor() {
        // 处理 js 异常
        window.onerror = (message) => {
            message = String(message);
            this.show({ message, classname: 'bg-danger text-light' });
        };
        // 处理 postMessage
        window.addEventListener('message', e => {
            if (typeof e.data === 'string') {
                this.show({ message: e.data, classname: 'bg-success text-light' });
            } else {
                e.data.message && this.show(e.data);
            }
        });
        // 处理 promise 未捕获的 rejection
        window.addEventListener("unhandledrejection", e => {
            this.show({ message: e.reason, classname: 'bg-danger text-light' });
            e.preventDefault && e.preventDefault();
        });
    }

    public show(toast: Toast) {
        this.items.push(toast);
    }

    public remove(toast: Toast) {
        this.items = this.items.filter((t) => t !== toast);
    }

    public clear() {
        this.items.splice(0, this.items.length);
    }

}

export interface Toast {
    classname?: string;
    message: string;
    delay?: number;
}
