// let account_url = "http://localhost:8080/api/accounts";
import { CONST } from '@/logic/const.js';

export const user = {
    account_url: CONST.account_url(),
    profile: async function(){
        const token = localStorage.getItem('token')
        try {
            const response = await fetch(this.account_url + "/profile", {
                headers: {
                    Authorization: `Bearer ${token}`
                }
            });
            const data = await response.json();
            console.log(data);
            if (response.ok) {
                return data;
            }else{
                user.directToLogin();
            }
        } catch (error) {
            // alert('ユーザー情報の取得に失敗しました')
            console.error('ユーザー情報の取得に失敗しました:', error);
            user.directToLogin();
        }
    },

    logout: function() {
        localStorage.removeItem('token');
        user.directToLogin();
    },

    login: async function(username, password) {
        const response = await fetch(this.account_url + "/login", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username: username, password: password })
        });

        const data = await response.json();
        if (response.ok){
            console.log(data);
            localStorage.setItem('token', data.token);
            //alert("ログイン成功！トップページへ移動します。");
            return {"ok": response.ok,"data":data};
        }else{
            console.error(data.error);
            alert(data.error);
            return {"ok": response.ok,"error":data.error};
        }
    },

    register: async function(username, password) {
        var error_message;
        const response = await fetch(this.account_url + "/register", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                username: username,
                password: password,
            }),
        });

        const data = await response.json();
        //console.log(data);
        if (response.ok){
            await this.login(username, password);
            return {"ok":true, "data":data};
        }else{
            error_message = data.error;
            console.error(error_message);
            return {"ok":false, "error":error_message};
        }
    },

    directToLogin: function(){
        window.location.href = CONST.authLink;
    }
}