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
  <HeaderDaily v-model:menu="activeMenu"/>

  <div id="main">
    <TodoList v-if="activeMenu === list" />
    <StudyLog v-if="activeMenu === studylog" />
    <Track v-if="activeMenu === track" />
    <Letter v-if="activeMenu === letter" />
  </div>
  <Footer />
</template>

<style scoped>
  #main {
    /* min-height: calc(100dvh - 80px - 60px - 60px); */
    height: calc(100dvh - 80px - 60px - 60px);;
    display: grid;
    grid-template-rows: auto 1fr auto; /* ヘッダー=auto / 中央=残り / フッター=auto */
  }
</style>