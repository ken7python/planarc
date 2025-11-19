import {ref} from "vue";

export const mic = {
    ok: false,
    rec: null,
    tmp: '',
    result: '',
    shouldRestart: ref(false),     // ← refのまま使う
    micON: ref(false),

    setup() {
        const SpeechRecognition = window.SpeechRecognition || window.webkitSpeechRecognition;
        if (!SpeechRecognition) {
            alert('このブラウザはWeb Speech APIに対応していません。Chrome系をおすすめします。');
            return;
        }
        this.ok = true;
        this.rec = new SpeechRecognition();
        this.rec.lang = 'ja-JP';
        this.rec.interimResults = true;
        this.rec.continuous = true;
        this.rec.maxAlternatives = 1;

        this.result = '';
        this.tmp = '';

        // エラー
        this.rec.onerror = (e) => {
            console.warn('Speech error:', e.error);
            // Android対策：no-speech は再開して握りつぶす
            if (this.shouldRestart.value && (e.error === 'no-speech' || e.error === 'aborted')) {
                try { this.rec.start(); } catch (_) {}
            }
        };

        // 予期せぬ終了→再開
        this.rec.onend = () => {
            if (this.shouldRestart.value) {
                try { this.rec.start(); } catch (_) {}
            }
        };
    },

    async start() {
        if (!this.ok) return;

        // ★ Android安定化：権限を確実に取得（HTTPS前提）
        try {
            const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
            // 使わないので即停止
            stream.getTracks().forEach(t => t.stop());
        } catch (e) {
            alert('マイク権限が必要です（設定でChromeのマイクを許可してね）');
            return;
        }
        mic.micON.value = true;
        this.shouldRestart.value = true;   // ← .value
        this.tmp = '';
        try { this.rec.start(); } catch (e) { /* 連打でAlready started等 */ }
    },

    stop() {
        this.shouldRestart.value = false;  // ← .value
        mic.micON.value = false;
        try { this.rec.stop(); } catch (_) {}
    },

    micStyle(wantRed) {
        return {
            width: '30px',
            height: '30px',
            cursor: 'pointer',
            fill: this.micON.value && wantRed ? '#FF0000' : '#AEB7BD'
        };
    }
};

mic.setup();

mic.rec.onerror = (e) => {
    console.warn('Speech error:', e.error);
    // よくある: 'not-allowed', 'no-speech', 'audio-capture', 'aborted'
};

mic.shouldRestart.value = false;
mic.rec.onend = () => {
    if (mic.shouldRestart.value) {
        try { mic.rec.start(); } catch (_) {}
    }
};