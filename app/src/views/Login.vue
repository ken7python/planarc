<script setup lang="ts">
  import { user } from '@/logic/user.js';
  import { ref } from 'vue';
  import Header from "@/components/Header.vue";

  const username = ref('');
  const password = ref('');
  const errorMessage = ref('');
  let res;

  const communication_sending = ref<boolean>(false);

  async function login_() {
    if (!communication_sending.value) {
      communication_sending.value = true;
      res = await user.login(username.value, password.value);
      console.log(res.error);
      if (await res.ok) {
        location.href = '/';
      } else {
        errorMessage.value = res.error;
        communication_sending.value = false;
      }
    }
  }
</script>

<template>
  <Header />
  <div id="LoginView">
      <h2>ログイン</h2>
      <form @submit.prevent="login_">
        <div class="mb-3">
<!--          <label for="username" class="form-label">ユーザ名</label>-->
          <input type="text" id="username" v-model="username" placeholder="ユーザ名" required />
        </div>
        <div class="mb-3">
<!--          <label for="password" class="form-label">パスワード</label>-->
          <input type="password" id="password" v-model="password" placeholder="パスワード" required />
        </div>
        <button class="btn" :disabled="communication_sending">
          <span v-if="!communication_sending">ログイン</span>
          <span v-if="communication_sending">通信中...</span>
        </button>
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
  #LoginView {
    margin-left: 2vw;
    margin-top: 70px;
  }
</style>