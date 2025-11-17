<script setup lang="ts">
  import { ref } from "vue";
  import { CONST } from "@/logic/const";
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

  const communication_loading = ref<boolean>(false);

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

    communication_loading.value = false;
  }

  communication_loading.value = true;
  loadData();

  async function deleteList(id :number) {
    await unfinishedModule.delete(id);
    DeletedialogVisible.value = false;
    loadData()
  }

  const dialogVisible = ref(false);
  const date = ref(CONST.getToday());
  const backId = ref(0);

  async function back(id :number, dateStr :string) {
    if (dateStr === '') {
      alert('日付を選択してください');
      return;
    }
    dialogVisible.value = false;
    await unfinishedModule.back(id, dateStr);
    loadData();
  }
  async function selectBackDate(id :number){
    dialogVisible.value = true;
    backId.value = id;
  }

  const DeletedialogVisible = ref(false);
  const nonCheckDelete = ref(false);
  const deleteTaskID = ref(0);
  function deleteTask(id :number) {
    if (nonCheckDelete.value) {
      deleteList(id);
      return;
    }
    DeletedialogVisible.value = true;
    deleteTaskID.value = id;
  }
  </script>

<template>
  <div>
    <el-dialog v-model="DeletedialogVisible" width="90vw" title="未完了リスト削除確認">
      <p>本当に未完了リストから削除しますか？</p>
      <el-checkbox v-model="nonCheckDelete">次回からは確認しない</el-checkbox>
      <template #footer>
        <el-button @click="DeletedialogVisible = false">キャンセル</el-button>
        <el-button type="primary" @click="deleteList(deleteTaskID)">OK</el-button>
      </template>
    </el-dialog>
  </div>

  <div>
    <el-dialog v-model="dialogVisible" width="90vw" title="日付選択">
      <input type="date" v-model="date" >
      <template #footer>
        <el-button @click="dialogVisible = false">キャンセル</el-button>
        <el-button type="primary" @click="back(backId, date)">OK</el-button>
      </template>
    </el-dialog>
  </div>

  <div id="page">
    <div id="List" v-if="!communication_loading">
      <ul class="list-ul" v-if="unfinishedTask.length > 0" :key="unfinishedTask">
        <li class="list-item" v-for="(task, index) in unfinishedTask" :key="index">
          <div class="left-group">
            <span :style="getColorboxStyle(task.Color)" style="margin-right: 4px;margin-left: 4px;"></span>
            <span class="task-title">{{ task.Title }}</span>
          </div>
          <div class="right">
            <button class="squareBtn btnEdit" style="margin-right: 4px;margin-left: 4px;" @click="selectBackDate(task.ID)"><moveIcon></moveIcon></button>
            <button class="squareBtn btnTrash" style="margin-right: 4px;margin-left: 4px;" @click="deleteTask(task.ID)"><deleteIcon></deleteIcon></button>
          </div>
        </li>
      </ul>
      <div v-else>
        <p style="color: white;text-align: center">未完了リストは空です</p>
      </div>
    </div>
    <div v-else>
      <p style="text-align: center; margin-top: 20px;">通信中...</p>
    </div>
  </div>
</template>

<style scoped>
  #List{
    height: calc(100vh - 86px);
    overflow-y: scroll;
  }
</style>