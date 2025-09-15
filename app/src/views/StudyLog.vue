<script setup lang="ts">
  import { ref } from 'vue';
  import Stopwatch from '@/views/Stopwatch.vue';
  import Handinput from '@/views/Handinput.vue';
  import stopwatchIcon from "@/assets/icons/stopwatch.svg";
  import handinputIcon from "@/assets/icons/input.svg";
  const strongGray = "#737373"
  const lightGray = "#F2F2F2"

  const stopwatch = "stopwatch";
  const handinput = "handinput";

  let selectedMenu = ref(stopwatch);

  function select(menu: string) {
    selectedMenu.value = menu;
  }

  function selected(menu: string): boolean {
    return selectedMenu.value === menu;
  }

 function getMenuStyle(menu: string): object {
    return {
      backgroundColor: selected(menu) ? strongGray : lightGray,
      color: selected(menu) ? "#FFFFFF" : "#4d5161",
      cursor: "pointer",
      padding: "10px",
      borderRadius: "5px"

    };
  }

  defineProps({
    date: String
  })
</script>

<template>
  <div>
<!--    {{ date }}-->
    <div id="menu">
      <div id="stopwatchMenu" :style="getMenuStyle(stopwatch)" @click="select(stopwatch)">
        <stopwatchIcon></stopwatchIcon>ストップウォッチ
      </div>
      <div id="handinputMenu" :style="getMenuStyle(handinput)" @click="select(handinput)">
        <handinputIcon></handinputIcon>手動入力
      </div>
    </div>

    <div id="content">
      <Stopwatch v-if="selectedMenu === stopwatch" :date="date"/>
      <Handinput v-if="selectedMenu === handinput" :date="date"/>
    </div>
  </div>
</template>

<style scoped>
  #menu {
    width: 100%;
    height: 40px;
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    background-color: #f0f0f0;
    border-bottom: 1px solid #ccc;
    text-align: center;
  }
  #stopwatchMenu, #handinputMenu {
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 14px;
    font-weight: bold;
  }

  #content {
    padding: 20px;
    text-align: center;
  }
</style>