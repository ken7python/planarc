<script setup lang="ts">
  // Chart.js円グラフ描画
  import {Chart} from "chart.js/auto";
  import { subjectModule } from '@/logic/subject';
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
            data: subjects.value.map(s => s.value || 1), // valueプロパティがなければ1で仮置き
            backgroundColor: subjects.value.map(s => s.Color), // Colorプロパティがなければ黒で仮置き
          }]
        },
        options: {
          responsive: true,
          plugins: {
            legend: { position: 'bottom' }
          }
        }
      });
    }
  }

  let subjects = ref<any[]>([]);

 async function loadData() {
   const subject_list = await subjectModule.getList();
   console.log(subject_list);
   subjects.value = subject_list;
   watch(subjects, () => {
     drawPieChart();
   }, {immediate: true});
 }

 loadData();
</script>

<template>
  <div>
    <h4>円グラフ（仮に全て１時間としている）</h4>
    <canvas id="pieChart"></canvas>
  </div>
</template>

<style scoped>

</style>