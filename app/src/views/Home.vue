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

  import goalIcon from '@/assets/icons/goal.svg';

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

    const studyLogs = await studyLog.getLog(today);
    console.log(studyLogs);

    studyLogs.map(log => {
      sumToday.value += log.StudyTime;
    })

    const subjectSet = new Set();
    // console.log(studyLogs);
    studyLogs.map(log => {
      // console.log(log.ID);
      console.log(log.SubjectID);
      subjectSet.add(log.SubjectID);
    })

    subject_list.map(subject => {
      if (subjectSet.has(subject.ID)) {
        subjects.value.push(subject);
      }
    })
    console.log(subjectSet);

    const TODOList = await todoModule.getList(today);
    console.log(TODOList);

    const finished = TODOList.filter(task => task.Checked)
    console.log(finished);

    numberOfToDO.value = TODOList.length || 0;
    finishedToDo.value = finished.length || 0;

    const bar = new ProgressBar.Line('#progress-bar', {
      strokeWidth: 3,
      color: '#1C409A',
      trailColor: '#eee',
      trailWidth: 10,
      svgStyle: { width: '80vw', height: '100%' }, // ← 横幅を90%にして左右に余白
    });

// 0% → 100% にアニメーション
    bar.animate((finishedToDo.value / numberOfToDO.value) || 0);  // 1.0 までの値を指定

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

  onMounted(() => {
    loadData();
  });
</script>

<template>
  <div id="tmp">
    <header>
      <Header />
<!--      <HeaderHome v-model:menu="activeMenu" />-->
    </header>
    <div id="main">
      <div id="progress">
        <h3>進捗バー</h3>
        <div v-show="numberOfToDO != null && finishedToDo != null">
          <div style="display: flex; align-items: center; justify-content: center;">
            <span>
              <div style="text-align: center;">
                <b>
                  <span v-if="finishedToDo === 0 || numberOfToDO === 0">0%</span>
                  <span v-else>{{ Math.round(finishedToDo / numberOfToDO * 100) }}%</span>
                </b>
              </div>
              <div id="progress-bar"></div>
            </span>
            <span>
              <goalIcon style="width: 30px;height: 30px;vertical-align: middle;margin-right: 8px;color: #007AFF;"></goalIcon>
            </span>
          </div>
        </div>
        <div v-if="numberOfToDO == null || finishedToDo == null">
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

#progress-bar {
  text-align: center;
}
</style>
