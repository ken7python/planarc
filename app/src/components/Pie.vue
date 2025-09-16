<script setup lang="ts">
  import { subjectModule } from '@/logic/subject';
  import { CONST } from '@/logic/const';
  import { studyLog } from "@/logic/StudyLog";
  import {onMounted, ref, watch, onUnmounted} from "vue";

  let chartInstance = null; // チャートインスタンスを保存

  function drawPieChart() {
    if (subjects.value && subjects.value.length > 0) {
      const ctx = (document.getElementById('pieChart') as HTMLCanvasElement).getContext('2d');
      if (!ctx) return;

      // 既存のチャートがあれば破棄
      if (chartInstance) {
        chartInstance.destroy();
        chartInstance = null;
      }

      chartInstance = new Chart(ctx, {
        type: 'pie',
        data: {
          labels: subjects.value.map(s => s.Name),
          datasets: [{
            data: subjects.value.map(s => s.StudyTime),
            backgroundColor: subjects.value.map(s => s.Color), // Colorプロパティがなければ黒で仮置き
          }]
        },
        options: {
          responsive: true,
          rotation: -90,
          circumference: 180,
          plugins: {
            legend: { position: 'bottom' }
          }
        }
      });
    }
  }

  // コンポーネント破棄時にチャートも破棄
  onUnmounted(() => {
    if (chartInstance) {
      chartInstance.destroy();
      chartInstance = null;
    }
  });

  let subjects = ref<any[]>([]);
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
     console.log(log.SubjectID);
     subjectSet.add(log.SubjectID);
   })

   let subject_list_included = [];
   subject_list.map(subject => {
      if (subjectSet.has(subject.ID)) {
        // subjects.value.push(subject);
        subject_list_included.push(subject);
      }
   })

   console.log(subjectSet);

   // console.log(subject_list_included);

   subject_list_included.map(subject => {
     // console.log(subject);
     const studyTime = studyLogs.filter(log => log.SubjectID === subject.ID)
       .reduce((total, log) => total + log.StudyTime, 0);
     // console.log(subject.ID,studyTime);
     subject["StudyTime"] = studyTime; // 各科目に合計学習時間を追加
   })

   // console.log(subject_list_included);

   subject_list_included.sort((a, b) => b.StudyTime - a.StudyTime);
   console.log(subject_list_included);
   subjects.value = subject_list_included;

   // subjects
   // sort((a, b) => a.StudyTime - b.StudyTime);
   console.log(subjects.value);

   drawPieChart();
 }

 loadData();

  watch(subjects, () => {
    drawPieChart();
  }, {deep: true});
</script>

<template>
  <div>
    <canvas id="pieChart"></canvas>
  </div>
</template>

<style scoped>
  #pieChart {
    max-width: 360px;
    max-height: 150px;
  }
</style>