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

async function loadData() {
  log.value = await studyLog.getLog(props.date);
  subjects.value = await subjectModule.getList();

  for (let i = 0; i < log.value.length; i++) {
    // console.log(log.value[i].SubjectID);
    const subject = subjects.value.find(s => s.ID === log.value[i].SubjectID);
    if (subject) {
      log.value[i].subjectName = subject.Name;
    } else {
      log.value[i].subjectName = "不明";
    }

    results.value.push({
      "name": log.value[i].subjectName,
      "StudyTime": log.value[i].StudyTime,
    });
  }
}

loadData();
</script>

<template>
  <div>
    <p>TIMETRACKを作成予定</p>
    <p>今は本日の勉強記録を掲載</p>
<!--    {{ log }}-->

    <ul>
      <li v-for="res in results" :key="res.ID">
        <p>{{ res.name }}: {{ res.StudyTime }}分</p>
      </li>
    </ul>

  </div>
</template>

<style scoped>

</style>