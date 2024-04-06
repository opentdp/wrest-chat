import { LoginGuard } from '../helpers/login-guard';

import { WelcomeComponent } from './welcome';
import { Alert404Component } from './alert/404';

import { ChatroomListComponent } from './chatroom/list';
import { ChatroomCreateComponent } from './chatroom/create';
import { ChatroomUpdateComponent } from './chatroom/update';

import { CronjobListComponent } from './cronjob/list';
import { CronjobCreateComponent } from './cronjob/create';
import { CronjobUpdateComponent } from './cronjob/update';

import { HandlerListComponent } from './handler/list';

import { LLModelCreateComponent } from './llmodel/create';
import { LLModelListComponent } from './llmodel/list';
import { LLModelUpdateComponent } from './llmodel/update';

import { KeywordListComponent } from './keyword/list';
import { KeywordCreateComponent } from './keyword/create';
import { KeywordUpdateComponent } from "./keyword/update";

import { PluginCronjobComponent } from './plugin/cronjob';
import { PluginKeywordComponent } from './plugin/keyword';

import { ProfileListComponent } from './profile/list';
import { ProfileCreateComponent } from './profile/create';
import { ProfileUpdateComponent } from './profile/update';

import { SettingListComponent } from './setting/list';
import { SettingCreateComponent } from './setting/create';
import { SettingUpdateComponent } from './setting/update';

import { WcferryChatroomComponent } from './wcferry/chatroom';
import { WcferryContactComponent } from './wcferry/contact';
import { WcferryDbqueryComponent } from './wcferry/dbquery';
import { WcferryReceiverComponent } from './wcferry/receiver';

import { WebhookCreateComponent } from './webhook/create';
import { WebhookListComponent } from "./webhook/list";


export const AppComponents = [
    WelcomeComponent,
    Alert404Component,

    ChatroomListComponent,
    ChatroomCreateComponent,
    ChatroomUpdateComponent,

    CronjobListComponent,
    CronjobCreateComponent,
    CronjobUpdateComponent,

    HandlerListComponent,

    LLModelCreateComponent,
    LLModelListComponent,
    LLModelUpdateComponent,

    KeywordListComponent,
    KeywordCreateComponent,
    KeywordUpdateComponent,

    PluginCronjobComponent,
    PluginKeywordComponent,

    ProfileListComponent,
    ProfileCreateComponent,
    ProfileUpdateComponent,

    SettingListComponent,
    SettingCreateComponent,
    SettingUpdateComponent,

    WcferryChatroomComponent,
    WcferryContactComponent,
    WcferryDbqueryComponent,
    WcferryReceiverComponent,

    WebhookCreateComponent,
    WebhookListComponent
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

    { path: 'handler/list', component: HandlerListComponent, canActivate: [LoginGuard] },

    { path: 'llmodel/create', component: LLModelCreateComponent, canActivate: [LoginGuard] },
    { path: 'llmodel/list', component: LLModelListComponent, canActivate: [LoginGuard] },
    { path: 'llmodel/update/:rd', component: LLModelUpdateComponent, canActivate: [LoginGuard] },

    { path: 'keyword/list', component: KeywordListComponent, canActivate: [LoginGuard] },
    { path: 'keyword/create', component: KeywordCreateComponent, canActivate: [LoginGuard] },
    { path: 'keyword/update/:rd', component: KeywordUpdateComponent, canActivate: [LoginGuard] },

    { path: 'plugin/cronjob', component: PluginCronjobComponent, canActivate: [LoginGuard] },
    { path: 'plugin/keyword', component: PluginKeywordComponent, canActivate: [LoginGuard] },

    { path: 'profile/list', component: ProfileListComponent, canActivate: [LoginGuard] },
    { path: 'profile/create', component: ProfileCreateComponent, canActivate: [LoginGuard] },
    { path: 'profile/update/:rd', component: ProfileUpdateComponent, canActivate: [LoginGuard] },

    { path: 'setting/list', component: SettingListComponent, canActivate: [LoginGuard] },
    { path: 'setting/create', component: SettingCreateComponent, canActivate: [LoginGuard] },
    { path: 'setting/update/:rd', component: SettingUpdateComponent, canActivate: [LoginGuard] },

    { path: 'wcferry/chatroom', component: WcferryChatroomComponent, canActivate: [LoginGuard] },
    { path: 'wcferry/contact', component: WcferryContactComponent, canActivate: [LoginGuard] },
    { path: 'wcferry/dbquery', component: WcferryDbqueryComponent, canActivate: [LoginGuard] },
    { path: 'wcferry/receiver', component: WcferryReceiverComponent, canActivate: [LoginGuard] },

    { path: 'webhook/list', component: WebhookListComponent, canActivate: [LoginGuard] },
    { path: 'webhook/create', component: WebhookCreateComponent, canActivate: [LoginGuard] },

    { path: '', redirectTo: 'welcome', pathMatch: 'full' },
    { path: '**', component: Alert404Component },
];
