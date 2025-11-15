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
    add: async function(date: string,title :string,suject :string,status :string, nDatetime: string){
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
        const body = await res.json()
        console.log(body)
        console.log(body["id"])
        if (await res.ok) {
            if (nDatetime) {
                await this.register(nDatetime, body["id"]);
            }
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
    },
    publicVapidKey: "BKdgyFaYbmA8NNQvlHbr6TQ6wJudtWWzmlcDmPogbp9ppkRuvB7kQThDjVw0LDwjynesVAQvlRlFkdfMu45KO6g",
    register: async function(datetime: string = null, nTask: string = null) {
        console.log("=== プッシュ通知登録開始 ===");

        try {
            if ("serviceWorker" in navigator) {
                console.log("1. Service Worker対応確認 OK");

                // ビルド後に正しいパスで登録
                const swPath = import.meta.env.PROD ? '/sw.js' : `${import.meta.env.BASE_URL}sw.js`;
                console.log("Service Worker登録パス:", swPath);

                const reg = await navigator.serviceWorker.register(swPath);
                console.log("2. Service Worker 登録完了:", reg);

                // 通知許可の確認
                console.log("3. 通知許可確認開始");
                const permission = await Notification.requestPermission();
                console.log("4. 通知許可結果:", permission);

                if (permission !== "granted") {
                    console.error("❌ 通知が許可されていません");
                    alert("通知が許可されていません。ブラウザの設定で通知を許可してください。");
                    return;
                }

                // プッシュマネージャーの確認
                if (!reg.pushManager) {
                    console.error("❌ プッシュマネージャーが利用できません");
                    alert("このブラウザはプッシュ通知に対応していません");
                    return;
                }

                console.log("5. プッシュ購読開始");
                console.log("VAPIDキー:", this.publicVapidKey.substring(0, 20) + "...");

                const sub = await reg.pushManager.subscribe({
                    userVisibleOnly: true,
                    applicationServerKey: this.urlBase64ToUint8Array(this.publicVapidKey),
                });
                console.log("6. プッシュ購読完了:", sub);
                console.log("購読エンドポイント:", sub.endpoint);

                console.log("7. サーバー送信開始");
                const apiUrl = `${CONST.api()}/notify/send`;
                console.log("API URL:", apiUrl);

                console.log(await user.getToken());

                const response = await fetch(apiUrl, {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                        Authorization: `Bearer ${await user.getToken()}`,
                        "datetime": datetime || "",
                        "Task": nTask || ""
                    },
                    body: JSON.stringify(sub),
                });
                console.log(response);

                console.log("8. サーバーレスポンス:", response.status, response.statusText);

                if (response.ok) {
                    console.log("✅ 通知登録完了");
                    alert("通知登録完了🎉");
                } else {
                    const errorText = await response.text();
                    console.error("❌ サーバーエラー:", response.status, errorText);
                    alert(`サーバーエラー: ${response.status}\n詳細: ${errorText}\n apiUrl: ${apiUrl}`);
                }
            } else {
                console.error("❌ Service Worker非対応");
                alert("このブラウザはService Workerに対応していません");
            }
        } catch (error) {
            console.error("❌ 通知登録エラー:", error);
            console.error("エラー詳細:", error.stack);
            alert(`エラーが発生しました: ${error.message}`);
        }
    },

    urlBase64ToUint8Array: function (base64String) {
        const padding = "=".repeat((4 - (base64String.length % 4)) % 4);
        const base64 = (base64String + padding).replace(/-/g, "+").replace(/_/g, "/");
        const rawData = atob(base64);
        return Uint8Array.from([...rawData].map((c) => c.charCodeAt(0)));
    }
}
