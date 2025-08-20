import { CONST } from "./const";
import { user } from "./user";

export const todoModule = {
    api: CONST.api() + '/todo',
    getList: async function(){
        const todos = await fetch(`${ this.api }/`,{headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${user.getToken()}`
            }})
        if (await todos.ok) {
            const res = await todos.json();
            res.reverse();
            return await res;
        }else {
            alert("サーバとの通信に失敗しました");
            return null;
        }
    },
    add: async function(date: string,title :string,suject :number,status :string){
        console.log(date, title, suject, status)
        const res = await fetch(`${this.api}/add`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${user.getToken()}`
            },
            body: JSON.stringify({
                "date": date,
                "title": title,
                "subjectID": suject,
                "status": status
            })
        })
        console.log(await res)
        if (await res.ok) {
            await this.getList()
        }else {
            console.error('Failed to fetch ToDos:', res.statusText);
            alert("サーバとの通信に失敗しました");
            return null;
        }
    },
    edit: async function(ID: number, AfterName: string){
        const res = await fetch(`${this.api}/edit`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${user.getToken()}`
            },
            body: JSON.stringify({
                "id": ID,
                "aftername": AfterName,
            })
        })
        if (res.ok) {
            await this.getList();
        }else {
            console.error('Failed to fetch ToDos:', res.statusText);
            alert(`サーバとの通信に失敗しました`);
            return null;
        }
    },
}