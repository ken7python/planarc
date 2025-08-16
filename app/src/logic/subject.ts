import { CONST } from '@/logic/const';
import { user } from '@/logic/user';

export const subjectModule = {
    api: CONST.api() + '/subject',
    getList: async function(){
        const subjects = await fetch(`${ subjectModule.api }/`,{headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${user.getToken()}`
        }}).then().catch(err => {console.error(err)});
        if (await subjects.ok) {
            return await subjects.json();
        }else {
            console.error('Failed to fetch subjects:', subjects.statusText);
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
        }).catch(err => console.error(err));
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
        fetch(`${subjectModule.api}/edit`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                id: ID,
                after_name: AfterName,
            })
        }).then(res => res.json()).then(data => { alert(data.message) }).catch(err => alert(err));
    },
}