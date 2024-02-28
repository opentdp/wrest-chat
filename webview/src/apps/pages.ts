import { LoginGuard } from '../helpers/login-guard';

import { WelcomeComponent } from './welcome';
import { ErrorComponent } from './error';

import { BotChatroomsComponent } from './robot/chatrooms';
import { BotKeywordsComponent } from './robot/keywords';
import { BotProfilesComponent } from './robot/profiles';

import { WcfChatroomsComponent } from './wcferry/chatrooms';
import { WcfContactsComponent } from './wcferry/contacts';
import { WcfReceiverComponent } from './wcferry/receiver';


export const AppComponents = [
    WelcomeComponent,
    ErrorComponent,

    BotChatroomsComponent,
    BotKeywordsComponent,
    BotProfilesComponent,

    WcfChatroomsComponent,
    WcfContactsComponent,
    WcfReceiverComponent,
];

//////////////////////////////////////////////////////////////////

import { Routes } from '@angular/router';

export const AppRoutes: Routes = [
    { path: 'welcome', component: WelcomeComponent },

    { path: 'bot/chatrooms', component: BotChatroomsComponent, canActivate: [LoginGuard] },
    { path: 'bot/keywords', component: BotKeywordsComponent, canActivate: [LoginGuard] },
    { path: 'bot/profiles', component: BotProfilesComponent, canActivate: [LoginGuard] },

    { path: 'wcf/chatrooms', component: WcfChatroomsComponent, canActivate: [LoginGuard] },
    { path: 'wcf/contacts', component: WcfContactsComponent, canActivate: [LoginGuard] },
    { path: 'wcf/receiver', component: WcfReceiverComponent, canActivate: [LoginGuard] },

    { path: '', redirectTo: 'welcome', pathMatch: 'full' },
    { path: '**', component: ErrorComponent, data: { error: 404 } },
];
