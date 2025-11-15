<script setup lang="ts">
  import { ref, onMounted } from 'vue';
  import { selectStyle } from '@/logic/style/selectStyle';
  import { todoModule } from "../logic/todo";
  import { mic } from '@/logic/mic';
  import { subjectModule } from "@/logic/subject";
  import { statusModule } from "@/logic/status";
  import ToDoCard from "../components/ToDoCard.vue";

  import { CONST } from '@/logic/const';
  import { user } from '@/logic/user.js';

  import Addicon from '@/assets/icons/add.svg';
  import MicIcon from '@/assets/icons/mic.svg';

  let subjectName = ref('');
  let todoText = ref('');
  let status = ref('MUST');

  import mode1 from '@/assets/icons/1.svg';
  import mode2 from '@/assets/icons/2.svg';
  import mode3 from '@/assets/icons/3.svg';
  import mode4 from '@/assets/icons/4.svg';

  import saveicon from '@/assets/icons/save.svg';

  const props = defineProps({
    date: String
  })

  const communication_loading = ref<boolean>(false);
  const communication_saving = ref<boolean>(false);

  const TODO_MUST = ref<any[]>([]);
  const TODO_WANT = ref<any[]>([]);
  const TODO_checked = ref<any[]>([]);

  let subjects = ref<any[]>([]);

  // const bellIDs = ref<number[]>([]);
  const bellSets = new Set();

  async function loadData() {
    const subject_list = await subjectModule.getList();
    //console.log(subject_list);
    subjects.value = subject_list;

    const ToDoList = await todoModule.getListGroup(props.date);

    const Notify = await todoModule.getNotify(props.date);
    if (Notify.length > 0) {
      await Notify.map((item => {
        // console.log(item);
        bellSets.add(item.to_do_id);
      }));
      console.log(bellSets);
    }

    ToDoList.MUST.map((item => {
      if (bellSets.has(item.ID)) {
        item.bell = true;
      } else {
        item.bell = false;
      }
    }));
    ToDoList.WANT.map((item => {
      if (bellSets.has(item.ID)) {
        item.bell = true;
      } else {
        item.bell = false;
      }
    }));
    ToDoList.checked.map((item => {
      if (bellSets.has(item.ID)) {
        item.bell = true;
      } else {
        item.bell = false;
      }
    }));

    console.log(await ToDoList);

    TODO_MUST.value = await ToDoList.MUST;
    TODO_WANT.value = await ToDoList.WANT;
    TODO_checked.value = await ToDoList.checked;

    console.log("MUST---");
    console.log(TODO_MUST.value);
    communication_loading.value = false;
  }

  const datetime = ref<string>('');

  const AddDialogVisible = ref(false);

  async function add() {
    communication_saving.value = true;
    await todoModule.add(props.date, todoText.value, subjectName.value, status.value, datetime.value);
    todoText.value = "";
    datetime.value = "";
    communication_saving.value = false;
    AddDialogVisible.value = false;
    loadData();
  };


  const tempStr = ref<string>('');
  const tempmode =ref<string>('');

  const micbtn = async () => {
    if (!mic.shouldRestart.value) {        // ← .value
      console.log("start");
      mic.result = '';
      // mic.micON.value = true;
      await mic.start();
    } else {
      console.log("stop");
      // mic.micON.value = false;
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

    tempStr.value = mic.result || mic.tmp;
    if (tempmode.value === 'todo') {
      todoText.value = tempStr.value;
    } else if (tempmode.value === 'enjoyment') {
      enjoyment.value = tempStr.value;
    }
  });

  function micTodo() {
    tempmode.value = 'todo';
    micbtn();
  }
  function micEnjoyment() {
    tempmode.value = 'enjoyment';
    micbtn();
  }

  let enjoyment = ref<string>('');

  async function setEnjoyment(enjoy: string) {
    //console.log(enjoy);
    await statusModule.setEnjoyment(props.date,enjoy);
    loadStatus();
  }

  let selectedMood = ref<number>(null);
  async function moodSelect(mood: number) {
    // selectedMood.value = mood;
    await statusModule.setMood(props.date,mood);
    loadStatus();
  }
  const getMoodStyle = (mood: number) => {
    return {
      color: selectedMood.value === mood ? '#4d5161' : 'gray',
      background: selectedMood.value === mood ? 'yellow' : 'white',
      cursor: 'pointer'
    };
  };


  async function loadStatus() {
    const status = await statusModule.get(props.date);
    //console.log(status);
    if (status) {
      enjoyment.value = status.Enjoyment || '';
      selectedMood.value = status.Mood || null;
    } else {
      enjoyment.value = '';
      selectedMood.value = null;
    }
  }

  communication_loading.value = true;

  loadData();
  loadStatus();
</script>

