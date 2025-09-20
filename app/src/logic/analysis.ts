import { CONST } from "./const";
import { user } from "./user";

export const analysisModule = {
    api: CONST.api() + '/analysis',
    get: async function(start: string, end: string) {
        const res = await fetch(`${this.api}/?start=${start}&end=${end}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${user.getToken()}`
            }
        })

        if (await res.ok) {
            const json = await res.json();
            if (json) {
                return json;
            } else {
                alert("データの取得に失敗しました");
                return null;
            }

        } else {
            alert("サーバとの通信に失敗しました");
            return null;
        }
    }
}