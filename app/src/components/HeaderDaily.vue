<script setup lang="ts">
  import { ref } from 'vue';
  const list = "TodoList";
  const studylog = "StudyLog";
  const track = "TimeTrack";
  const letter = "Letter";

  import listIcon from '@/assets/icons/list.svg';
  import studylogIcon from '@/assets/icons/timer.svg';
  import trackIcon from '@/assets/icons/barefoot.svg';
  import letterIcon from '@/assets/icons/ink_pen.svg';

  let menu = ref(list);

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
    }
  }

  function getLinkStyle(trueMenu: string){
    const weight = menu.value === trueMenu ? 'bold' : 'normal';
    return {
      color: getColor(trueMenu),
      fontWeight: weight,
      fontSize: '16px',
    }
  }
</script>

<template>
  <div id="header_home">
    <!-- TODOLIST -->
    <div class="menu">
      <div :style="getLinkStyle(list)" @click="menu = list;" class="menuGrid">
        <div class="iconDiv">
          <listIcon />
        </div>
        <span class="menuText">
          TODO
          <br>
          LIST
        </span>
      </div>
      <div v-if="menu === list" class="underline"></div>
      <div v-else style="height: 4px;"></div>
    </div>

    <!-- 勉強記録 -->
    <div class="menu">
      <a :style="getLinkStyle(studylog)" @click="menu = studylog" class="menuGrid">
        <div class="iconDiv">
          <studylogIcon />
        </div>
        <span class="menuText">
          勉強記録
        </span>
      </a>
      <div v-if="menu === studylog" class="underline"></div>
      <div v-else style="height: 4px;"></div>
    </div>
    <!-- TIME TRACK -->
    <div class="menu">
      <a :style="getLinkStyle(track)" @click="menu = track" class="menuGrid">
        <div class="iconDiv">
          <trackIcon />
        </div>
        <span class="menuText">
          TIME<br>TRACK
        </span>
      </a>
      <div v-if="menu === track" class="underline"></div>
      <div v-else style="height: 4px;"></div>
    </div>
    <!-- 振り返り -->
    <div class="menu">
      <a :style="getLinkStyle(letter)" @click="menu = letter" class="menuGrid">
        <div class="iconDiv">
          <letterIcon />
        </div>
        <span class="menuText">
          振り返り
        </span>
      </a>
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
    display: flex;
    flex-direction: row;
    justify-content: space-around;
    align-items: center;
    background-color: #1C409A;
    text-align: center;
  }
  .menu {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: end;
    flex: 1;
    height: 100%;
  }
  .underline {
    width: 100%;
    height: 4px;
    background-color: #66BEBC;
    margin-top: 8px;
  }
  .menuGrid {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    gap: 6px;
    width: 100%

  }
</style>