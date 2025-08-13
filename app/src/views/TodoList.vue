<script setup lang="ts">
  import { ref } from 'vue';
  import { selectStyle } from '@/logic/style/selectStyle';
  import {getColorboxStyle} from "@/logic/style/colorbox";

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

  const TODO = ref([
    { name: '国語', color: '#FF5733' },
    { name: '数学', color: '#33FF57' },
    { name: '英語', color: '#3357FF' },
    { name: '理科', color: '#F1C40F' },
    { name: '社会', color: '#8E44AD' },
    { name: '歴史', color: '#FF8C00' },
    { name: '地理', color: '#00CED1' },
    { name: '物理', color: '#8A2BE2' },
    { name: '化学', color: '#FF4500' },
    { name: '生物', color: '#2E8B57' }
  ]);
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
      <button class="btn" style="margin: 0 auto;">追加</button>
    </div>
    <br>
  </div>

  <div id="List">
    <p style="color: white;line-height: 0">TODOリストのサンプル(まだ追加できません)</p>
    <ul class="list-ul">
      <li class="list-item" v-for="(task, index) in TODO" :key="index">
        <div>
          <span :style="getColorboxStyle(task.color)" style="margin-right: 4px;margin-left: 4px;"></span>
          <span>{{ task.name }}</span>
        </div>
        <div class="right">
<!--          <button class="squareBtn btnTrash" style="margin-right: 4px;margin-left: 4px;"></button>-->
          <input type="checkbox" class="squareBtn btnCheck" style="margin-right: 4px;margin-left: 4px;" />
          <button class="squareBtn btnUnfinished" style="margin-right: 4px;margin-left: 4px;"></button>

        </div>
      </li>
    </ul>
  </div>
</template>

<style scoped>
  #AddTodo {
    text-align: center;
  }
  #TODOList {
    text-align: center;
  }
</style>