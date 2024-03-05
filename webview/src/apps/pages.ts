import { LoginGuard } from '../helpers/login-guard';

import { WelcomeComponent } from './welcome';
import { Alert404Component } from './alert/404';

import { ChatroomListComponent } from './chatroom/list';
import { ChatroomCreateComponent } from './chatroom/create';
import { ChatroomUpdateComponent } from './chatroom/update';

import { CronjobListComponent } from './cronjob/list';
import { CronjobCreateComponent } from './cronjob/create';
import { CronjobUpdateComponent } from './cronjob/update';

import { LLModelCreateComponent } from './llmodel/create';
import { LLModelListComponent } from './llmodel/list';
import { LLModelUpdateComponent } from './llmodel/update';

import { KeywordListComponent } from './keyword/list';
import { KeywordCreateComponent } from './keyword/create';
import { KeywordUpdateComponent } from "./keyword/update";

import { ProfileListComponent } from './profile/list';
import { ProfileCreateComponent } from './profile/create';
import { ProfileUpdateComponent } from './profile/update';

import { SettingListComponent } from './setting/list';
import { SettingCreateComponent } from './setting/create';
import { SettingUpdateComponent } from './setting/update';

import { WcferryChatroomComponent } from './wcferry/chatroom';
import { WcferryContactComponent } from './wcferry/contact';
import { WcferryReceiverComponent } from './wcferry/receiver';


export const AppComponents = [
    WelcomeComponent,
    Alert404Component,

    ChatroomListComponent,
    ChatroomCreateComponent,
    ChatroomUpdateComponent,

    CronjobListComponent,
    CronjobCreateComponent,
    CronjobUpdateComponent,

    LLModelCreateComponent,
    LLModelListComponent,
    LLModelUpdateComponent,

    KeywordListComponent,
    KeywordCreateComponent,
    KeywordUpdateComponent,

    ProfileListComponent,
    ProfileCreateComponent,
    ProfileUpdateComponent,

    SettingListComponent,
    SettingCreateComponent,
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
    { path: 'chatroom/update/:rd', component: ChatroomUpdateComponent, canActivate: [LoginGuard] },

    { path: 'cronjob/list', component: CronjobListComponent, canActivate: [LoginGuard] },
    { path: 'cronjob/create', component: CronjobCreateComponent, canActivate: [LoginGuard] },
    { path: 'cronjob/update/:rd', component: CronjobUpdateComponent, canActivate: [LoginGuard] },

    { path: 'llmodel/create', component: LLModelCreateComponent, canActivate: [LoginGuard] },
    { path: 'llmodel/list', component: LLModelListComponent, canActivate: [LoginGuard] },
    { path: 'llmodel/update/:rd', component: LLModelUpdateComponent, canActivate: [LoginGuard] },

    { path: 'keyword/list', component: KeywordListComponent, canActivate: [LoginGuard] },
    { path: 'keyword/create', component: KeywordCreateComponent, canActivate: [LoginGuard] },
    { path: 'keyword/update/:rd', component: KeywordUpdateComponent, canActivate: [LoginGuard] },

    { path: 'profile/list', component: ProfileListComponent, canActivate: [LoginGuard] },
    { path: 'profile/create', component: ProfileCreateComponent, canActivate: [LoginGuard] },
    { path: 'profile/update/:rd', component: ProfileUpdateComponent, canActivate: [LoginGuard] },

    { path: 'setting/list', component: SettingListComponent, canActivate: [LoginGuard] },
    { path: 'setting/create', component: SettingCreateComponent, canActivate: [LoginGuard] },
    { path: 'setting/update/:rd', component: SettingUpdateComponent, canActivate: [LoginGuard] },

    { path: 'wcferry/chatroom', component: WcferryChatroomComponent, canActivate: [LoginGuard] },
    { path: 'wcferry/contact', component: WcferryContactComponent, canActivate: [LoginGuard] },
    { path: 'wcferry/receiver', component: WcferryReceiverComponent, canActivate: [LoginGuard] },

    { path: '', redirectTo: 'welcome', pathMatch: 'full' },
    { path: '**', component: Alert404Component },
];
