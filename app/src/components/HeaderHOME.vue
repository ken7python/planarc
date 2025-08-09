<script setup lang="ts">
  import { ref } from 'vue';
  const log = "StudyLog";
  const subject = "AddSubject";

  import noteIcon from '@/assets/icons/note_alt.svg';
  import subjectIcon from '@/assets/icons/subject.svg';

  let menu = ref(log);

  function getColor(trueMenu: string){
    if (menu.value === trueMenu) {
      return '#FFFFFF';
    } else {
      return '#739DDB';
    }
  }

  function getIconStyle(trueMenu: string){
    return {
      color: getColor(trueMenu),
      width: '30px',
      height: '30px',
      display: 'inline-block',
      position: 'relative',
      top: '5px',
    }
  }

  function getLinkStyle(trueMenu: string){
    const weight = menu.value === trueMenu ? 'bold' : 'normal';
    return {
      color: getColor(trueMenu),
      fontWeight: weight,
      fontSize: '24px',
    }
  }

  function toStudyLog() {
    menu.value = log;
  }
  function toAddSubject() {
    menu.value = subject;
  }
</script>

<template>
  <div id="header_home">
    <!-- 勉強時間 -->
    <div class="menu">
      <a :style="getLinkStyle(log)" @click="toStudyLog()">
        <noteIcon :style='getIconStyle(log)'></noteIcon>
        勉強時間
      </a>
      <div v-if="menu === log" class="underline"></div>
      <div v-else style="height: 4px;"></div>
    </div>
    <!-- 科目作成 -->
    <div class="menu">
      <a :style="getLinkStyle(subject)" @click="toAddSubject()">
        <subjectIcon :style='getIconStyle(subject)'></subjectIcon>
        科目作成
      </a>
      <div v-if="menu === subject" class="underline"></div>
      <div v-else style="height: 4px;"></div>
    </div>
  </div>
</template>

<style scoped>
  #header_home {
    position: absolute;
    top: 45px;
    left: 0;
    width: 100%;
    height: 60px;
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    background-color: #1C409A;
    text-align: center;
  }
  .menu {
    display: flex;
    justify-content: end;
    align-items: center;
    flex-direction: column;
  }
  .underline {
    width: 100%;
    height: 4px;
    background-color: #66BEBC;
  }
</style>