<script setup lang="ts">
  let subjectName = ref<string>('');
  let subjectColor = ref<string>('#000000');
</script>

<template>
  <div>
    <div id="addSubject">
      <input type="text" placeholder="科目名を入力" v-model="subjectName" />
      <span>色を選択</span>
      <input type="color" v-model="subjectName" />
      <button type="button">追加</button>
    </div>
    <div id="subjectList">
      <p>科目のリストを追加予定</p>
  </div>
</template>

<style scoped>

</style>