<script setup lang="ts">
  // Chart.js円グラフ描画
  import {Chart} from "chart.js/auto";
  import { subjectModule } from '@/logic/subject';
  import { CONST } from '@/logic/const';
  import { studyLog } from "@/logic/StudyLog";
  import {onMounted, ref, watch} from "vue";

  function drawPieChart() {
    if (subjects.value && subjects.value.length > 0) {
      const ctx = (document.getElementById('pieChart') as HTMLCanvasElement).getContext('2d');
      if (!ctx) return;
      new Chart(ctx, {
        type: 'pie',
        data: {
          labels: subjects.value.map(s => s.Name),
          datasets: [{
            data: studyTimeBySubject,
            backgroundColor: subjects.value.map(s => s.Color), // Colorプロパティがなければ黒で仮置き
          }]
        },
        options: {
          responsive: true,
          rotation: -90,
          circumference: 180,
          plugins: {
            legend: { position: 'right' }
          }
        }
      });
    }
  }

  let subjects = ref<any[]>([]);
  const studyTimeBySubject = [];
  const today = CONST.getToday();

 async function loadData() {
   const subject_list = await subjectModule.getList();
   // console.log(subject_list);
   // subjects.value = subject_list;

   const subjectSet = new Set()
   const studyLogs = await studyLog.getLog(today);
   // console.log(studyLogs);
   studyLogs.map(log => {
     // console.log(log.ID);
     subjectSet.add(log.ID);
   })

   subject_list.map(subject => {
      if (subjectSet.has(subject.ID)) {
        subjects.value.push(subject);
      }
   })

   console.log(subjectSet);

   subjects.value.map(subject => {
     // console.log(subject);
     const studyTime = studyLogs.filter(log => log.SubjectID === subject.ID)
       .reduce((total, log) => total + log.StudyTime, 0);
     // console.log(subject.ID,studyTime);
     studyTimeBySubject.push(studyTime);
   })

   drawPieChart();
 }

 loadData();

  watch(subjects, () => {
    drawPieChart();
  }, {immediate: true});
</script>

<template>
  <div>
    <canvas id="pieChart"></canvas>
  </div>
</template>

<style scoped>
  #pieChart {
    max-width: 360px;
    max-height: 100px;
  }
</style>