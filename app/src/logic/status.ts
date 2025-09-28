import { CONST } from "@/logic/const.ts";
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
    setEnjoyment: async function(date :string,enjoyment :string) {
        if (enjoyment === "") {
            alert("楽しみを入力してください");
            return null;
        }
        const res = await fetch(`${this.api}/enjoyment?date=${date}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${user.getToken()}`
            },
            body: JSON.stringify({
                "enjoyment": enjoyment
            })
        })
        if (res.ok) {
            alert("今日の楽しみを保存しました");
            return true;
        } else {
            console.error('Failed to fetch status:', res.statusText);
            alert("サーバとの通信に失敗しました");
            return null;
        }
    },
    setMood: async function(date :string,mood :number) {
        const res = await fetch(`${this.api}/mood?date=${date}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${user.getToken()}`
            },
            body: JSON.stringify({
                "mood": mood
            })
        })
        if (res.ok) {
            return true;
        } else {
            console.error('Failed to fetch status:', res.statusText);
            alert("サーバとの通信に失敗しました");
            return null;
        }
    },
}