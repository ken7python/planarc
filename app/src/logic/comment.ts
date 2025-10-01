import { CONST } from "@/logic/const.ts";
import { user } from "./user";
import { ref } from "vue";

export const CommentModule = {
    api: CONST.api() + '/comment',
    refComment: ref(""),
    refUserNote: ref(""),
    ask: async function(date :string, note :string,chr :string){
        if (chr === "") {
            alert("キャラを選択してください");
            return null;
        }
        if (note === "") {
            alert("日記が空です");
            return null;
        }

        this.refComment.value = "通信中...";

        const res = await fetch(`${ this.api }/ask`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${user.getToken()}`
            },
            body: JSON.stringify({
                "date": date,
                "note": note,
                "chr": chr
            })
        })

        const reader = res.body.getReader();
        const decoder = new TextDecoder();

        if (await res.ok) {
            while (true) {
                const { done, value } = await reader.read();
                if (done) break;
                const decodedValue = decoder.decode(value);
                this.refComment.value += decodedValue;
                console.log(decodedValue);
            }
            return true;
        }
        else {
            alert("サーバとの通信に失敗しました");
            return null;
        }
    },
    get: async function(date :string, note :string){
        const res = await fetch(`${ this.api }/?date=${date}`,{
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${user.getToken()}`
            }
        })

        if (await res.ok) {
            const json = await res.json();
            if (json) {
                // console.log(json);
                this.refComment.value = json.Note;
                this.refUserNote.value = json.UserNote;
                return json;
            }else {
                this.refComment.value = "";
                this.refUserNote.value = "";
            }

        }
    }
}