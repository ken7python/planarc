<script setup lang="ts">
  import commentIcon from '@/assets/icons/send.svg';
  import MicIcon from '@/assets/icons/mic.svg';
  import { ref, onMounted, onBeforeUnmount } from 'vue';
  import { mic } from '@/logic/mic';

  import { CommentModule } from "../logic/comment";
  
  const message = ref<string>("");

  async function micbtn() {
    // alert('マイクボタンが押されました。');
    if (!mic.shouldRestart.value) {
      console.log("start");
      mic.result = message.value;
      // mic.micON.value = true;
      await mic.start();
    } else {
      console.log("stop");
      // mic.micON.value = false;
      await mic.stop();
    }
  }

  mic.rec.onresult = (e) => {
    let interim = '';
    for (let i = e.resultIndex; i < e.results.length; i++) {
      const transcript = e.results[i][0].transcript;
      if (e.results[i].isFinal) {
        mic.result = transcript;
        mic.tmp = '';
        message.value = mic.result;
      } else {
        interim += transcript;
        mic.tmp = interim;
        message.value = mic.result;
      }
    }
  };

  const props = defineProps({
    date: String
  })

  const disabled = ref(false);
  const noteHeight = ref('auto');
  const memoBottom = ref('80px');
  const NOTE_LAYOUT_OFFSET = 70 + 60 + 30;

  const updateViewportMetrics = () => {
    const viewport = window.visualViewport;
    const viewportHeight = viewport ? viewport.height : window.innerHeight;
    const keyboardOffset = Math.max(0, window.innerHeight - viewportHeight);

    noteHeight.value = `${Math.max(viewportHeight - NOTE_LAYOUT_OFFSET, 160)}px`;
    memoBottom.value = `${Math.max(16, 80 + keyboardOffset)}px`;
  };

  onMounted(() => {
    updateViewportMetrics();
    window.addEventListener('resize', updateViewportMetrics);
    if (window.visualViewport) {
      window.visualViewport.addEventListener('resize', updateViewportMetrics);
      window.visualViewport.addEventListener('scroll', updateViewportMetrics);
    }
  });

  onBeforeUnmount(() => {
    window.removeEventListener('resize', updateViewportMetrics);
    if (window.visualViewport) {
      window.visualViewport.removeEventListener('resize', updateViewportMetrics);
      window.visualViewport.removeEventListener('scroll', updateViewportMetrics);
    }
  });
  CommentModule.refComment.value = "";

  const chr = ref("");

  async function ask() {
    disabled.value = true;
    await mic.stop();
    if (!(message.value === "") && !(chr.value === "")) {
      CommentModule.refUserNote.value = message.value;
      // CommentModule.refComment.value = "通信中...";
    }

    const comment = await CommentModule.ask(props.date,message.value,chr.value);
    console.log(comment);
    if (await comment) {
      // message.value = '';
      await loadData();
    }else {
      disabled.value = false;
    }
    console.log(comment);
    // ここでcommentを表示する処理を追加
  }

  async function loadData() {
    const comment = await CommentModule.get(props.date);
    console.log(comment);
  }
  loadData();
</script>

<template>
  <div id="note" :style="{ height: noteHeight }">
    <div id="mine" v-show="CommentModule.refUserNote.value">
<!--    <div id="mine">-->
<!--      <div v-if="CommentModule.refUserNote.value">-->
        <span class="yourComment">
          {{ CommentModule.refUserNote }}
        </span>
<!--      </div>-->
<!--      <div v-else>-->
<!--        <span>ひとこと</span>-->
<!--      </div>-->
    </div>

    <div id="comment" v-show="CommentModule.refComment.value">
<!--      <div v-if="CommentModule.refComment.value">-->
        <span>{{ CommentModule.refComment }}</span>
<!--      </div>-->
<!--      <div v-else>-->
<!--        <span>AIからのコメント</span>-->
<!--      </div>-->
    </div>

    <div v-if="!CommentModule.refComment.value && !CommentModule.refUserNote.value" id="about-note">
      <span>感想を書いてAIからメッセージをもらおう！</span>
    </div>

    <div class="micdiv" id="memo-input" v-if="!CommentModule.refComment.value" :style="{ bottom: memoBottom }">
      <div class="input-area">
        <select class="selectbox" v-model="chr">
          <option v-if="chr===''" disabled value="">AIサポーターを選択</option>
          <option value="先生">先生</option>
          <option value="親">親</option>
          <option value="友達">友達</option>
        </select>
        <textarea placeholder="感想" v-model="message"></textarea>
      </div>
      <div>
        <div>
          <mic-icon class="mic" :style="mic.micStyle(true)" @click="micbtn"></mic-icon>
        </div>
        <div v-if="message.length > 0">
          <button id="send-btn" style="margin: 0 auto;" @click="ask" :disabled="disabled">
            <commentIcon id="send-icon"></commentIcon>
          </button>
        </div>
        <div v-else>
          <button id="send-btn" style="margin: 0 auto;" disabled>
            <commentIcon id="send-icon" style="opacity: 0;"></commentIcon>
          </button>
        </div>
      </div>
    </div>

  </div>
</template>

<style scoped>
  html, body {
    margin: 0;
    height: 100dvh;
    overflow: hidden; /* ← ページ全体はスクロールさせない */
  }

  #note {
    text-align: center;
    overflow-y: auto;
  }

  #about-note {
    text-align: center;
    padding-top: 10px;
  }

  #memo-input {
    position: fixed;
    width: 90vw;
  }

  .input-area {
    display: flex;
    flex-direction: column;
    gap: 8px;
    margin-right: 8px;
    margin-bottom: 8px;
  }
  textarea {
    width: 100%;
    height: 22dvh;
    min-height: 80px;
    max-height: 120px;
    background-color: #FFFFFF;
    padding: 10px;
    font-size: 16px;
    border: 1px solid #E2E8F8;
    border-radius: 4px;
    box-sizing: border-box;
    resize: none;
  }
  
  .mic, #send-btn {
    width: 40px;
    height: 40px;
  }

  #send-icon {
    color: #4071e9;
  }

  #send-btn {
    background: none;
    border: none;
    cursor: pointer;
  }

  #mine, #comment {
    width: 75%;
    margin: 4px auto;
    padding: 10px;
    font-size: 16px;
    margin-bottom: 16px;
    border: 1px solid #E2E8F8;
    border-radius: 4px;
    clear: both;
    word-wrap: break-word;
    white-space: pre-wrap;
    box-sizing: border-box;
    min-height: 50px;
    max-width: 90%;
    text-align: left;

    display: inline-block;
    position: relative;
  }

  #mine::after, #comment::after {
    content: "";
    position: absolute;
    top: 100%;  /* 本体の下に配置 */
    border-width: 20px;
    border-style: solid;
    /* 上だけ背景色、それ以外は透明 */
  }
  #mine::after {
    border-color: #93df83 transparent transparent transparent;
    right: 20px; /* 左から20pxの位置に */
  }
  #comment::after {
    border-color: #E2E8F8 transparent transparent transparent;
    left: 20px; /* 左から20pxの位置に */
  }

  #mine{
    background-color: #93df83;
    float: right;
  }

  #comment{
    background-color: #FFFFFF;
    float: left;
  }
</style>