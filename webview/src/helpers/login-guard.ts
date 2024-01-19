import { CanActivateFn } from '@angular/router';

import { WrestApi } from '../openapi/wcfrest';


export const LoginGuard: CanActivateFn = () => {
   return WrestApi.isLogin();
};

