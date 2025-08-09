<script setup>
  import { useRoute } from 'vue-router'
  const route = useRoute()
  import { CONST } from "@/logic/const.js";
</script>

<template>
  <div>
    <router-view />
  </div>
</template>

<style scoped>
</style>
