<script setup lang="ts">
  import {ref} from "vue";
  import { user } from "@/logic/user.js";
  import { getColorboxStyle } from "@/logic/style/colorbox";
  import moveIcon from '@/assets/icons/move.svg';
  import deleteIcon from '@/assets/icons/delete.svg';

  const username = ref<string | null>(null);

  const unfinishedTask = ref([
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

  async function loadData() {
    const profile = await user.profile();
    // console.log(profile);
    username.value = profile.username;
  }
  loadData();
</script>

<template>
  <div id="page">
    <div id="List">
      <p style="color: white;line-height: 0">未完了リストのサンプル(まだ追加できません)</p>
      <ul class="list-ul">
        <li class="list-item" v-for="(task, index) in unfinishedTask" :key="index">
          <div>
            <span :style="getColorboxStyle(task.color)" style="margin-right: 4px;margin-left: 4px;"></span>
            <span>{{ task.name }}</span>
          </div>
          <div class="right">
            <button class="squareBtn btnEdit" style="margin-right: 4px;margin-left: 4px;"><moveIcon></moveIcon></button>
            <button class="squareBtn btnTrash" style="margin-right: 4px;margin-left: 4px;"><deleteIcon></deleteIcon></button>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<style scoped>
  #List{
    height: calc(100vh - 145px);
    overflow-y: scroll;
  }
</style>