<script setup lang="ts">
import {ref, watch} from 'vue';
  import { selectStyle } from '@/logic/style/selectStyle';
  import { subjectModule } from '@/logic/subject';
  import { studyLog } from "@/logic/StudyLog";
  import { stopwatch } from '@/logic/StudyLog';

  import writeIcon from '@/assets/icons/write.svg';

  const props = defineProps({
    date: String
  })

  const communication_saving = ref<boolean>(false);

  let subjectID = ref<string>('');

  let startTime = ref<string>(null);
  let endTime = ref<string>(null);

  let subjects = ref<any[]>([]);
  async function loadData() {
    const subject_list = await subjectModule.getList();
    //console.log(subject_list);
    subjects.value = subject_list;
    //console.log(subject_list);
    subjects.value = subject_list;

    stopwatch.init();
    subjectID.value = stopwatch.subject.value;

    await studyLog.getLog(props.date);
  }

  loadData();

  watch(subjectID, (newVal, oldVal) => {
    if (newVal != undefined) {
      //console.log(newVal);
      //console.log(stopwatch);
      stopwatch.subject.value = newVal;
      stopwatch.save();
      stopwatch.init();
    }
  });

  async function record() {
    communication_saving.value = true;
    console.log(props.date);
    const ok = await studyLog.writeStr(props.date,subjectID.value,startTime.value,endTime.value)
    console.log(ok);
    if (await ok) {
      startTime.value = null;
      endTime.value = null;
      // subjectID.value = '';
    }
    communication_saving.value = false;
  }
</script>

<template>
<!--  {{ date }}-->
  <div id="headerinput">
    <select class="selectbox" :style="selectStyle.getSelectStyle(subjectID)" v-model="subjectID">
      <option disabled value="">科目を選択</option>
      <option v-for="subject in subjects" :key="subject.value" :value="subject.ID">
        {{ subject.Name }}
      </option>
    </select>
    <div id="inputbox">
      <div>
        <h2>開始時間</h2>
        <input type="time" id="startTime" name="startTime" value="--:--" v-model="startTime">
      </div>
      <div>
        <h2>終了時間</h2>
        <input type="time" id="startTime" name="startTime" value="--:--" v-model="endTime">
      </div>
    </div>
    <div id="res">
      <button @click="record" class="btn"  style="margin: 0 auto;" :disabled="communication_saving">
        <writeIcon v-if="!communication_saving"></writeIcon>
        <span v-if="!communication_saving">記録</span>

        <span v-if="communication_saving">通信中...</span>
      </button>
    </div>
  </div>
</template>

<style scoped>
  #headerinput {
    text-align: center;
  }

  #inputbox {
    display: flex;
    justify-content: space-around;
    align-items: center;
    margin: 10px 0;
    margin-bottom: 80px;
  }
</style>
