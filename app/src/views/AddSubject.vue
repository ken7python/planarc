<script setup lang="ts">
  import { ref } from 'vue';
  import {getColorboxStyle} from "@/logic/style/colorbox";
  import { subjectModule } from '@/logic/subject';

  import AddIcon from '@/assets/icons/add.svg';
  import EditIcon from '@/assets/icons/edit.svg';
  import SaveIcon from '@/assets/icons/save.svg';
  import MicIcon from '@/assets/icons/mic.svg';

  let subjectName = ref<string>('');
  let subjectColor = ref<string>('#000000');

  // const subjects = ref([
  //   { name: '国語', color: '#FF5733' },
  //   { name: '数学', color: '#33FF57' },
  //   { name: '英語', color: '#3357FF' },
  //   { name: '理科', color: '#F1C40F' },
  //   { name: '社会', color: '#8E44AD' },
  //   { name: '歴史', color: '#FF8C00' },
  //   { name: '地理', color: '#00CED1' },
  //   { name: '物理', color: '#8A2BE2' },
  //   { name: '化学', color: '#FF4500' },
  //   { name: '生物', color: '#2E8B57' }
  // ]);
  const subjects = ref<any[]>([]);
  let editID = ref<number>(0);

  async function loadData() {
    const subject_list = await subjectModule.getList();
    subjectName.value = '';
    subjectColor.value = '#000000';
    console.log(subject_list);
    subjects.value = subject_list;
  }
  loadData();

  async function add() {
    await subjectModule.add(subjectName.value, subjectColor.value);
    await loadData();
  }

  async function edit() {
    let id = subjects.value.find(item => item.ID === editID.value);
    await subjectModule.edit(editID.value, id.Name);
    editID.value = 0;
    await loadData();
  }
</script>

<template>
  <div>
    <div id="addSubject">
      <div class="micdiv">
        <input type="text" placeholder="科目名を入力" v-model="subjectName" />
        <MicIcon class="mic"></MicIcon>
      </div>
      <br>

      <label class="color-field">
        <span class="color-label">色を選択：</span>
        <input type="color" value="#5B8DEF" aria-label="Select accent color" v-model="subjectColor" />
      </label>
      <br>

      <button class="btn" style="margin: 0 auto;" @click="add" :disabled="!subjectName || !subjectColor">
        <AddIcon class="icon"></AddIcon>
        追加
      </button>
    </div>
    <br>
  </div>

  <div id="List">
<!--    <p style="color: white;line-height: 0">科目リストのサンプル(まだ追加できません)</p>-->
    <div v-if="subjects.length === 0" style="text-align: center; margin-top: 20px;">
      <p style="color: white;line-height: 0">科目がまだ追加されていません</p>
    </div>
    <ul class="list-ul" v-if="subjects">
      <li class="list-item" v-for="(subject, index) in subjects" :key="index">
        <div>
          <span :style="getColorboxStyle(subject.Color)" style="margin-right: 4px;margin-left: 4px;"></span>
          <input type="text" v-if="editID === subject.ID" v-model="subject.Name" style="width: 64vw;height: 20px;">
          <span v-else>{{ subject.Name }}</span>
        </div>
        <div class="right">
          <button v-if="subject.ID != editID" @click="editID=subject.ID" class="squareBtn btnEdit" style="margin-right: 4px;margin-left: 4px;">
            <EditIcon></EditIcon>
          </button>
          <button v-else @click="edit" class="squareBtn btnTrash" style="margin-right: 4px;margin-left: 4px;">
            <SaveIcon></SaveIcon>
          </button>
        </div>
      </li>
    </ul>
    <p style="text-align: center;color: white;" v-else>サーバとの通信に失敗しました</p>
  </div>
</template>

<style scoped>
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