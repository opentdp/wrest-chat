import { CanActivateFn } from '@angular/router';

import { WrestApi } from '../openapi/wcfrest';


export const LoginGuard: CanActivateFn = async () => {

   const isLogin = await WrestApi.isLogin();

   isLogin || location.assign('/#/welcome');
   return isLogin;

};
