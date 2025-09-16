import { ref } from 'vue';
import { CONST } from "./const";
import { user } from "./user";

const nullStr = "--";

export const stopwatch = {
    subject : ref<number>(null),
    startDate: ref(null),
    endDate: ref(null),

    DateDiff: {
        hours: ref(null),
        minutes: ref(null),
    },

    sHours: ref(null),
    sMinutes: ref(null),
    eHours: ref(null),
    eMinutes: ref(null),
    dHours: ref(null),
    dMinutes: ref(null),

    init: function() {
        this.startDate.value = null;
        this.endDate.value = null;
        this.DateDiff.value = null;

        this.sHours.value = null;
        this.sMinutes.value = null;
        this.eHours.value = null;
        this.eMinutes.value = null;
        this.dHours.value = null;
        this.dMinutes.value = null;

        this.load();
        this.update();
    },
    saveRef: function(key, refitem) {
        //console.log("saveRef", key, refitem);
        if (refitem.value != null) {
            localStorage.setItem(key, refitem.value)
        }
    },

    getItem: function(key :string) {
        return localStorage.getItem(key);
    },

    save: function() {
        // localStorage.setItem("sHours", this.sHours.value || null);
        // localStorage.setItem("sMinutes", this.sMinutes.value || null);
        // localStorage.setItem("eHours", this.eHours.value || null);
        // localStorage.setItem("eMinutes", this.eMinutes.value || null);
        this.saveRef("startDate", this.startDate);
        this.saveRef("endDate", this.endDate);
        this.saveRef("sHours", this.sHours);
        this.saveRef("sMinutes", this.sMinutes);
        this.saveRef("eHours", this.eHours);
        this.saveRef("eMinutes", this.eMinutes);
        this.saveRef("subject", this.subject);
        // console.log(this.subject);
        console.log("Saved stopwatch state to localStorage");
    },
    load: function() {
        this.sHours.value = this.getItem("sHours");
        this.sMinutes.value = this.getItem("sMinutes");
        this.eHours.value = this.getItem("eHours");
        this.eMinutes.value = this.getItem("eMinutes");
        const startDate = this.getItem("startDate");
        const endDate = this.getItem("endDate");
        if (startDate) {
            this.startDate.value = new Date(startDate);
        }
        if (endDate) {
            this.endDate.value = new Date(endDate);
        }

        this.subject.value = this.getItem("subject") || '';
        // console.log(this.subject.value);

        this.update();
    },

    update: function() {
        const sd = this.startDate.value;
        const ed = this.endDate.value;

        let dMinutes = this.dMinutes
        let dHours = this.dHours;

        let eMinutes = this.eMinutes;
        let eHours = this.eHours;
        let sMinutes = this.sMinutes;
        let sHours = this.sHours;
        if (sd) {
            sMinutes.value = this.startDate.value.getMinutes();
            sHours.value = this.startDate.value.getHours();
        }
        if (ed) {
            eMinutes.value = this.endDate.value.getMinutes();
            eHours.value = this.endDate.value.getHours();
        }
        if (sd) {
            let diff;
            const tStart = sMinutes.value + sHours.value * 60;
            if (ed) {
                const tStop = eMinutes.value + eHours.value * 60;
                diff = tStop - tStart;
            }else {
                const now = new Date();
                const nMinute = now.getMinutes();
                const nHour = now.getHours();
                const tNow = nMinute + nHour * 60;
                diff = tNow - tStart;
            }
            dMinutes.value = diff % 60;
            dHours.value = Math.floor(diff / 60);
        }
        // console.log(this);
    },

    start: function() {
        this.reset();
        this.startDate.value = new Date();
        console.log(this.startDate.value);
        this.endDate.value = null;

        this.update();
        this.save();
    },
    stop: function() {
        if (this.startDate.value === null) {
            alert("ストップウォッチは開始していません")
        }
        this.endDate.value = new Date();

        this.update();
        this.save();
    },
    reset: function() {
        this.startDate.value = null;
        this.endDate.value = null;

        this.DateDiff.hours.value = null;
        this.DateDiff.minutes.value = null;

        this.sHours.value = null;
        this.sMinutes.value = null;
        this.eHours.value = null;
        this.eMinutes.value = null;
        this.dHours.value = null;
        this.dMinutes.value = null;

        localStorage.removeItem("sHours");
        localStorage.removeItem("sMinutes");
        localStorage.removeItem("eHours");
        localStorage.removeItem("eMinutes");
        localStorage.removeItem("startDate");
        localStorage.removeItem("endDate");
    }
}

export const studyLog = {
    api: CONST.api() + '/studylog',
    separate: function(time :string) {
        if (!time || time === nullStr) {
            return { hours: null, minutes: null };
        }
        const parts = time.split(':');
        if (parts.length !== 2) {
            throw new Error("Invalid time format. Expected format is 'HH:MM'.");
        }
        const hours = parseInt(parts[0], 10);
        const minutes = parseInt(parts[1], 10);
        if (isNaN(hours) || isNaN(minutes)) {
            throw new Error("Invalid time format. Hours and minutes must be numbers.");
        }
        return { hours, minutes };
    },
    isNull(value :any) {
        return value === null;
    },

    write: async function(date :string,subject :number,sHours :number,sMinutes :number, eHours :number,eMinutes :number) {
        console.log(subject);
        console.log(sHours);
        console.log(sMinutes);
        console.log(eHours);
        console.log(eMinutes);

        if (this.isNull(sHours) || this.isNull(sMinutes)) {
            alert("開始時間を入力してください");
            return false;
        }
        if ( this.isNull(eHours) || this.isNull(eMinutes) ){
            alert("終了時間を入力してください");
            return false;
        }
        if (!subject) {
            alert("科目を入力してください");
            return false;
        }

        // alert(`科目ID:${subject}勉強時間: ${sHours}時${sMinutes}分〜${eHours}時${eMinutes}分`);

        const res = await fetch(`${this.api}/add`,{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${user.getToken()}`
            },
            body: JSON.stringify({
                "date": date,
                "sHours": sHours,
                "sMinutes": sMinutes,
                "eHours": eHours,
                "eMinutes": eMinutes,
                "subjectID": Number(subject)
            })
        })
        console.log(await res)
        if (res.ok) {
            stopwatch.reset();
            stopwatch.init();
            alert("勉強記録を追加しました");
            return true;
        }else {
            console.error('Failed to fetch study_logs:', res.statusText);
            alert("サーバとの通信に失敗しました");
            return false;
        }
    },
    getLog: async function(dateStr) {
        const res = await fetch(`${this.api}/?date=${dateStr}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${user.getToken()}`
            },
        }).catch(err => {
            alert(err);
            return null;
        });
        if (!res.ok) {
            throw new Error("Failed to fetch study logs");
        }
        return await res.json();
    },
    writeStr: async function(date: string,subject :number, startTime :string, endTime :string) {
        // console.log(startTime, endTime);
        const start = this.separate(startTime);
        const end = this.separate(endTime);
        // console.log(start, end);
        const ok = await this.write(date, subject, start.hours, start.minutes, end.hours, end.minutes);
        // console.log(res);
        return ok;
    }
}
