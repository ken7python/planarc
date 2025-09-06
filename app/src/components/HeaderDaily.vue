<script setup lang="ts">
  import {ref, watch} from 'vue';

  import listIcon from '@/assets/icons/list.svg';
  import studylogIcon from '@/assets/icons/timer.svg';
  import trackIcon from '@/assets/icons/barefoot.svg';
  import letterIcon from '@/assets/icons/ink_pen.svg';
  import { header } from '@/logic/header.js';

  // const list = "TodoList";
  // const studylog = "StudyLog";
  // const track = "TimeTrack";
  // const letter = "Letter";
  const list = header.list;
  const studylog = header.studylog;
  const track = header.track;
  const letter = header.letter;

  let menu = ref(list);
  header.fontsize = '14px';

  const props = defineProps<{
    menu: string;
  }>();
  const emit = defineEmits<{
    (e: 'update:menu', menu: string): void;
  }>();

  watch(menu, (newMenu) => {
    console.log('menu changed:', newMenu);
    emit('update:menu', newMenu);
  });
</script>

<template>
  <div>
    <div id="header_home">
      <!-- ToDoLIST -->
      <div class="menu">
        <div :style="header.getLinkStyle(menu, list)" @click="menu = list;" class="menuGrid">
          <div class="iconDiv">
            <listIcon />
          </div>
          <span class="menuText">
            ToDo
            <br>
            リスト
          </span>
        </div>
        <div :style="header.getUnderlineStyle(menu, list)"></div>
      </div>

      <!-- 学習記録 -->
      <div class="menu">
        <a :style="header.getLinkStyle(menu, studylog)" @click="menu = studylog" class="menuGrid">
          <div class="iconDiv">
            <studylogIcon />
          </div>
          <span class="menuText">
            学習記録
          </span>
        </a>
        <div :style="header.getUnderlineStyle(menu, studylog)"></div>
      </div>
      <!-- TIME TRACK -->
      <div class="menu">
        <a :style="header.getLinkStyle(menu, track)" @click="menu = track" class="menuGrid">
          <div class="iconDiv">
            <trackIcon />
          </div>
          <span class="menuText">
            TIME<br>TRACK
          </span>
        </a>
        <div :style="header.getUnderlineStyle(menu, track)"></div>
      </div>
      <!-- 振り返り -->
      <div class="menu">
        <a :style="header.getLinkStyle(menu, letter)" @click="menu = letter" class="menuGrid">
          <div class="iconDiv">
            <letterIcon />
          </div>
          <span class="menuText">
            振り返り
          </span>
        </a>
        <div :style="header.getUnderlineStyle(menu, letter)"></div>
      </div>
    </div>
  </div>
</template>

<style scoped>
  #header_home {
    /*
    position: absolute;
    top: 75px;
    left: 0;
    */
    width: 100%;
    height: 44px;
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
  .menuGrid {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    gap: 6px;
    width: 100%

  }
</style>