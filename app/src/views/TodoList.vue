<script setup lang="ts">
  import { ref } from 'vue';
  let subjectName = ref('');
  let todoText = ref('');
  let status = ref('MUST');
</script>

<template>
  <div>
    <div id="AddTodo">
      <input type="text" placeholder="科目を選択" v-model="subjectName" />
      <input type="text" placeholder="TODOを入力" v-model="todoText" />
      <input type="text" placeholder="MUST or WANT" v-model="status">
      <button class="btn">追加</button>
    </div>
    <div id="TODOList">
      <p>TODOのリストを追加予定</p>
    </div>
  </div>
</template>

<style scoped>

</style>