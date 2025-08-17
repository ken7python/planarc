<script setup lang="ts">
  import { ref } from 'vue';
  import { selectStyle } from '@/logic/style/selectStyle';
  import { subjectModule } from '@/logic/subject';
  import { stopwatch } from '@/logic/StudyLog';

  import startIcon from '@/assets/icons/start.svg';
  import stopIcon from '@/assets/icons/stop.svg';
  import writeIcon from '@/assets/icons/write.svg';
  import {studyLog} from "@/logic/StudyLog";

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
        <span v-if="stopwatch.sHours.value != null && stopwatch.sHours.value<10">0{{ stopwatch.sHours.value }}</span>
        <span v-else>{{ stopwatch.sHours.value }}</span>
        <span v-if="stopwatch.sMinutes.value != null">:</span>
        <span v-if="stopwatch.sMinutes.value != null && stopwatch.sMinutes.value<10">0{{ stopwatch.sMinutes }}</span>
        <span v-else>{{ stopwatch.sMinutes }}</span>
      </span>
    </div>
    <div id="end">
      <button class="btn" @click="stopwatch.stop()"><stopIcon></stopIcon>終了</button>
      <span class="time">
        <span v-if="stopwatch.eHours.value != null && stopwatch.eHours.value<10">0{{ stopwatch.eHours.value }}</span>
        <span v-else>{{ stopwatch.eHours.value }}</span>
        <span v-if="stopwatch.eMinutes.value != null">:</span>
        <span v-if="stopwatch.eMinutes.value != null && stopwatch.eMinutes.value<10">0{{ stopwatch.eMinutes }}</span>
        <span v-else>{{ stopwatch.eMinutes }}</span>
      </span>
    </div>
    <div id="res">
      <span class="time">
        <span v-if="stopwatch.dHours.value != null">{{ stopwatch.dHours }}時間</span>
        <span v-if="stopwatch.dMinutes.value != null">{{ stopwatch.dMinutes.value }}分</span>
      </span>
      <button class="btn"  style="margin: 0 auto;" @click="studyLog.write(stopwatch.sHours.value,stopwatch.sMinutes.value,stopwatch.eHours.value,stopwatch.eMinutes.value)"><writeIcon></writeIcon>記録</button>
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