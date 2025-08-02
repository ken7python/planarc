<script setup lang="ts">
  import { login } from '@/assets/user.js';
  import { ref } from 'vue';

  const username = ref('');
  const password = ref('');
  const errorMessage = ref('');
  let res;

  async function login_() {
    res = await login(username.value, password.value);
    console.log(res.error);
    if (res.ok) {
      location.href = '/';
    } else {
      errorMessage.value = res.error;
    }
  }
</script>

<template>
  <div>
      <h2>ログイン</h2>
      <form @submit.prevent="login_">
        <div class="mb-3">
<!--          <label for="username" class="form-label">ユーザ名</label>-->
          <input type="text" id="username" v-model="username" placeholder="ユーザ名">
        </div>
        <div class="mb-3">
<!--          <label for="password" class="form-label">パスワード</label>-->
          <input type="password" id="password" v-model="password" placeholder="パスワード">
        </div>
        <button type="submit" class="btn">ログイン</button>
      </form>

      <div v-if="errorMessage != ''" class="alert">
        {{ errorMessage }}
      </div>

<!--      <div class="mt-3">-->
<!--        <a href="/signup">ユーザ登録</a>-->
<!--      </div>-->
  </div>
</template>

<style scoped>

</style>