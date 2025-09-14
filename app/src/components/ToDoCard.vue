<script setup lang="ts">
import {todoModule} from "../logic/todo";
import {ref} from "vue";
import {unfinishedModule} from "../logic/unfinished";
import { getColorboxStyle } from "@/logic/style/colorbox";

import EditIcon from '@/assets/icons/edit.svg';
import SaveIcon from '@/assets/icons/save.svg';
import MoveIcon from '@/assets/icons/move.svg';

const props = defineProps({
  LIST: Array
})

console.log(props.LIST);

const LIST = ref<any[]>(props.LIST);


async function check(id :number) {
  await todoModule.check(id);
  reload();
}

let editId = ref<number>(0);
let editedText = ref<string>('');

function editing(id :number, title :string) {
  editId.value = id;
  editedText.value = title;
}

async function edit() {
  await todoModule.edit(editId.value, editedText.value);
  editId.value = 0;
  editedText.value = '';
  reload();
}

async function move(id :number) {
  const res :boolean = await unfinishedModule.move(id);
  if (res) {
    reload();
  }
}

const emit = defineEmits(["reload"])

function reload() {
  emit("reload") // 親にイベントを通知する
}

function getTitleStyle(checked :boolean) {
  if (checked) {
    return {
      textDecoration: 'line-through',
      color: '#888888'
    }
  } else {
    return {
      textDecoration: 'none',
      color: '#000000'
    }
  }
}
</script>

<template>
    <ul class="list-ul" v-for="(task, index) in LIST" :key="index">
      <li class="list-item" style="width: calc(100dvw - 30px);">
        <div class="left-group">
          <span :style="getColorboxStyle(task.Color)" style="margin-right: 4px;margin-left: 4px;"></span>
          <span v-if="editId != task.ID" class="task-title" :style="getTitleStyle(task.Checked)">
              {{ task.Title }}
            </span>
          <input v-else type="text" v-model="editedText" class="task-input"/>
        </div>
        <div class="right">
          <!--          <button class="squareBtn btnTrash" style="margin-right: 4px;margin-left: 4px;"></button>-->
          <span v-if="editId != task.ID">
              <button class="squareBtn btnEdit" @click="editing(task.ID,task.Title)"><EditIcon></EditIcon></button>
            </span>
          <span v-else>
              <button class="squareBtn btnSave" @click="edit()"><SaveIcon></SaveIcon></button>
            </span>

          <input type="checkbox" class="squareBtn btnCheck" style="margin-right: 4px;margin-left: 4px;" v-model="task.Checked" @click="check(task.ID)" />
          <button class="squareBtn btnUnfinished" style="margin-right: 4px;margin-left: 4px;" @click="move(task.ID)"><MoveIcon></MoveIcon></button>
        </div>
      </li>
    </ul>
    <ul v-if="LIST.length === 0" class="list-ul">
      <li style="width: 100%;text-align: center;">
        <div>
          <span>タスクがありません</span>
        </div>
      </li>
    </ul>
</template>

<style scoped>

</style>