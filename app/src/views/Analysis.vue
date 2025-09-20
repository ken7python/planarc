<script setup lang="ts">
  import { ref, watch } from "vue";
  import Header from "@/components/Header.vue";
  import Footer from "@/components/Footer.vue";

  import { analysisModule } from "@/logic/analysis";
  import { subjectModule } from '@/logic/subject';

  let subjects = ref<any[]>([]);
  let sum = ref<number>(0);

  let chartInstance = null; // チャートインスタンスを保存
  let barChartInstance = null;

  function drawPieChart() {
    if (subjects.value && subjects.value.length > 0) {
      const ctx = (document.getElementById('pieChart') as HTMLCanvasElement).getContext('2d');
      if (!ctx) return;

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

  function drawBarChart(dateArr, datasets) {
    const ctx = document.getElementById('barChart').getContext('2d');

    const data = {
      labels: dateArr,
      datasets: datasets,
    };

    console.log(data);

    const config = {
      type: 'bar',
      data: data,
      options: {
        responsive: true,
        plugins: {
          // title: { display: true, text: '月別：学習＋活動（積み上げ）' },
          legend: {
            position: 'bottom',
            display: false,
          }
        },
        scales: {
          x: {
            stacked: true,
            categoryPercentage: 1.0,
            barPercentage: 1.0
          },   // ← ここで積み上げを有効にする
          y: { stacked: true }    // ← 縦軸も積み上げに
        }
      }
    };

    barChartInstance = new Chart(ctx, config);
  }

  function toLocalYMD(date) {
    const z = n => String(n).padStart(2, '0');
    return `${date.getFullYear()}-${z(date.getMonth()+1)}-${z(date.getDate())}`;
  }

  function getThisMonthBounds(now = new Date()) {
    const year = now.getFullYear();
    const month = now.getMonth(); // 0=1月, 1=2月, ..., 11=12月

    const firstDay = toLocalYMD(new Date(year, month, 1));           // 今月の1日 00:00:00 (ローカル)
    const lastDay  = toLocalYMD(new Date(year, month + 1, 0));      // 「次の月の0日」＝今月の最終日

    return { firstDay, lastDay };
  }

  const startDay = ref<string>('');
  const endDay = ref<string>('');

  const loadOK = ref<boolean>(false);
  const loading = ref<boolean>(true);

  async function loadData() {
    loading.value = true;
    loadOK.value = false;
    sum.value = 0;

    const subject_list = await subjectModule.getList();
    // console.log(subject_list);
    // subjects.value = subject_list;

    const subjectSet = new Set();
    const studyLogs = await analysisModule.get(startDay.value, endDay.value);
    console.log(studyLogs);

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

    subjects.value.map(subject => {
      sum.value += subject.StudyTime;
    })

    if (sum.value > 0) {
      loadOK.value = true;
    }
    loading.value = false;

    console.log(studyLogs);

    const dateSet = new Set();

    let datasets = [];
    const dateArr = [];

    studyLogs.map(log => {
      dateSet.add(log.Date);
    })
    dateSet.forEach(date => {
      const [year, month, day] = date.split('-');
      dateArr.push(`${Number(month)}.${Number(day)}`);
    });

    console.log(dateSet);

    datasets = [];
    subjectSet.forEach(subject => {
      const daily = [];
      const listfind = subject_list.find(s => s.ID === subject);
      const name = listfind ? listfind.Name : '不明';
      const color = listfind ? listfind.Color : '#000000';
      dateSet.forEach(date => {
        let sum = 0;
        studyLogs.forEach(log => {
          if (log.SubjectID === subject && log.Date === date) {
            sum += log.StudyTime;
          }
        })
        daily.push(sum);
      })

      datasets.push({
        label: name,
        data: daily,
        backgroundColor: color,
        stack: '勉強',
      })
    });

    console.log(datasets);
    console.log(dateArr);

    // 既存のチャートがあれば破棄
    if (chartInstance) {
      chartInstance.destroy();
      chartInstance = null;
    }
    if (barChartInstance) {
      barChartInstance.destroy();
      barChartInstance = null;
    }

    if (loadOK.value) {
      drawPieChart();
      drawBarChart(dateArr, datasets);
    }
  }

  const { firstDay, lastDay } = getThisMonthBounds();
  console.log(firstDay);
  console.log(lastDay);

  startDay.value = firstDay;
  endDay.value = lastDay;

  loadData();

  watch(startDay, subject => {
    loadData()
  })

  watch(endDay, subject => {
    loadData()
  })
</script>

<template>
  <header>
    <Header />
  </header>

  <div>
    <div id="main">
        <!-- Analysis content goes here -->
      <div id="datepicker">
        <h2>指定期間</h2>
        <b>開始</b><input type="date" v-model="startDay"  />
        <br>
        <b>終了</b><input type="date" v-model="endDay" />
      </div>

      <div v-if="loading">
        <p>通信中...</p>
      </div>

      <div id="circle">
        <p v-if="sum > 0">期間合計：<span v-if="Math.floor(sum / 60) > 0">{{ Math.floor(sum / 60) }}時間</span><span v-if="sum % 60 > 0">{{sum % 60}}分</span></p>
        <p v-else>期間合計：0時間</p>
        <canvas id="pieChart"></canvas>
      </div>

      <div id="bar">
        <canvas id="barChart"></canvas>
      </div>
    </div>
  </div>
  <footer>
    <Footer />
  </footer>
</template>

<style scoped>
#main {
  height: calc(100dvh - 80px); /* ヘッダーとフッターの高さを引く */
  overflow-y: auto;     /* 縦スクロールを自動で表示 */
  overflow-x: hidden;   /* 横スクロールを隠す（必要なら auto にする） */
}
  #circle, #bar {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    margin-top: 16px;
  }
  #pieChart {
    max-width: 90vw;
    max-height: 32vh;
  }
  #barChart {
    max-width: 90vw;
    max-height: 32vh;
  }
</style>