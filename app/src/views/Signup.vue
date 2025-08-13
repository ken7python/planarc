<script setup lang="ts">
import { user } from '@/logic/user.js';
import { ref } from 'vue';
import Header from "@/components/Header.vue";

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
  <Header />
  <div id="SignupView">
    <h1 class="mb-4">ユーザ登録</h1>
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
        <button type="submit" class="btn">ユーザ登録</button>
      </div>
      <div v-if="errorMessage" class="alert">
        {{ errorMessage }}
      </div>
  <!--    <div class="mt-3">-->
  <!--      <a href="/login">ログイン</a>-->
  <!--    </div>-->
    </form>
  </div>
</template>
<style scoped>
  #SignupView {
    margin-left: 2vw;
    margin-top: 70px;
  }
</style>