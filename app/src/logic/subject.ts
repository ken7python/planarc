import { CONST } from "@/logic/const.ts";
import { user } from './user';

export const subjectModule = {
    api: CONST.api() + '/subject',
    getList: async function(){
        const subjects = await fetch(`${ subjectModule.api }/`,{headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${user.getToken()}`
        }})
        if (await subjects.ok) {
            const res = await subjects.json();
            res.reverse()
            return await res;
        }else {
            alert("サーバとの通信に失敗しました");
            return null;
        }
    },
    add: async function(Name: string,Color: string){
        console.log(Name, Color)
        const res = await fetch(`${subjectModule.api}/add`,{
            method: 'POST',
            headers: {
               'Content-Type': 'application/json',
                Authorization: `Bearer ${user.getToken()}`
            },
            body: JSON.stringify({
                "name": Name,
                "color": Color
            })
        })
        console.log(await res)
        if (res.ok) {
            await subjectModule.getList()
        }else {
            console.error('Failed to fetch subjects:', res.statusText);
            alert("サーバとの通信に失敗しました");
            return null;
        }
    },
    edit: async function(ID: number, AfterName: string){
        const res = await fetch(`${subjectModule.api}/edit`,{
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
            await subjectModule.getList();
        }else {
            console.error('Failed to fetch subjects:', res.statusText);
            const err :string = await res.json();
            if (err){
                alert(err.error)
            }else {
                alert(`サーバとの通信に失敗しました`);
            }
            return null;
        }
    },
}