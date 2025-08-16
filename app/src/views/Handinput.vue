<script setup lang="ts">
  import { ref } from 'vue';
  import { selectStyle } from '@/logic/style/selectStyle';
  import { subjectModule } from '@/logic/subject';

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
  <div id="headerinput">
    <select class="selectbox" :style="selectStyle.getSelectStyle(subjectName)" v-model="subjectName">
      <option value="">科目を選択</option>
      <option v-for="subject in subjects" :key="subject.value" :value="subject.ID">
        {{ subject.Name }}
      </option>
    </select>
    <div id="inputbox">
      <div>
        <h2>開始時間</h2>
        <input type="time" id="startTime" name="startTime" value="--:--">
      </div>
      <div>
        <h2>終了時間</h2>
        <input type="time" id="startTime" name="startTime" value="--:--">
      </div>
    </div>
    <div id="res">
      <button class="btn"  style="margin: 0 auto;"><writeIcon></writeIcon>記録</button>
    </div>
  </div>
</template>

<style scoped>
  #headerinput {
    text-align: center;
  }

  #inputbox {
    display: flex;
    justify-content: space-around;
    align-items: center;
    margin: 10px 0;
    margin-bottom: 80px;
  }

  .btn{
    width: 200px;
  }
</style>