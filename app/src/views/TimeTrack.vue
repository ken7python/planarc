<script setup lang="ts">
import { ref } from 'vue';
import {studyLog} from "@/logic/StudyLog";
import {subjectModule} from "@/logic/subject";

const props = defineProps({
    date: String
})

const log = ref<any[]>([]);
const subjects = ref<any[]>([]);

const results = ref<any[]>([]);

const sum = ref<number>(0)

const communication_loading = ref<boolean>(false)

async function loadData() {
  communication_loading.value = true;

  log.value = await studyLog.getLog(props.date);
  subjects.value = await subjectModule.getList();

  console.log(log.value);

  for (let i = 0; i < log.value.length; i++) {
    // console.log(log.value[i].SubjectID);
    const subject = subjects.value.find(s => s.ID === log.value[i].SubjectID);
    if (subject) {
      log.value[i].subjectName = subject.Name;
    } else {
      log.value[i].subjectName = "不明";
    }

    const startHours = log.value[i].StartHours;
    const startMinutes = log.value[i].StartMinutes;
    const endHours = log.value[i].EndHours;
    const endMinutes = log.value[i].EndMinutes;

    console.log(startHours, startMinutes, endHours);

    sum.value += log.value[i].StudyTime;

    results.value.push({
      "name": log.value[i].subjectName,
      "StudyTime": log.value[i].StudyTime,
      "sHours": startHours,
      "sMinutes": startMinutes,
      "eHours": endHours,
      "eMinutes": endMinutes,
    });

    communication_loading.value = false;
  }
}

loadData();
</script>

<template>
  <div v-if="!communication_loading">
    <p>TIMETRACKを作成予定</p>
    <p>今は本日の勉強記録を掲載</p>
<!--    {{ log }}-->

    <ul>
      <li v-for="res in results" :key="res.ID">
        <p>
          <span>{{ res.name }}　</span>
          <span v-if="res.sHours != null && res.sHours<10">0{{ res.sHours }}</span>
          <span v-else>{{ res.sHours }}</span>
          <span v-if="res.sMinutes != null">:</span>
          <span v-if="res.sMinutes != null && res.sMinutes<10">0{{ res.sMinutes }}</span>
          <span v-else>{{ res.sMinutes }}</span>

          <span>〜</span>

          <span v-if="res.eHours != null && res.sHours<10">0{{ res.eHours }}</span>
          <span v-else>{{ res.eHours }}</span>
          <span v-if="res.eMinutes != null">:</span>
          <span v-if="res.eMinutes != null && res.eMinutes<10">0{{ res.eMinutes }}</span>
          <span v-else>{{ res.eMinutes }}</span>

          <span>　{{ res.StudyTime }}分</span>


        </p>
      </li>
    </ul>

    <p>本日合計：<span v-if="Math.floor(sum / 60) > 0">{{ Math.floor(sum / 60) }}時間</span><span v-if="sum % 60 > 0">{{sum % 60}}分</span></p>

  </div>

  <div v-else style="text-align: center; margin-top: 20px;">
    <p>通信中...</p>
  </div>
</template>

<style scoped>

</style>