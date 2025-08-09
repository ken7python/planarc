<script setup lang="ts">
  import { useRoute } from 'vue-router'
  import homeIcon from '@/assets/icons/home.svg';
  import dailyIcon from '@/assets/icons/calendar_today.svg';
  import flagIcon from '@/assets/icons/flag.svg';

  import { CONST } from '@/logic/const.js';

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
      width: '40px',
      height: '40px',
      display: 'inline-block',
    }
  }

  function getLinkStyle(truePath: string){
    const weight = route.path === truePath ? 'bold' : 'normal';
    return {
      color: getColor(truePath),
      fontWeight: weight,
    }
  }
</script>

<template>
  <div id="header">

  </div>
</template>

<style scoped>
  #header {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 80px;
    background-color: #1C409A;
  }
</style>