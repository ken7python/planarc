<script setup lang="ts">
  import { ref } from 'vue';
  import { user } from '@/logic/user.js';
  import Footer from "@/components/Footer.vue";
  import Header from "@/components/Header.vue";
  import HeaderHome from "../components/HeaderHOME.vue";

  let data: any;
  let username = ref('');

  const log = 'StudyLog';
  const subject = 'AddSubject';
  const activeMenu = ref< log | subject >(log);


  async function loadData() {
    data = await user.profile();
    // console.log(data);
    username.value = data.username;
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
<!--        <p>ようこそ、{{ username }}さん！</p>-->
<!--      </div>-->
<!--      <div v-else>-->
<!--        <p>ユーザー情報を取得中...</p>-->
<!--      </div>-->
    {{ activeMenu }}
    </div>
    <footer>
      <Footer />
    </footer>
  </div>
</template>

<style scoped>

</style>
