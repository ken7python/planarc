<script setup lang="ts">
  import { ref } from 'vue';
  import { selectStyle } from '@/logic/style/selectStyle';
  import startIcon from '@/assets/icons/start.svg';
  import stopIcon from '@/assets/icons/stop.svg';
  import writeIcon from '@/assets/icons/write.svg';
  let subjectName = ref<string>('');

  const subjects = [
    { name: "国語", value :"kokugo" },
    { name: "数学", value :"sugaku" },
    { name: "英語", value :"eigo" },
    { name: "理科", value :"rika" },
    { name: "社会", value :"shakai" }
  ]
</script>

<template>
  <div id="stopwatch">
    <select class="selectbox" :style="selectStyle.getSelectStyle(subjectName)" v-model="subjectName">
      <option value="">科目を選択</option>
      <option v-for="subject in subjects" :key="subject.value" :value="subject.value">
        {{ subject.name }}
      </option>
    </select>
    <div id="start">
      <button class="btn"><startIcon></startIcon>開始</button>
      <span class="time">--:--</span>
    </div>
    <div id="end">
      <button class="btn"><stopIcon></stopIcon>終了</button>
      <span class="time">--:--</span>
    </div>
    <div id="res">
      <span class="time">--:--</span>
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
    font-size: 64px;
    margin-left: 10px;
  }
</style>