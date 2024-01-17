import { ChatroomsComponent } from './chatrooms';
import { ContactsComponent } from './contacts';
import { ReceiverComponent } from './receiver';
import { OwnerComponent } from './owner';

import { ErrorComponent } from './error';


export const AppComponents = [
    ChatroomsComponent,
    ContactsComponent,
    ReceiverComponent,
    OwnerComponent,
    ErrorComponent,
];

//////////////////////////////////////////////////////////////////

import { Routes } from '@angular/router';

export const AppRoutes: Routes = [
    { path: '', redirectTo: 'owner', pathMatch: 'full' },
    { path: 'chatrooms', component: ChatroomsComponent },
    { path: 'contacts', component: ContactsComponent },
    { path: 'receiver', component: ReceiverComponent },
    { path: 'owner', component: OwnerComponent },
    { path: '**', component: ErrorComponent, data: { error: 404 } }
];
