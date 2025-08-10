<script setup lang="ts">
  import Footer from "../components/Footer.vue";
  import Header from "../components/Header.vue";
  import HeaderDaily from "../components/HeaderDaily.vue";
  import {ref} from "vue";

  import { header } from "@/logic/header.js";

  import TodoList from "@/views/TodoList.vue";
  import StudyLog from "@/views/StudyLog.vue";
  import Track from "@/views/TimeTrack.vue";
  import Letter from "@/views/Note.vue";

  const list = header.list;
  const studylog = header.studylog;
  const track = header.track;
  const letter = header.letter;
  const activeMenu = ref<string>(list);
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

</style>