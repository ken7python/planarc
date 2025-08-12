// const.js
import { api } from '@/logic/api.js';

export const CONST = {
    homeLink: '/',
    dailyLink: '/daily',
    flagLink: '/flag',
    authLink: '/auth',
    signupLink: '/signup',
    loginLink: '/login',

    api: function(){
        return api;
    },
    account_url: function() {
        return `${CONST.api()}/accounts`;
    },
}
