<script setup lang="ts">
  import { ref, onMounted, watch } from 'vue';
  import { user } from '@/logic/user.js';
  import { subjectModule } from '@/logic/subject';
  import { todoModule } from "@/logic/todo";
  import { studyLog } from "@/logic/StudyLog";
  import { CONST } from "@/logic/const";

  import Footer from "@/components/Footer.vue";
  import Header from "@/components/Header.vue";

  import { selectStyle } from '@/logic/style/selectStyle';


  import Pie from "@/components/Pie.vue";



  import { mic } from '@/logic/mic';

  let subjectName = ref<string>('');

  let subjects = ref<any[]>([]);

  const numberOfToDO = ref<number>(null)
  const finishedToDo = ref<number>(null);

  const sumToday = ref<number>(0);

  const sumTodayofSubject = ref<number>(0);
  async function getTodayofSubject(subjectId: number) {
    const studyLogs = await studyLog.getLog(today);
    console.log(studyLogs);

    studyLogs.map( log => {
      if (log.SubjectID === subjectId) {
        sumTodayofSubject.value += log.StudyTime;
      }
    });
  }

  watch (subjectName, (newVal) => {
    if (newVal) {
      sumTodayofSubject.value = 0;
      getTodayofSubject(Number(newVal));
    }
  });

  const today = CONST.getToday();
  console.log(today);

  const uSubjectNames = ref<string[]>([]);

  const username = ref<string | null>(null);
  async function loadData() {
    const profile = await user.profile();
    // console.log(profile);
    username.value = profile.username;

    const subject_list = await subjectModule.getList();
    console.log(subject_list);
    subjects.value = subject_list;

    const studyLogs = await studyLog.getLog(today);
    console.log(studyLogs);

    studyLogs.map(log => {
      sumToday.value += log.StudyTime;
    })

    const TODOList = await todoModule.getList(today);
    console.log(TODOList);

    const finished = TODOList.filter(task => task.Checked)
    console.log(finished);

    numberOfToDO.value = TODOList.length || 0;
    finishedToDo.value = finished.length || 0;

    const uSubjectIDs = [...new Set(TODOList.map(task => {
      if (!task.Checked) {
        return task.SubjectID;
      }
    }))];
    console.log(uSubjectIDs);
    uSubjectNames.value = subject_list.filter(subject => uSubjectIDs.includes(subject.ID)).map(subject => {
      return subject.Name;
    });
    console.log(uSubjectNames.value);
  }
  loadData();
</script>

<template>
  <div id="tmp">
    <header>
      <Header />
<!--      <HeaderHome v-model:menu="activeMenu" />-->
    </header>
    <div id="main">
      <div id="studyplan">
        <h3>学習予定一覧</h3>
        <ul v-if="uSubjectNames.length != 0" v-for="subject in uSubjectNames" :key="subject">
          <li>{{ subject }}</li>
        </ul>
        <div v-else>
          <p>今日のタスクはありません</p>
        </div>
      </div>

      <hr>

      <div id="progress">
        <h3>進捗バー</h3>
        <div v-if="numberOfToDO != null && finishedToDo != null">
          <p style="text-align: right;line-height: 0.5;">{{ ( Math.round(finishedToDo * 100 / numberOfToDO) ) || 0 }}%</p>
          <progress class="my-progress" :value="finishedToDo" :max="numberOfToDO"></progress>
        </div>
        <div v-else>
          <p>通信中...</p>
        </div>
      </div>

      <hr>

      <div id="studyTime">
        <h3>学習時間　<span class="underlined">{{ Math.floor(sumToday / 60) }}時間{{ sumToday % 60 }}分</span></h3>
        <Pie></Pie>

        <hr>

        <h3>学習時間（科目別）　<span class="underlined">{{ Math.floor(sumTodayofSubject / 60) }}時間{{ sumTodayofSubject % 60 }}分</span></h3>

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

 svg {
    width: 30px;
    height: 30px;
    cursor: pointer;
 }

#main {
  height: calc(100dvh - 80px);
  overflow-y: scroll;
}
</style>
