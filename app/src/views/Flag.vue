<script setup lang="ts">
  import Footer from "@/components/Footer.vue";
  import Header from "@/components/Header.vue";
  import UnfinishedList from "@/views/UnfinishedList.vue";
</script>

<template>
  <div>
    <Header />
    <div>
      <unfinished-list />
    </div>
    <Footer />
  </div>
</template>

<style scoped>

</style>