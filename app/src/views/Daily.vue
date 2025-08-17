<script setup lang="ts">
  import Footer from "../components/Footer.vue";
  import Header from "../components/Header.vue";
  import HeaderDaily from "../components/HeaderDaily.vue";
  import {ref} from "vue";

  import { header } from "@/logic/header.js";
  import { user } from "@/logic/user.js";

  import TodoList from "@/views/TodoList.vue";
  import StudyLog from "@/views/StudyLog.vue";
  import Track from "@/views/TimeTrack.vue";
  import Letter from "@/views/Note.vue";

  const list = header.list;
  const studylog = header.studylog;
  const track = header.track;
  const letter = header.letter;
  const activeMenu = ref<string>(list);

  const today = new Date();
  const yyyy = today.getFullYear();
  const mm = String(today.getMonth() + 1).padStart(2, '0');
  const dd = String(today.getDate()).padStart(2, '0');
  const date = ref<string>(`${yyyy}-${mm}-${dd}`); // 日本時間の今日の日付

  const username = ref<string | null>(null);
  async function loadData() {
    const profile = await user.profile();
    // console.log(profile);
    username.value = profile.username;
  }
  loadData();
</script>

<template>
  <Header />
  <div id="calendar">
    <input type="date" v-model="date">
  </div>
  <HeaderDaily v-model:menu="activeMenu"/>

  <div id="main">
    <TodoList v-if="activeMenu === list" />
    <StudyLog v-if="activeMenu === studylog" :date="date" />
    <Track v-if="activeMenu === track" />
    <Letter v-if="activeMenu === letter" />
  </div>
  <Footer />
</template>

<style scoped>
  #main {
    height: calc(100dvh - 80px - 60px - 60px - 40px);
    display: grid;
    grid-template-rows: auto auto auto 1fr auto; /* ヘッダー=auto / 中央=残り / フッター=auto */
  }

  #calendar {
    width: 100%;
    height: 40px;
    background-color: #1C409A;
    display: flex;
    justify-content: center;
    align-items: center;
    color: #FFFFFF;
    font-size: 24px;
    font-weight: bold;
  }
</style>