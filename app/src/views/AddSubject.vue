<script setup lang="ts">
  import { ref } from 'vue';
  import {getColorboxStyle} from "@/logic/style/colorbox";
  import { subjectModule } from '@/logic/subject';
  import { mic } from '@/logic/mic';

  import AddIcon from '@/assets/icons/add.svg';
  import EditIcon from '@/assets/icons/edit.svg';
  import SaveIcon from '@/assets/icons/save.svg';
  import MicIcon from '@/assets/icons/mic.svg';
  import Header from "../components/Header.vue";
  import Footer from "../components/Footer.vue";

  let subjectName = ref<string>('');
  let subjectColor = ref<string>('#000000');

  const subjects = ref<any[]>([]);
  let editID = ref<number>(0);

  const communication_loading = ref<boolean>(false);
  const communication_saving = ref<boolean>(false);

  async function loadData() {
    communication_loading.value = true;
    const subject_list = await subjectModule.getList();
    communication_loading.value = false;

    subjectName.value = '';
    subjectColor.value = '#000000';
    console.log(subject_list);
    subjects.value = subject_list;
  }
  loadData();

  async function add() {
    communication_saving.value = true;
    await subjectModule.add(subjectName.value, subjectColor.value);
    communication_saving.value = false;
    await loadData();
    mic.result = '';
  }

  async function edit() {
    let id = subjects.value.find(item => item.ID === editID.value);
    await subjectModule.edit(editID.value, id.Name);
    editID.value = 0;
    await loadData();
  }

  async function micbtn() {
    // alert('マイクボタンが押されました。');
    if (!mic.shouldRestart.value) {
      console.log("start");
      mic.result = '';
      // mic.micON.value = true;
      await mic.start();
    } else {
      console.log("stop");
      // mic.micON.value = false;
      await mic.stop();
    }
  }

  mic.rec.onresult = (e) => {
    let interim = '';
    for (let i = e.resultIndex; i < e.results.length; i++) {
      const transcript = e.results[i][0].transcript;
      if (e.results[i].isFinal) {
        mic.result = transcript;
        mic.tmp = '';
        subjectName.value = mic.result;
      } else {
        interim += transcript;
        mic.tmp = interim;
        subjectName.value = mic.result;
      }
    }
  };
</script>

<template>
  <Header></Header>
  <div id="subjectMain">
    <div>
      <div id="addSubject">
        <div class="micdiv">
          <input type="text" placeholder="科目名を入力" v-model="subjectName" />
          <MicIcon class="mic" @click="micbtn" :style="mic.micStyle(true)"></MicIcon>
        </div>
        <br>

        <label class="color-field">
          <span class="color-label">色を選択：</span>
          <input type="color" value="#5B8DEF" aria-label="Select accent color" v-model="subjectColor" />
        </label>
        <br>

        <button class="btn" style="margin: 0 auto;" @click="add" :disabled="!subjectName || !subjectColor || communication_saving">
          <span v-if="communication_saving">通信中</span>

          <AddIcon v-if="!communication_saving" class="icon"></AddIcon>
          <span v-if="!communication_saving">追加</span>
        </button>
      </div>
      <br>
    </div>

    <div id="List" v-if="communication_loading === false">
  <!--    <p style="color: white;line-height: 0">科目リストのサンプル(まだ追加できません)</p>-->
      <div v-if="subjects.length === 0" style="text-align: center; margin-top: 20px;">
        <p style="color: white;line-height: 0">科目がまだ追加されていません</p>
      </div>
      <ul class="list-ul" v-if="subjects" v-for="(subject, index) in subjects" :key="index">
        <li class="list-item">
          <div class="left-group">
            <span :style="getColorboxStyle(subject.Color)" style="margin-right: 4px;margin-left: 4px;"></span>
            <input type="text" v-if="editID === subject.ID" v-model="subject.Name" class="task-input" style="height: 20px;">
            <span class="task-title" v-else>{{ subject.Name }}</span>
          </div>
          <div class="right">
            <button v-if="subject.ID != editID" @click="editID=subject.ID" class="squareBtn btnEdit" style="margin-right: 4px;margin-left: 4px;">
              <EditIcon></EditIcon>
            </button>
            <button v-else @click="edit" class="squareBtn btnSave" style="margin-right: 4px;margin-left: 4px;">
              <SaveIcon></SaveIcon>
            </button>
          </div>
        </li>
      </ul>
      <p style="text-align: center;color: white;" v-else>サーバとの通信に失敗しました</p>
    </div>
    <div v-else style="text-align: center; margin-top: 20px;">
      <p>通信中...</p>
    </div>
  </div>

  <Footer></Footer>
</template>

<style scoped>
  #subjectMain {
    height: calc(100dvh - 50px - 30px); /* ヘッダー80px / フッター60px */
    display: grid;
    grid-template-rows: auto 1fr; /* ヘッダー=auto / 中央=残り / フッター=auto */
  }
  #addSubject {
    text-align: center;
  }

  .color-field {
    display: inline-flex;
    align-items: center;
    gap: 8px;
    font-size: 14px;
  }

  .color-field input[type="color"] {
    width: 40px;
    height: 40px;
    border: none;
    padding: 0;
    border-radius: 6px;
    cursor: pointer;
    appearance: none;
  }

  .color-field input[type="color"]::-webkit-color-swatch {
    border: none;
    border-radius: 4px;
  }

  .color-field input[type="color"]::-moz-color-swatch {
    border: none;
    border-radius: 4px;
  }

</style>