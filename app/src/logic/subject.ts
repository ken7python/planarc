import { CONST } from '@/logic/const';
import { user } from '@/logic/user';

export const subjectModule = {
    api: CONST.api() + '/subject',
    getList: async function(){
        const subjects = await fetch(`${ subjectModule.api }/`,{headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${user.getToken()}`
        }}).then().catch(err => {alert(err)});
        if (await subjects.ok) {
            return await subjects.json();
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
        }).catch(err => alart(err));
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
        }).catch(err => alert(err));
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