<template>
  <div>
    <el-dialog v-model="AddDialogVisible" width="90vw" title="ToDo作成">
      <div id="AddTodo">
        <select class="selectbox" :style="selectStyle.getSelectStyle(subjectName)" v-model="subjectName">
          <option disabled value="">科目を選択</option>
          <option v-for="subject in subjects" :key="subject.value" :value="subject.ID">
            {{ subject.Name }}
          </option>
        </select>
        <div class="micdiv">
          <input type="text" placeholder="ToDoを入力" v-model="todoText" />
          <MicIcon class="mic" @click="micTodo" :style="mic.micStyle(tempmode === 'todo')"></MicIcon>
        </div>
        <select class="selectbox" :style="selectStyle.getSelectStyle(status)" v-model="status">
          <option value="MUST">MUST</option>
          <option value="WANT">WANT</option>
        </select>
        <span>通知日時</span>
        <input class="selectbox" type="datetime-local" id="meeting" name="meeting" v-model="datetime">
      </div>
      <template #footer>
        <el-button @click="AddDialogVisible = false">キャンセル</el-button>
        <el-button type="primary" @click="add" :disabled="communication_loading">
          <Addicon v-if="!communication_saving"></Addicon>
          <span v-if="!communication_saving">追加</span>

          <span v-if="communication_saving" >通信中</span>
        </el-button>
      </template>
    </el-dialog>
  </div>

  <div id="ToDoPage">
    <div id="top-menu">
    <div id="statusMenu">
        <details class="frame">
          <summary>今日の楽しみ</summary>
          <div>
            <div>
              <!--        <h3 style="line-height: 0;margin-bottom: 0">今日の楽しみ</h3>-->
              <div style="display: flex; align-items: center; justify-content: center;">
                <div class="micdiv">
                  <input type="text" placeholder="今日の楽しみを入力" style="width: 70vw" v-model="enjoyment" />
                  <mic-icon :style="mic.micStyle(tempmode === 'enjoyment')" @click="micEnjoyment"></mic-icon>
                </div>
                <button class="squareBtn btnSave" style="margin-left: 20px;" @click="setEnjoyment(enjoyment)"><saveicon></saveicon></button>
              </div>
            </div>
          </div>
        </details>
        <details class="frame">
          <summary>今日の気分</summary>
          <div>
            <div id="feeling">
              <!--          <h3 style="line-height: 0;margin-bottom: 0">今日の気分</h3>-->
              <div id="feeling-icons">
                <mode4 @click="moodSelect(4)" :style="getMoodStyle(4)"></mode4>
                <mode3 @click="moodSelect(3)" :style="getMoodStyle(3)"></mode3>
                <mode2 @click="moodSelect(2)" :style="getMoodStyle(2)"></mode2>
                <mode1 @click="moodSelect(1)" :style="getMoodStyle(1)"></mode1>
              </div>
            </div>
          </div>
        </details>
      </div>
    <div>
        <button class="btn" style="margin: 10px auto;" @click="AddDialogVisible = true">
          <Addicon></Addicon>
          <span>ToDo追加</span>
        </button>
    </div>

    </div>
    <div id="List" v-if="!communication_loading">
      <div>
        <div id="MUST" class="frame">
          <span>MUST</span>
          <span v-if="TODO_MUST.length > 0">({{ TODO_MUST.length }})</span>
          <ToDoCard :LIST="TODO_MUST" @reload="loadData" :key="TODO_MUST"/>
        </div>


        <div id="WANT" class="frame">
          <span>WANT</span>
          <span v-if="TODO_WANT.length > 0">({{ TODO_WANT.length }})</span>
          <ToDoCard :LIST="TODO_WANT" @reload="loadData" :key="TODO_WANT" />
        </div>


        <div id="Checked" class="frame">
          <span>完了</span>
          <span v-if="TODO_checked.length > 0">({{ TODO_checked.length }})</span>
          <ToDoCard :LIST="TODO_checked" @reload="loadData" :key="TODO_checked" />
        </div>
      </div>
    </div>
    <div v-else style="color: white;text-align: center; margin-top: 20px;">
      <p>通信中...</p>
    </div>

<!--    <button @click="register()" style="position: fixed; bottom: 80px; right: 20px; z-index: 1000;">-->
<!--      🔔 通知を許可して送信-->
<!--    </button>-->

<!--    &lt;!&ndash; デバッグ用テスト通知ボタン &ndash;&gt;-->
<!--    <button @click="testNotification" style="position: fixed; bottom: 140px; right: 20px; z-index: 1000; background: orange; color: white; padding: 8px; border: none; border-radius: 4px;">-->
<!--      🧪 テスト通知-->
<!--    </button>-->
  </div>
</template>

<style scoped>
  #AddTodo {
    margin-top: 2%;
  }
  #AddTodo span {
    display: flex;
    justify-content: center;
  }
  #ToDoPage {
    /* background-color: #3d7fe0; */
    height: calc(100dvh - 80px - 34px - 40px - 10px);
    display: grid;
    grid-template-rows: auto 1fr;
  }

  #top-menu {
    margin-top: 10px;
    margin-bottom: 0px;
    margin-right: 0px;
    margin-left: 0px;
  }

  .selectbox,input[type="text"] {
    padding-bottom: 4px;
    margin-bottom: 0px;
    margin-top: 0px;
  }

  .frame {
    border: 2px solid var(--border-color);
    background: var(--white-color);
    border-radius: 8px;
    margin: 5px;
    padding: 5px;
  }

  #feeling-icons {
    display: flex;
    justify-content: space-around;
    align-items: center;
    height: 50px;
  }

  h3 {
    line-height: 0.75;
  }
</style>
