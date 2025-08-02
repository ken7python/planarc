<script setup lang="ts">
import { register } from '@/assets/user.js';
import { ref } from 'vue';

let username = ref('');
let password = ref('');
let confirmPassword = ref('');
let errorMessage = ref('');
let res;

async function handleSubmit() {
  if (password.value !== confirmPassword.value) {
    this.errorMessage = "パスワードが一致しません。";
    return;
  }
  res = await register(username.value, password.value);
  console.log(res);
  if (res.ok) {
    location.href = '/';
  }else{
    this.errorMessage = res.error;
  }
}
</script>

<template>
  <h1 class="mb-4">サインアップ</h1>
  <form @submit.prevent="handleSubmit">
    <div class="mb-3">
      <label for="username" class="form-label">ユーザー名:</label>
      <input type="text" id="username" v-model="username" class="form-control" required />
    </div>
    <div class="mb-3">
      <label for="password" class="form-label">パスワード:</label>
      <input type="password" id="password" v-model="password" class="form-control" required />
    </div>
    <div class="mb-3">
      <label for="confirmPassword" class="form-label">パスワード確認:</label>
      <input type="password" id="confirmPassword" v-model="confirmPassword" class="form-control" required />
    </div>
    <div class="mb-3">
      <button type="submit" class="btn btn-primary">サインアップ</button>
    </div>
    <div v-if="errorMessage" class="alert alert-danger">
      {{ errorMessage }}
    </div>
    <div class="mt-3">
      <a href="/login">ログイン</a>
    </div>
  </form>
</template>

<style scoped>

</style>