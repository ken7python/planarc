<script setup lang="ts">
  import { ref } from 'vue';
  const list = "TodoList";
  const studylog = "StudyLog";
  const track = "TimeTrack";
  const letter = "Letter";

  let menu = ref(list);

  function getColor(trueMenu: string){
    if (menu.value === trueMenu) {
      return '#FFFFFF';
    } else {
      return '#739DDB';
    }
  }

  function getIconStyle(trueMenu: string,filename: string){
    return {
      color: getColor(trueMenu),
      width: '40px',
      height: '40px',
      display: 'inline-block',
    }
  }

  function getLinkStyle(trueMenu: string){
    const weight = menu.value === trueMenu ? 'bold' : 'normal';
    return {
      color: getColor(trueMenu),
      fontWeight: weight,
      fontSize: '14px',
    }
  }
</script>

<template>
  <div id="header_home">
    <!-- TODOLIST -->
    <div class="menu">
      <a :style="getLinkStyle(list)" @click="menu = list;">TODO<br>LIST</a>
      <div v-if="menu === list" class="underline"></div>
      <div v-else style="height: 4px;"></div>
    </div>
    <!-- 勉強記録 -->
    <div class="menu">
      <a :style="getLinkStyle(studylog)" @click="menu = studylog">勉強記録</a>
      <div v-if="menu === studylog" class="underline"></div>
      <div v-else style="height: 4px;"></div>
    </div>
    <!-- TIME TRACK -->
    <div class="menu">
      <a :style="getLinkStyle(track)" @click="menu = track">TIME<br>TRACK</a>
      <div v-if="menu === track" class="underline"></div>
      <div v-else style="height: 4px;"></div>
    </div>
    <!-- 振り返り -->
    <div class="menu">
      <a :style="getLinkStyle(letter)" @click="menu = letter">振り返り</a>
      <div v-if="menu === letter" class="underline"></div>
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
    grid-template-columns: repeat(4, 1fr);
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