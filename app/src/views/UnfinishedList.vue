<script setup lang="ts">
  import {ref} from "vue";
  import { user } from "@/logic/user.js";

  const username = ref<string | null>(null);
  async function loadData() {
    const profile = await user.profile();
    // console.log(profile);
    username.value = profile.username;
  }
  loadData();
</script>

<template>
  <div>
    <p>未完了リストを追加予定</p>
  </div>
</template>

<style scoped>

</style>