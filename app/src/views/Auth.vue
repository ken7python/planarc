<script setup lang="ts">
  import { CONST } from '@/logic/const.js';
  import Header from "@/components/Header.vue";
</script>

<template>
  <div>
    <Header />
    <div id="main">
      <router-link :to="CONST.signupLink" class="btn">
        <p>ユーザ登録</p>
      </router-link>

      <router-link :to="CONST.loginLink" class="btn">
        <p>ログイン</p>
      </router-link>
    </div>
  </div>
</template>

<style scoped>
  #main {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 80vh;
  }
  .btn {
    margin-top: 5rem;
    margin-bottom: 5rem;
  }
</style>