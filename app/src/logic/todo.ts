import { CONST } from "@/logic/const.ts";
import { user } from "./user";
import {ref} from "vue";
import { subjectModule } from "@/logic/subject";

export const todoModule = {
    api: CONST.api() + '/todo',
    getList: async function(date :string){
        const todos = await fetch(`${ this.api }/?date=${ date }`,{headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${await user.getToken()}`
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
    getListGroup: async function(date :string){
        const todos = await fetch(`${ this.api }/group?date=${ date }`,{headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${await user.getToken()}`
            }})
        if (await todos.ok) {
            const res = await todos.json();

            if (!res.MUST) {
                res.MUST = [];
            }
            if (!res.WANT) {
                res.WANT = [];
            }
            if (!res.checked) {
                res.checked = [];
            }

            res.MUST.reverse();
            res.WANT.reverse();
            res.checked.reverse();

            let subjects = ref<any[]>([]);

            const subject_list = await subjectModule.getList();
            //console.log(subject_list);
            subjects.value = subject_list;

            res.MUST.forEach((item) => {
                item["Color"] = subjects.value.find((subject) => subject.ID === item.SubjectID)?.Color || '#000000';
            })

            res.WANT.forEach((item) => {
                item["Color"] = subjects.value.find((subject) => subject.ID === item.SubjectID)?.Color || '#000000';
            })

            res.checked.forEach((item) => {
                item["Color"] = subjects.value.find((subject) => subject.ID === item.SubjectID)?.Color || '#000000';
            })
            return await res;
        }else {
            alert("サーバとの通信に失敗しました");
            return null;
        }
    },
    add: async function(date: string,title :string,suject :number,status :string){
        console.log(date, title, suject, status)
        if( !title ) {
            alert("タイトルを入力してください");
            return;
        }
        if( !suject ) {
            alert("科目を入力してください");
            return;
        }
        if( !status ) {
            alert("MUSTかWANTを入力してください");
            return;
        }
        if ( !date ) {
            alert("日付の入力に失敗しました");
            return;
        }
        const res = await fetch(`${this.api}/add`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${await user.getToken()}`
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
    edit: async function(ID: number, newTitle: string){
        const res = await fetch(`${this.api}/edit`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${await user.getToken()}`
            },
            body: JSON.stringify({
                "id": ID,
                "newtitle": newTitle,
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
    async check(ID: number) {
        return fetch(`${this.api}/check`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${await user.getToken()}`
            },
            body: JSON.stringify({
                "id": ID
            })
        }).then(res => {
            if (res.ok) {
                return res.json();
            } else {
                throw new Error('Failed to check ToDo');
            }
        });
    }
}
