<script setup lang="ts">
  import { ref } from 'vue';
  import { profile,logout } from '@/assets/user.js';
0
  let data: any;
  let username = ref('');
  async function loadData() {
    data = await profile();
    // console.log(data);
    username.value = data.username;
  }
  loadData();
</script>

<template>
  <div>
    <h1>Welcome to PlanArc</h1>
    <div v-if="username">
      <p>ようこそ、{{ username }}さん！</p>
    </div>
    <div v-else>
      <p>ユーザー情報を取得中...</p>
    </div>
    <div class="mt-3">
      <a @click="logout" href="">ログアウト</a>
    </div>
  </div>
</template>

<style scoped>

</style>