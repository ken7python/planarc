<script setup lang="ts">
  import { ref } from 'vue';
  import { profile,logout } from '@/assets/user.js';
  import Footer from "@/components/Footer.vue";
  import Header from "@/components/Header.vue";

  let data: any;
  let username = ref('');
  async function loadData() {
    data = await profile();
    // console.log(data);
    username.value = data.username;
  }
  loadData();

  // const onLongpress = (e: CustomEvent) => {
  //   alert("長押しイベントが発生しました: ");
  // }
</script>

<template>
  <div>
    <header>
      <Header />
    </header>
    <div id="main">
      <h1>Welcome to PlanArc</h1>
      <div v-if="username">
        <p>ようこそ、{{ username }}さん！</p>
      </div>
      <div v-else>
        <p>ユーザー情報を取得中...</p>
      </div>
      <button class="btn" @click="logout">Logout</button>
    </div>
    <footer>
      <Footer />
    </footer>
  </div>
</template>

<style scoped>

</style>
