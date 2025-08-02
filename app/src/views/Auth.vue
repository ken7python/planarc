<script setup lang="ts">

</script>

<template>
  <div>
    <div id="main">
      <router-link to="/signup" class="btn">ユーザ登録</router-link>
      <router-link to="/login" class="btn">ログイン</router-link>
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