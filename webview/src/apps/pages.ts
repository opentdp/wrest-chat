import { LoginGuard } from '../helpers/login-guard';

import { WelcomeComponent } from './welcome';
import { Alert404Component } from './alert/404';

import { ChatroomListComponent } from './chatroom/list';
import { ChatroomCreateComponent } from './chatroom/create';

import { CronjobListComponent } from './cronjob/list';
import { CronjobCreateComponent } from './cronjob/create';
import { CronjobUpdateComponent } from './cronjob/update';

import { LLModelCreateComponent } from './llmodel/create';
import { LLModelListComponent } from './llmodel/list';

import { KeywordListComponent } from './keyword/list';
import { KeywordCreateComponent } from './keyword/create';

import { ProfileListComponent } from './profile/list';
import { ProfileCreateComponent } from './profile/create';
import { ProfileUpdateComponent } from './profile/update';

import { SettingListComponent } from './setting/list';
import { SettingUpdateComponent } from './setting/update';

import { WcferryChatroomComponent } from './wcferry/chatroom';
import { WcferryContactComponent } from './wcferry/contact';
import { WcferryReceiverComponent } from './wcferry/receiver';


export const AppComponents = [
    WelcomeComponent,
    Alert404Component,

    ChatroomListComponent,
    ChatroomCreateComponent,

    CronjobListComponent,
    CronjobCreateComponent,
    CronjobUpdateComponent,

    LLModelCreateComponent,
    LLModelListComponent,

    KeywordListComponent,
    KeywordCreateComponent,

    ProfileListComponent,
    ProfileCreateComponent,
    ProfileUpdateComponent,

    SettingListComponent,
    SettingUpdateComponent,

    WcferryChatroomComponent,
    WcferryContactComponent,
    WcferryReceiverComponent,
];

//////////////////////////////////////////////////////////////////

import { Routes } from '@angular/router';

export const AppRoutes: Routes = [
    { path: 'welcome', component: WelcomeComponent },

    { path: 'chatroom/list', component: ChatroomListComponent, canActivate: [LoginGuard] },
    { path: 'chatroom/create', component: ChatroomCreateComponent, canActivate: [LoginGuard] },

    { path: 'cronjob/list', component: CronjobListComponent, canActivate: [LoginGuard] },
    { path: 'cronjob/create', component: CronjobCreateComponent, canActivate: [LoginGuard] },
    { path: 'cronjob/update/:rd', component: CronjobUpdateComponent, canActivate: [LoginGuard] },

    { path: 'llmodel/create', component: LLModelCreateComponent, canActivate: [LoginGuard] },
    { path: 'llmodel/list', component: LLModelListComponent, canActivate: [LoginGuard] },

    { path: 'keyword/list', component: KeywordListComponent, canActivate: [LoginGuard] },
    { path: 'keyword/create', component: KeywordCreateComponent, canActivate: [LoginGuard] },

    { path: 'profile/list', component: ProfileListComponent, canActivate: [LoginGuard] },
    { path: 'profile/create', component: ProfileCreateComponent, canActivate: [LoginGuard] },
    { path: 'profile/update/:rd', component: ProfileUpdateComponent, canActivate: [LoginGuard] },

    { path: 'setting/list', component: SettingListComponent, canActivate: [LoginGuard] },
    { path: 'setting/update/:name', component: SettingUpdateComponent, canActivate: [LoginGuard] },

    { path: 'wcferry/chatroom', component: WcferryChatroomComponent, canActivate: [LoginGuard] },
    { path: 'wcferry/contact', component: WcferryContactComponent, canActivate: [LoginGuard] },
    { path: 'wcferry/receiver', component: WcferryReceiverComponent, canActivate: [LoginGuard] },

    { path: '', redirectTo: 'welcome', pathMatch: 'full' },
    { path: '**', component: Alert404Component },
];
