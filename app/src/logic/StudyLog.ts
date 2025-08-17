import { ref } from 'vue';
import { CONST } from "@/logic/const";

const nullStr = "--";

export const stopwatch = {
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
        this.startDate.value = new Date();
        console.log(this.startDate.value);
        this.endDate.value = null;

        this.update();
    },
    stop: function() {
        if (this.startDate.value === null) {
            alert("ストップウォッチは開始していません")
        }
        this.endDate.value = new Date();

        this.update();
    },
}

export const studyLog = {
    write: function(sHours :number,sMinutes :number, eHours :number,eMinutes :number) {
        if (!sHours || !sMinutes || !eHours || !eMinutes) {
            alert("開始時間と終了時間を設定してください");
            return;
        }
        alert(`記録を送信する機能を実装予定です。勉強記録: ${sHours}:${sMinutes} - ${eHours}:${eMinutes}`);
    }
}