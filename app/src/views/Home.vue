<script setup lang="ts">
  import { ref, onMounted, watch } from 'vue';
  import { user } from '@/logic/user.js';

  import Footer from "@/components/Footer.vue";
  import Header from "@/components/Header.vue";

  import { selectStyle } from '@/logic/style/selectStyle';
  import { subjectModule } from '@/logic/subject';

  import Pie from "@/components/Pie.vue";

  import mode1 from '@/assets/icons/1.svg';
  import mode2 from '@/assets/icons/2.svg';
  import mode3 from '@/assets/icons/3.svg';
  import mode4 from '@/assets/icons/4.svg';

  import { mic } from '@/logic/mic';
  import MicIcon from '@/assets/icons/mic.svg';

  import { Chart } from "chart.js/auto";

  let subjectName = ref<string>('');

  let enjoyment = ref<string>('');

  let subjects = ref<any[]>([]);

  let selectedMood = ref<number>(null);
  function moodSelect(mood: number) {
    selectedMood.value = mood;
  }
  const getMoodStyle = (mood: number) => {
    return {
      color: selectedMood.value === mood ? 'black' : 'gray',
      background: selectedMood.value === mood ? 'yellow' : 'white',
      cursor: 'pointer'
    };
  };

  const numberOfToDO = ref<number>(10)
  const finishedToDo = ref<number>(5);

  const username = ref<string | null>(null);
  async function loadData() {
    const profile = await user.profile();
    // console.log(profile);
    username.value = profile.username;

    const subject_list = await subjectModule.getList();
    console.log(subject_list);
    subjects.value = subject_list;
  }
  loadData();

  async function micbtn() {
    // alert('マイクボタンが押されました。');
    if (!mic.shouldRestart.value) {
      console.log("start");
      mic.result = '';
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
        enjoyment.value = mic.result;
      } else {
        interim += transcript;
        mic.tmp = interim;
        enjoyment.value = mic.result + mic.tmp;
      }
    }
  };
</script>

<template>
  <div id="tmp">
    <header>
      <Header />
<!--      <HeaderHome v-model:menu="activeMenu" />-->
    </header>
    <div id="main">
      <h3>今日の楽しみ</h3>
      <div class="micdiv">
        <input type="text" placeholder="今日の楽しみを入力してください" v-model="enjoyment" />
        <mic-icon :style="mic.micStyle()" @click="micbtn"></mic-icon>
      </div>

      <hr>

      <div id="feeling">
        <h3 style="line-height: 0;margin-bottom: 0">今日の気分</h3>
        <div id="feeling-icons">
          <mode4 @click="moodSelect(4)" :style="getMoodStyle(4)"></mode4>
          <mode3 @click="moodSelect(3)" :style="getMoodStyle(3)"></mode3>
          <mode2 @click="moodSelect(2)" :style="getMoodStyle(2)"></mode2>
          <mode1 @click="moodSelect(1)" :style="getMoodStyle(1)"></mode1>
        </div>
      </div>

      <hr>

      <div id="loginbonus">
        <h3>〇日連続ログイン</h3>
        <h3>累計〇日使用</h3>
      </div>

      <hr>
      

      <div id="studyplan">
        <h3>学習予定一覧</h3>
        <ul>
          <li>科目1</li>
          <li>科目2</li>
          <li>科目3</li>
        </ul>
      </div>

      <hr>

      <div id="progress">
        <h3>進捗バー</h3>
        <div v-if="numberOfToDO && finishedToDo">
          <p style="text-align: right;line-height: 0.5;">{{ Math.round(finishedToDo * 100 / numberOfToDO) }}%</p>
          <progress class="my-progress" :value="finishedToDo" :max="numberOfToDO"></progress>
        </div>
        <div v-else>
          <p>通信中...</p>
        </div>
      </div>

      <hr>

      <div id="studyTime">
        <h3>学習時間　<span class="underlined">〇〇時間</span></h3>
        <Pie></Pie>

        <hr>

        <h3>学習時間（科目別）　<span class="underlined">〇〇時間</span></h3>

        <select class="selectbox" :style="selectStyle.getSelectStyle(subjectName)" v-model="subjectName">
          <option value="">科目を選択</option>
          <option v-for="subject in subjects" :key="subject.value" :value="subject.ID">
            {{ subject.Name }}
          </option>
        </select>
      </div>

    </div>
    <footer>
      <Footer />
    </footer>
  </div>
</template>

<style scoped>
h3 {
  line-height: 0.75;
}
.underlined {
  text-decoration: underline;
}
progress {
  width: 100%;
  height: 20px;
}

 #main {
   overflow-y: scroll;
 }
 #tmp {
   height: calc(100dvh - 80px);
   display: grid;
   grid-template-rows: auto 1fr auto; /* ヘッダー=auto / 中央=残り / フッター=auto */
 }

 #feeling-icons {
    display: flex;
    justify-content: space-around;
    align-items: center;
    height: 60px;
 }
 svg {
    width: 30px;
    height: 30px;
    cursor: pointer;
 }
</style>
