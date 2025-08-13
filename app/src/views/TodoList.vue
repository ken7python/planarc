<script setup lang="ts">
  import { ref } from 'vue';
  import { selectStyle } from '@/logic/style/selectStyle.ts';

  let subjectName = ref('');
  let todoText = ref('');
  let status = ref('MUST');

  const subjects = [
    { name: "国語", value :"kokugo" },
    { name: "数学", value :"sugaku" },
    { name: "英語", value :"eigo" },
    { name: "理科", value :"rika" },
    { name: "社会", value :"shakai" }
  ]


</script>

<template>
  <div>
    <div id="AddTodo">
      <select class="selectbox" :style="selectStyle.getSelectStyle(subjectName)" v-model="subjectName">
        <option value="">科目を選択</option>
        <option v-for="subject in subjects" :key="subject.value" :value="subject.value">
          {{ subject.name }}
        </option>
      </select>
      <br>
      <input type="text" placeholder="TODOを入力" v-model="todoText" />
      <br>
      <select class="selectbox" :style="selectStyle.getSelectStyle(status)" v-model="status">
        <option value="MUST">MUST</option>
        <option value="WANt">WANT</option>
      </select>
      <br>
      <br>
      <button class="btn">追加</button>
    </div>
    <div id="TODOList">
      <p>TODOのリストを追加予定</p>
    </div>
  </div>
</template>

<style scoped>
  #AddTodo {
    text-align: center;
  }
</style>