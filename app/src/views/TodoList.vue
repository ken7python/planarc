<script setup lang="ts">
  import { ref, onMounted } from 'vue';
  import { selectStyle } from '@/logic/style/selectStyle';
  import { todoModule } from "../logic/todo";
  import { mic } from '@/logic/mic';
  import { subjectModule } from "@/logic/subject";
  import { unfinishedModule } from "../logic/unfinished";
  import ToDoCard from "../components/ToDoCard.vue";

  import Addicon from '@/assets/icons/add.svg';
  import MicIcon from '@/assets/icons/mic.svg';

  let subjectName = ref('');
  let todoText = ref('');
  let status = ref('MUST');

  const props = defineProps({
    date: String
  })

  const communication_loading = ref<boolean>(false);
  const communication_saving = ref<boolean>(false);

  const TODO_MUST = ref<any[]>([]);
  const TODO_WANT = ref<any[]>([]);
  const TODO_checked = ref<any[]>([]);

  let subjects = ref<any[]>([]);

  async function loadData() {
    communication_loading.value = true;

    const subject_list = await subjectModule.getList();
    console.log(subject_list);
    subjects.value = subject_list;

    const ToDoList = await todoModule.getListGroup(props.date);
    console.log(await ToDoList);

    TODO_MUST.value = await ToDoList.MUST;
    TODO_WANT.value = await ToDoList.WANT;
    TODO_checked.value = await ToDoList.checked;

    console.log("MUST---");
    console.log(TODO_MUST.value);

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

      <div>
        <div id="MUST" class="frame">
          <span>MUST</span>
          <ToDoCard :LIST="TODO_MUST" @reload="loadData" />
        </div>


        <div id="WANT" class="frame">
          <span>WANT</span>
          <ToDoCard :LIST="TODO_WANT" @reload="loadData" />
        </div>


        <div id="Checked" class="frame">
          <span>完了</span>
          <ToDoCard :LIST="TODO_checked" @reload="loadData" />
        </div>
      </div>
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
    height: calc(100dvh - 80px - 34px - 40px - 10px);
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

  .frame {
    border: 1px solid ##e4f2ff;
    background-color: #e4f2ff;
    border-radius: 8px;
    margin: 5px;
    padding: 5px;
  }
</style>