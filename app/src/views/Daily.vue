<script setup lang="ts">
  import Footer from "../components/Footer.vue";
  import Header from "../components/Header.vue";
  import HeaderDaily from "../components/HeaderDaily.vue";
  import {ref} from "vue";

  const list = "TodoList";
  const studylog = "StudyLog";
  const track = "TimeTrack";
  const letter = "Letter";
  const activeMenu = ref<list | studylog | track | letter>(list);
</script>

<template>
  <Header />
  <HeaderDaily v-model:menu="activeMenu"/>
  <div id="main">
    {{ activeMenu }}
  </div>
  <Footer />
</template>

<style scoped>

</style>