<script setup lang="ts">
  import { ref } from 'vue';
  import { user } from '@/logic/user.js';

  const username = ref<string | null>(null);
  async function loadData() {
    const profile = await user.profile();
    username.value = profile.username;
  }

  loadData();
</script>

<template>
  <div>
    <div id="welcome" v-if="username">
      <p>ようこそ、{{ username }}さん！</p>
      <hr>
    </div>
    <div v-else>
      <p>ユーザー情報を取得中...</p>
    </div>
    <h2>今日の勉強時間(全体)</h2>
    <p>勉強時間の確認ができるページを作成予定</p>
  </div>
</template>

<style scoped>
  #welcome {
    /* 行間をなくす */
    line-height: 2px;
  }
</style>