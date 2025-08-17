<script setup lang="ts">
  import { ref } from 'vue';
  import { selectStyle } from '@/logic/style/selectStyle';
  import { subjectModule } from '@/logic/subject';
  import { stopwatch } from '@/logic/StudyLog';

  import startIcon from '@/assets/icons/start.svg';
  import stopIcon from '@/assets/icons/stop.svg';
  import writeIcon from '@/assets/icons/write.svg';

  let subjectName = ref<string>('');

  let subjects = ref<any[]>([]);
  async function loadData() {
    const subject_list = await subjectModule.getList();
    console.log(subject_list);
    subjects.value = subject_list;

    stopwatch.init()
  }

  loadData();

  setInterval(function() {
    stopwatch.update();
  }, 1000);
</script>

<template>
  <div id="stopwatch">
    <select class="selectbox" :style="selectStyle.getSelectStyle(subjectName)" v-model="subjectName">
      <option value="">科目を選択</option>
      <option v-for="subject in subjects" :key="subject.value" :value="subject.ID">
        {{ subject.Name }}
      </option>
    </select>

    <div id="start">
      <button class="btn" @click="stopwatch.start()"><startIcon></startIcon>開始</button>
      <span class="time">
        <span v-if="stopwatch.sHours<10">0</span>{{ stopwatch.sHours }}:<span v-if="stopwatch.sMinutes<10">0</span>{{ stopwatch.sMinutes }}
      </span>
    </div>
    <div id="end">
      <button class="btn" @click="stopwatch.stop()"><stopIcon></stopIcon>終了</button>
      <span class="time">
        <span v-if="stopwatch.eHours<10">0</span>{{ stopwatch.eHours }}:<span v-if="stopwatch.eMinutes<10">0</span>{{ stopwatch.eMinutes }}
      </span>
    </div>
    <div id="res">
      <span class="time">
        <span v-if="stopwatch.dHours.value != null">{{ stopwatch.dHours }}時間</span>
        <span v-if="stopwatch.dMinutes.value != null">{{ stopwatch.dMinutes.value }}分</span>
      </span>
      <button class="btn"  style="margin: 0 auto;"><writeIcon></writeIcon>記録</button>
    </div>
  </div>
</template>

<style scoped>
  #stopwatch {
    text-align: center;
  }
  #start, #end{
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 10px 0;
  }

  .btn{
    width: 200px;
  }
  .time {
    font-size: 42px;
    margin-left: 10px;
  }
</style>