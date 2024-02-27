import { NgModule, LOCALE_ID } from '@angular/core';

import { registerLocaleData } from '@angular/common';
import zh from '@angular/common/locales/zh';
registerLocaleData(zh, 'zh');

import { FormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { NgbModule } from '@ng-bootstrap/ng-bootstrap';

import { QrCodeModule } from 'ng-qrcode';


@NgModule({
    exports: [
        FormsModule,
        BrowserAnimationsModule,
        NgbModule,
        QrCodeModule,
    ],
    providers: [
        { provide: LOCALE_ID, useValue: 'zh' },
    ],
})
export class BasicModule { }
