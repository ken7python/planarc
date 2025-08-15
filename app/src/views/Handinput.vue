<script setup lang="ts">
  import { ref } from 'vue';
  import { selectStyle } from '@/logic/style/selectStyle';
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
  <div id="headerinput">
    <select class="selectbox" :style="selectStyle.getSelectStyle(subjectName)" v-model="subjectName">
      <option value="">科目を選択</option>
      <option v-for="subject in subjects" :key="subject.value" :value="subject.value">
        {{ subject.name }}
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