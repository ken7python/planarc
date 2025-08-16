<script setup lang="ts">
  import commentIcon from '@/assets/icons/comment.svg';
  import MicIcon from '@/assets/icons/mic.svg';
  import { ref } from 'vue';
  import { mic } from '@/logic/mic';

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
        mic.result += transcript;
        mic.tmp = '';
        message.value = mic.result;
      } else {
        interim += transcript;
        mic.tmp = interim;
        message.value = mic.result + mic.tmp;
      }
    }
  };
</script>

<template>
  <div id="note">
    <div class="micdiv">
      <textarea placeholder="メモ" v-model="message"></textarea>
      <mic-icon class="mic" :style="mic.micStyle()" @click="micbtn"></mic-icon>
    </div>
    <br>
    <button class="btn" style="margin: 0 auto;">
      <commentIcon style="margin-right: 8px;"></commentIcon>
      AIからのコメント
    </button>
    <br>
    <div id="comment">
      <p>AIからのコメントがここに表示されます。</p>
    </div>
  </div>
</template>

<style scoped>
  #note {
    text-align: center;
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