<script setup lang="ts">
  import { ref } from 'vue';
  import { selectStyle } from '@/logic/style/selectStyle';
  import { getColorboxStyle } from "@/logic/style/colorbox";
  import { subjectModule } from "@/logic/subject";
  import { todoModule } from "../logic/todo";
  import { mic } from '@/logic/mic';
  import { unfinishedModule } from "../logic/unfinished";

  import Addicon from '@/assets/icons/add.svg';
  import EditIcon from '@/assets/icons/edit.svg';
  import SaveIcon from '@/assets/icons/save.svg';
  import MoveIcon from '@/assets/icons/move.svg';
  import MicIcon from '@/assets/icons/mic.svg';

  let subjectName = ref('');
  let todoText = ref('');
  let status = ref('MUST');

  const props = defineProps({
    date: String
  })

  const communication_loading = ref<boolean>(false);
  const communication_saving = ref<boolean>(false);

  const TODO = ref<any[]>([]);
  let subjects = ref<any[]>([]);
  async function loadData() {
    communication_loading.value = true;
    const subject_list = await subjectModule.getList();
    console.log(subject_list);
    subjects.value = subject_list;

    const ToDoList = await todoModule.getList(props.date);

    let i :number = 0;
    while (i < ToDoList.length) {
      ToDoList[i]["Color"] = subjects.value.find((subject) => subject.ID === ToDoList[i].SubjectID)?.Color || '#000000';
      ++i;
    }

    console.log(ToDoList);
    TODO.value = ToDoList;

    communication_loading.value = false;
  }

  loadData();

  async function add() {
    communication_saving.value = true;
    await todoModule.add(props.date, todoText.value, subjectName.value, status.value);
    todoText.value = "";
    communication_saving.value = false;
    loadData();
  };

  async function move(id :number) {
    const res :boolean = await unfinishedModule.move(id);
    if (res) {
      todoText.value = "";
      loadData();
    }
  };

  async function check(id :number) {
    await todoModule.check(id);
    loadData();
  };

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
    loadData();
  }

  const micbtn = async () => {
    if (!mic.shouldRestart.value) {        // ← .value
      console.log("start");
      mic.result = '';
      mic.micON.value = true;
      await mic.start();
    } else {
      console.log("stop");
      mic.micON.value = false;
      mic.stop();
    }
  };

  // onresult は setup後に一度だけ付ける
  mic.rec && (mic.rec.onresult = (e) => {
    let interim = '';
    for (let i = e.resultIndex; i < e.results.length; i++) {
      const transcript = e.results[i][0].transcript;
      if (e.results[i].isFinal) {
        mic.result = transcript;
        mic.tmp = '';
      } else {
        interim += transcript;
        mic.tmp = interim;
      }
    }

    todoText.value = mic.result || mic.tmp;
  });
</script>

<template>
  <div id="ToDoPage">
    <div>
      <details>
        <summary>リスト作成</summary>
        <div id="AddTodo">
          <select class="selectbox" :style="selectStyle.getSelectStyle(subjectName)" v-model="subjectName">
            <option value="">科目を選択</option>
            <option v-for="subject in subjects" :key="subject.value" :value="subject.ID">
              {{ subject.Name }}
            </option>
          </select>
          <div class="micdiv">
            <input type="text" placeholder="ToDoを入力" v-model="todoText" />
            <MicIcon class="mic" @click="micbtn" :style="mic.micStyle()"></MicIcon>
          </div>
          <select class="selectbox" :style="selectStyle.getSelectStyle(status)" v-model="status">
            <option value="MUST">MUST</option>
            <option value="WANT">WANT</option>
          </select>
          <br>
          <br>
          <button class="btn" style="margin: 0 auto;" @click="add" :disabled="communication_loading">
            <Addicon v-if="!communication_loading"></Addicon>
            <span v-if="!communication_loading">追加</span>

            <span v-if="communication_saving">通信中</span>
          </button>
        </div>
        <br>
      </details>
    </div>

    <div id="List" v-if="!communication_loading">
      <ul class="list-ul" v-for="(task, index) in TODO" :key="index">
        <li class="list-item" style="width: calc(100dvw - 10px);">
          <div class="left-group">
            <span v-if="task.Status === 'MUST'">M</span>
            <span v-if="task.Status === 'WANT'">W</span>
            <span :style="getColorboxStyle(task.Color)" style="margin-right: 4px;margin-left: 4px;"></span>
            <span v-if="editId != task.ID" class="task-title">
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
      <ul v-if="TODO.length === 0" class="list-ul">
        <li style="width: 100%;text-align: center;">
          <div>
            <span style="color: white;">タスクがありません</span>
          </div>
        </li>
      </ul>
    </div>
    <div v-else style="text-align: center; margin-top: 20px;">
      <p>通信中...</p>
    </div>
  </div>
</template>

<style scoped>
  #AddTodo {
    text-align: center;
  }
  #TODOList {
    text-align: center;
  }
  #ToDoPage {
    height: calc(100dvh - 80px - 34px - 40px - 40px);
    display: grid;
    grid-template-rows: auto 1fr
  }

  .selectbox,input[type="text"] {
    padding-bottom: 4px;
    margin-bottom: 0px;
    margin-top: 0px;
  }

  #List {
    margin-right: -5px;
    margin-left: -5px;
  }
</style>