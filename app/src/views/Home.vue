<script setup lang="ts">
  import { ref } from 'vue';
  import { user } from '@/logic/user.js';

  import Footer from "@/components/Footer.vue";
  import Header from "@/components/Header.vue";
  import HeaderHome from "@/components/HeaderHOME.vue";
  import { header } from "@/logic/header.js";

  import StudyTime from "./StudyTime.vue";
  import AddSubject from "./AddSubject.vue";

  const log = header.log;
  const subject = header.subject;
  const activeMenu = ref< log | subject >(log);


  const username = ref<string | null>(null);
  async function loadData() {
    const profile = await user.profile();
    // console.log(profile);
    username.value = profile.username;
  }
  loadData();
</script>

<template>
  <div>
    <header>
      <Header />
      <HeaderHome v-model:menu="activeMenu" />
    </header>
    <div id="main">
<!--    {{ activeMenu }}-->
      <study-time v-if="activeMenu === log" />
      <add-subject v-if="activeMenu === subject" />
    </div>
    <footer>
      <Footer />
    </footer>
  </div>
</template>

<style scoped>
  #main {
    //min-height: calc(100dvh - 80px - 60px);
    height: calc(100dvh - 80px - 60px - 60px); /* ヘッダー80px / フッター60px */
    display: grid;
    grid-template-rows: auto 1fr auto; /* ヘッダー=auto / 中央=残り / フッター=auto */
  }
</style>
