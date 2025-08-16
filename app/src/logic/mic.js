import {ref} from "vue";

export const mic = {
    ok: false,
    req: null,
    tmp: '',
    result: '',
    shouldRestart: ref(false), // 連打対策＆再開フラグ
    setup: function(){
        const SpeechRecognition = window.SpeechRecognition || window.webkitSpeechRecognition;
        if (!SpeechRecognition) {
            alert('このブラウザはWeb Speech APIに対応していません。Chrome系をおすすめします。');
        }else {
            mic.ok = true;
            mic.rec = new SpeechRecognition();
            mic.rec.lang = 'ja-JP';
            mic.rec.interimResults = true;   // 途中結果を受け取る
            mic.rec.continuous = true;       // 続けて聞く（ただし時々endが来る）

            // const finalEl = document.getElementById('final');
            // const interimEl = document.getElementById('interim');

            mic.result = '';
            mic.tmp = '';
        }
    },
    start: async function() {
        // 連打対策＆再開フラグ
        mic.shouldRestart = true;
        // mic.result = '';
        mic.tmp = ''
        try { mic.rec.start(); } catch (e) { /* 既に開始中など */ }
    },
    stop: function() {
        mic.shouldRestart = false;
        mic.rec.stop();
    },

    micON: ref(false),
    micStyle: function() {
        return {
            width: '40px',
            height: '40px',
            cursor: 'pointer',
            fill: mic.micON.value ? '#FF0000' : '#AEB7BD'
        }
    }
};

mic.setup()

mic.rec.onerror = (e) => {
    console.warn('Speech error:', e.error);
    // よくある: 'not-allowed', 'no-speech', 'audio-capture', 'aborted'
};

mic.shouldRestart = false;
mic.rec.onend = () => {
    if (mic.shouldRestart) {
        try { mic.rec.start(); } catch (_) {}
    }
};