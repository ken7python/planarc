<script setup lang="ts">
  import { ref } from 'vue';
  let subjectName = ref<string>('');
  let subjectColor = ref<string>('#000000');
</script>

<template>
  <div>
    <div id="addSubject">
      <input type="text" placeholder="科目名を入力" v-model="subjectName" />
      <p>色を選択<input type="color" v-model="subjectName" /></p>

      <button class="btn">追加</button>
    </div>
    <div id="subjectList">
      <p>科目のリストを追加予定</p>
    </div>
  </div>
</template>

<style scoped>

</style>