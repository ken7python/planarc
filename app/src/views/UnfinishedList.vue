<script setup lang="ts">
  import {ref} from "vue";
  import { user } from "@/logic/user.js";
  import { getColorboxStyle } from "@/logic/style/colorbox";
  import { unfinishedModule } from "../logic/unfinished";
  import { subjectModule } from "@/logic/subject";
  import moveIcon from '@/assets/icons/move.svg';
  import deleteIcon from '@/assets/icons/delete.svg';
  import {todoModule} from "../logic/todo";

  const username = ref<string | null>(null);

  let subjects = ref<any[]>([]);
  const unfinishedTask = ref([]);

  async function loadData() {
    const profile = await user.profile();
    // console.log(profile);

    const subject_list = await subjectModule.getList();
    console.log(subject_list);
    subjects.value = subject_list;

    const unfinished = await unfinishedModule.getList();

    let i :number = 0;
    while (i < unfinished.length) {
      unfinished[i]["Color"] = subjects.value.find((subject) => subject.ID === unfinished[i].SubjectID)?.Color || '#000000';
      ++i;
    }

    unfinishedTask.value = unfinished;
    console.log(unfinished);
    username.value = profile.username;
  }
  loadData();

  async function deleteList(id :number) {
    await unfinishedModule.delete(id);
    loadData();
  };

  async function back(id :number) {
    await unfinishedModule.back(id);
    loadData();
  }
</script>

<template>
  <div id="page">
    <div id="List">
      <ul class="list-ul" v-if="unfinishedTask.length > 0">
        <li class="list-item" v-for="(task, index) in unfinishedTask" :key="index">
          <div class="left-group">
            <span :style="getColorboxStyle(task.Color)" style="margin-right: 4px;margin-left: 4px;"></span>
            <span class="task-title">{{ task.Title }}</span>
          </div>
          <div class="right">
            <button class="squareBtn btnEdit" style="margin-right: 4px;margin-left: 4px;" @click="back(task.ID)"><moveIcon></moveIcon></button>
            <button class="squareBtn btnTrash" style="margin-right: 4px;margin-left: 4px;" @click="deleteList(task.ID)"><deleteIcon></deleteIcon></button>
          </div>
        </li>
      </ul>
      <div v-else>
        <p style="color: white;text-align: center">未完了リストは空です</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
  #List{
    height: calc(100vh - 114px);
    overflow-y: scroll;
  }
</style>