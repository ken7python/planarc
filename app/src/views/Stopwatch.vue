<script setup lang="ts">
  import { ref } from 'vue';
  import { selectStyle } from '@/logic/style/selectStyle';
  import { subjectModule } from '@/logic/subject';

  import startIcon from '@/assets/icons/start.svg';
  import stopIcon from '@/assets/icons/stop.svg';
  import writeIcon from '@/assets/icons/write.svg';
  let subjectName = ref<string>('');

  let subjects = ref<any[]>([]);
  async function loadData() {
    const subject_list = await subjectModule.getList();
    console.log(subject_list);
    subjects.value = subject_list;
  }

  loadData();
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