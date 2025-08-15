<script setup lang="ts">
  import { ref } from 'vue';
  import Stopwatch from '@/views/Stopwatch.vue';
  import Handinput from '@/views/Handinput.vue';
  import _default from "chart.js/dist/plugins/plugin.tooltip";
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
      color: selected(menu) ? "#FFFFFF" : "#000000",
      cursor: "pointer",
      padding: "10px",
      borderRadius: "5px"
    };
  }
</script>

<template>
  <div>
    <div id="menu">
      <div id="stopwatch" :style="getMenuStyle(stopwatch)" @click="select(stopwatch)">
        <a>ストップウォッチ</a>
      </div>
      <div id="handinput" :style="getMenuStyle(handinput)" @click="select(handinput)">
        <a>手動入力</a>
      </div>
    </div>

    <div id="content">
      <Stopwatch v-if="selectedMenu === stopwatch" />
      <Handinput v-if="selectedMenu === handinput" />
    </div>
  </div>
</template>

<style scoped>
  #menu {
    width: 100%;
    height: 60px;
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    background-color: #f0f0f0;
    border-bottom: 1px solid #ccc;
    text-align: center;
  }
  #stopwatch, #handinput{
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 18px;
    color: #333;
  }
</style>