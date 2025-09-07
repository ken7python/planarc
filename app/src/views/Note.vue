<script setup lang="ts">
  import commentIcon from '@/assets/icons/comment.svg';
  import MicIcon from '@/assets/icons/mic.svg';
  import { ref } from 'vue';
  import { mic } from '@/logic/mic';

  import { CommentModule } from "../logic/comment";

  const message = ref<string>('');

  async function micbtn() {
    // alert('マイクボタンが押されました。');
    if (!mic.shouldRestart.value) {
      console.log("start");
      mic.result = message.value;
      mic.micON.value = true;
      await mic.start();
    } else {
      console.log("stop");
      mic.micON.value = false;
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
  CommentModule.refComment.value = "";

  async function ask() {
    disabled.value = true;
    const comment = await CommentModule.ask(props.date,message.value);
    console.log(comment);
    // ここでcommentを表示する処理を追加
  }
</script>

<template>
  <div id="note">
    <div class="micdiv">
      <textarea placeholder="メモ" v-model="message"></textarea>
      <mic-icon class="mic" :style="mic.micStyle()" @click="micbtn"></mic-icon>
    </div>
    <br>
    <button class="btn" style="margin: 0 auto;" @click="ask" :disabled="disabled">
      <commentIcon style="margin-right: 8px;"></commentIcon>
      <span v-if="!disabled">AIからのコメント</span>
      <span v-else>通信中...</span>
    </button>
    <br>
    <div id="comment">
      <div v-if="CommentModule.refComment.value">
        <p>{{ CommentModule.refComment }}</p>
      </div>
      <div v-else>
        <p>コメントはまだありません。</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
  #note {
    text-align: center;
    overflow-y: auto;
    height: calc(100vh - 60px - 60px - 60px);
  }

  textarea {
    width: 70%;
    height: 22vh;
    background-color: #FFFFFF;
    padding: 10px;
    font-size: 16px;
    margin: 4px 0;
    border: 1px solid #E2E8F8;
    border-radius: 4px;
  }
</style>