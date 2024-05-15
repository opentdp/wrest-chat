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
        this.register();
    }

    public create(toast: Toast) {
        toast.classname = `bg-${toast.classname || 'success'} text-light`;
        this.items.push(toast);
    }

    public remove(toast: Toast) {
        this.items = this.items.filter((t) => t !== toast);
    }

    public clear() {
        this.items.splice(0, this.items.length);
    }

    private register() {
        // 处理 js 异常
        window.onerror = (message) => {
            this.create({ message: String(message), classname: 'danger' });
        };
        // 处理 promise 未捕获的 rejection
        window.addEventListener('unhandledrejection', e => {
            this.create({ message: String(e.reason), classname: 'danger' });
            e.preventDefault && e.preventDefault();
        });
        // 处理 postMessage 信息
        window.addEventListener('message', e => {
            if (e && e.data && e.data.type) {
                this.create({ message: String(e.data.message), classname: e.data.type });
            }
        });
    }

}

export interface Toast {
    classname: string;
    message: string;
    delay?: number;
}
