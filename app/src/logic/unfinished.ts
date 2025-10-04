import { CONST } from "@/logic/const.ts";
import { user } from "./user";
import { subjectModule } from "@/logic/subject";

export const unfinishedModule = {
    api: CONST.api() + '/unfinished',
    getList: async function(date :string){
        const lists = await fetch(`${ this.api }/`,{headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${await user.getToken()}`
            }})
        if (await lists.ok) {
            const res = await lists.json();
            res.reverse();
            return await res;
        }else {
            alert("サーバとの通信に失敗しました");
            return null;
        }
    },
    move: async function(id :number){
        const res = await fetch(`${this.api}/move`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${await user.getToken()}`
            },
            body: JSON.stringify({
                "id": id
            })
        })
        console.log(await res)
        if (res.ok) {
            return true;
        }else {
            console.error('Failed to fetch unfinished:', res.statusText);
            alert("サーバとの通信に失敗しました");
            return null;
        }
    },
    delete: async function(id :number){
        const res = await fetch(`${this.api}/delete`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${await user.getToken()}`
            },
            body: JSON.stringify({
                "id": id
            })
        })
        console.log(await res)
        if (res.ok) {
            return true;
        }else {
            console.error('Failed to fetch unfinished:', res.statusText);
            alert("サーバとの通信に失敗しました");
            return null;
        }
    },
    back: async function(id :number, dateStr? :string){
        const res = await fetch(`${this.api}/back`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${await user.getToken()}`
            },
            body: JSON.stringify({
                "id": id,
                "date": dateStr
            })
        })
        console.log(await res)
        if (res.ok) {
            return true;
        }else {
            console.error('Failed to fetch unfinished:', res.statusText);
            alert("サーバとの通信に失敗しました");
            return null;
        }
    },
}