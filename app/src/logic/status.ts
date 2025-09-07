import { CONST } from "./const";
import { user } from "./user";

export const statusModule = {
    api: CONST.api() + '/status',
    get: async function(date :string){
        const res = await fetch(`${ this.api }/?date=${date}`,{headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${user.getToken()}`
            }})
        if (await res.ok) {
            return await res.json();
        }else {
            alert("サーバとの通信に失敗しました");
            return null;
        }
    },
    set: async function(status :string) {
        const res = await fetch(`${this.api}/set`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${user.getToken()}`
            },
            body: JSON.stringify({
                "status": status
            })
        })
        if (res.ok) {
            return true;
        } else {
            console.error('Failed to fetch status:', res.statusText);
            alert("サーバとの通信に失敗しました");
            return null;
        }
    }
}