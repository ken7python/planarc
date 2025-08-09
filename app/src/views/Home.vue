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

  // const onLongpress = (e: CustomEvent) => {
  //   alert("長押しイベントが発生しました: ");
  // }
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
    <button class="btn" @click="logout">Logout</button>
  </div>

<!--    <div>-->
<!--      &lt;!&ndash; デフォルト（500ms） &ndash;&gt;-->
<!--      <div v-use-longpress="1000" @longpress="onLongpress"><button class="btn">長押しのサンプル</button></div>-->
<!--    </div>-->
</template>

<style scoped>

</style>
