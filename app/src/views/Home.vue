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
<!--      <h1>Welcome to PlanArc</h1>-->
<!--      <div v-if="username">-->
      <div id="welcome" style="">
        <div v-if="username">
          <p>ようこそ、{{ username }}さん！</p>
        </div>
        <div v-else>
          <p>ユーザー情報を取得中...</p>
        </div>
        <hr>
      </div>
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
  #welcome {
    /* 行間をなくす */
    line-height: 2px;
  }
</style>
