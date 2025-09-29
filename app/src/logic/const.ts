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
        return this.dateToString(today);
    },
    dateToString: function(date :Date) :string{
        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, '0'); // Months are zero-based
        const day = String(date.getDate()).padStart(2, '0');
        return `${year}-${month}-${day}`;
    },
    timeToString: function(date :Date) :string {
        const hours = String(date.getHours()).padStart(2, '0');
        const minutes = String(date.getMinutes()).padStart(2, '0');
        return `${hours}:${minutes}`;
    }
}
