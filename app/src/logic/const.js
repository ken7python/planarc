// const.js
import { api } from '@/logic/api.js';

export const CONST = {
    progressLink: '/progress',
    dailyLink: '/',
    flagLink: '/flag',
    authLink: '/auth',
    signupLink: '/signup',
    loginLink: '/login',
    subjectLink: '/subject',
    analysisLink: '/analysis',

    api: function(){
        return api;
    },
    account_url: function() {
        return `${CONST.api()}/accounts`;
    },

    getToday: function() {
        const today = new Date();
        const year = today.getFullYear();
        const month = String(today.getMonth() + 1).padStart(2, '0'); // Months are zero-based
        const day = String(today.getDate()).padStart(2, '0');
        return `${year}-${month}-${day}`;
    }
}
