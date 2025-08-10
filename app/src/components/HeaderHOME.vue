<script setup lang="ts">
  import { ref, watch } from 'vue';
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

  function underlineStyle(trueMenu: string) {
    return {
      width: '100%',
      height: '8px',
      backgroundColor: menu.value === trueMenu ? '#66BEBC' : '',
    };
  }

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
  <div id="header_home">
    <!-- 勉強時間 -->
    <div class="menu">
      <a :style="getLinkStyle(log)" @click="menu = log;">
        <noteIcon :style='getIconStyle(log)'></noteIcon>
        勉強時間
      </a>
      <div :style="underlineStyle(log)"></div>
    </div>
    <!-- 科目作成 -->
    <div class="menu">
      <a :style="getLinkStyle(subject)" @click="menu = subject">
        <subjectIcon :style='getIconStyle(subject)'></subjectIcon>
        科目作成
      </a>
      <div :style="underlineStyle(subject)"></div>
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
</style>