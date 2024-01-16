import { ChatroomsComponent } from './chatrooms';
import { ContactsComponent } from './contacts';
import { OwnerComponent } from './owner';

import { ErrorComponent } from './error';


export const AppComponents = [
    ChatroomsComponent,
    ContactsComponent,
    OwnerComponent,
    ErrorComponent,
];

//////////////////////////////////////////////////////////////////

import { Routes } from '@angular/router';

export const AppRoutes: Routes = [
    { path: '', redirectTo: 'owner', pathMatch: 'full' },
    { path: 'chatrooms', component: ChatroomsComponent },
    { path: 'contacts', component: ContactsComponent },
    { path: 'owner', component: OwnerComponent },
    { path: '**', component: ErrorComponent, data: { error: 404 } }
];
