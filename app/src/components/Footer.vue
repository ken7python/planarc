<script setup lang="ts">
  import { useRoute } from 'vue-router'
  import homeIcon from '@/assets/icons/home.svg';
  import subjectIcon from '@/assets/icons/subject.svg';
  import dailyIcon from '@/assets/icons/calendar_today.svg';
  import flagIcon from '@/assets/icons/flag.svg';
  import analysisIcon from '@/assets/icons/analytics.svg';

  import { CONST } from '@/logic/const.js';
  import Analysis from "../views/Analysis.vue";

  const route = useRoute()

  function getColor(truePath: string){
    if (route.path === truePath) {
      return '#FFFFFF';
    } else {
      return '#739DDB';
    }
  }

  function getIconStyle(truePath: string,filename: string){
    return {
      color: getColor(truePath),
      width: '30px',
      height: '30px',
      display: 'inline-block',
    }
  }

  function getLinkStyle(truePath: string){
    const weight = route.path === truePath ? 'bold' : 'normal';
    return {
      color: getColor(truePath),
      fontWeight: weight,
      fontSize: '12px',
    }
  }

  // console.log(CONST);
</script>

<template>
  <div>
    <div id="footer">
      <!-- HOME -->
      <router-link class="footer-item" :to="CONST.homeLink" >
        <homeIcon :style='getIconStyle(CONST.homeLink)'></homeIcon>
        <p :style="getLinkStyle(CONST.homeLink)">ホーム</p>
      </router-link>
      <!-- FLAG -->
      <router-link class="footer-item" :to="CONST.subjectLink">
        <subjectIcon :style='getIconStyle(CONST.subjectLink)'></subjectIcon>
        <p :style="getLinkStyle(CONST.subjectLink)">科目作成</p>
      </router-link>
      <!-- DAILY -->
      <router-link class="footer-item" :to="CONST.dailyLink">
        <dailyIcon :style='getIconStyle(CONST.dailyLink)'></dailyIcon>
        <p :style="getLinkStyle(CONST.dailyLink)">日</p>
      </router-link>

      <router-link class="footer-item" :to="CONST.flagLink">
        <flagIcon :style='getIconStyle(CONST.flagLink)'></flagIcon>
        <p :style="getLinkStyle(CONST.flagLink)">未完了リスト</p>
      </router-link>

      <router-link class="footer-item" :to="CONST.analysisLink">
        <analysisIcon :style='getIconStyle(CONST.analysisLink)'></analysisIcon>
        <p :style="getLinkStyle(CONST.analysisLink)">分析</p>
      </router-link>

    </div>
  </div>
</template>

<style scoped>
  #footer{
    position: fixed;
    bottom: 0;
    left: 0;
    width: 100%;
    height: 50px;
    display: grid;
    grid-template-columns: repeat(5, 1fr);
    background-color: #265DC5;
    text-align: center;
  }
  #footer p {
    margin: 0;
    font-size: 12px;
    color: #FFFFFF;
  }
  .footer-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
    text-decoration: none;
  }
  svg{
    margin: 0;
  }
</style>