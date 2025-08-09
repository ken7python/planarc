// const.js

export const CONST = {
    homeLink: '/',
    dailyLink: '/daily',
    flagLink: '/flag',
    authLink: '/auth',
    signupLink: '/signup',
    loginLink: '/login',

    api: function(){
        return 'http://localhost:8080/api';
    },
    account_url: function() {
        return `${CONST.api()}/accounts`;
    },
}