<script setup lang="ts">
import { user } from '@/logic/user.js';
import { ref } from 'vue';

let username = ref('');
let password = ref('');
let confirmPassword = ref('');
let errorMessage = ref('');
let res;

async function handleSubmit() {
  if (password.value !== confirmPassword.value) {
    errorMessage.value = "パスワードが一致しません。";
    return;
  }
  res = await user.register(username.value, password.value);
  console.log(res);
  if (res.ok) {
    location.href = '/';
  }else{
    errorMessage.value = res.error;
  }
}
</script>

<template>
  <h1 class="mb-4">サインアップ</h1>
  <form @submit.prevent="handleSubmit">
    <div class="mb-3">
<!--      <label for="username" class="form-label">ユーザー名:</label>-->
      <input type="text" id="username" v-model="username" placeholder="ユーザ名" required />
    </div>
    <div class="mb-3">
<!--      <label for="password" class="form-label">パスワード:</label>-->
      <input type="password" id="password" v-model="password" placeholder="パスワード" required />
    </div>
    <div class="mb-3">
<!--      <label for="confirmPassword" class="form-label">パスワード確認:</label>-->
      <input type="password" id="confirmPassword" v-model="confirmPassword" placeholder="パスワード確認:" required />
    </div>
    <div class="mb-3">
      <button type="submit" class="btn">サインアップ</button>
    </div>
    <div v-if="errorMessage" class="alert">
      {{ errorMessage }}
    </div>
<!--    <div class="mt-3">-->
<!--      <a href="/login">ログイン</a>-->
<!--    </div>-->
  </form>
</template>

<style scoped>

</style>