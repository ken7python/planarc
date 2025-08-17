import { ref } from 'vue';
import { CONST } from "@/logic/const";

const stopwatch = {
    startTime: ref(null),
    endTime: ref(null),
    start: function() {
        this.startTime.value = new Date();
        this.endTime.value = null;
    },
    end: function() {
        if (this.startTime.value === null) {
            alert("ストップウォッチは開始していません")
        }
        this.endTime.value = new Date();
    }
}