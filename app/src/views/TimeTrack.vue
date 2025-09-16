<script setup lang="ts">
import {  onMounted ,ref } from 'vue';
import {studyLog} from "@/logic/StudyLog";
import {subjectModule} from "@/logic/subject";

import {createCalendar, TimeGrid} from '@event-calendar/core';

const calendar = ref(null)

const props = defineProps({
    date: String
})

const log = ref<any[]>([]);
const subjects = ref<any[]>([]);

const results = ref<any[]>([]);

const sum = ref<number>(0)

const communication_loading = ref<boolean>(false);

function getTimeRange(events) {
  console.log('-getTimeRange-')
  let sHours = Infinity;
  let sMinutes = 0;
  let eHours = -Infinity;
  let eMinutes = 0;

  events.forEach(e => {
    let sD = new Date(new Date(e.start).getTime() - 15 * 60 * 1000);
    let eD = new Date(new Date(e.end).getTime() + 15 * 60 * 1000);
    let sH = sD.getHours();
    let sM = sD.getMinutes();
    let eH = eD.getHours();
    let eM = eD.getMinutes();
    if (sH * 60 + sM < sHours * 60 + sMinutes) sHours = sH;sMinutes = sM;
    if (eH * 60 + eM > eHours * 60 + eMinutes) eHours = eH;eMinutes = eM;
 })
  console.log(sHours, sMinutes, eHours, eMinutes);

  return { sHours, sMinutes, eHours, eMinutes };
}


function getContrastColor(hex) {
  // 先頭の # を削除
  hex = hex.replace("#", "");

  // RGBに分解
  let r = parseInt(hex.substring(0, 2), 16);
  let g = parseInt(hex.substring(2, 4), 16);
  let b = parseInt(hex.substring(4, 6), 16);

  // 明度(YIQ計算式)
  let yiq = (r * 299 + g * 587 + b * 114) / 1000;

  // 明度が128未満なら濃い色とみなして白を返す
  return yiq < 128 ? "#FFFFFF" : "#000000";
}

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
      log.value[i].subjectColor = subject.Color;
    } else {
      log.value[i].subjectName = "不明";
      log.value[i].subjectColor = "#000000";
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
      "color": log.value[i].subjectColor
    });
  }

  console.log(...results.value);

  const events = [
    ...results.value.map(item => ({
      title: item.name,
      start: props.date + 'T' + String(item.sHours).padStart(2, '0') + ':' + String(item.sMinutes).padStart(2, '0') + ':00',
      end: props.date + 'T' + String(item.eHours).padStart(2, '0') + ':' + String(item.eMinutes).padStart(2, '0') + ':00',
      color: item.color,
      textColor: getContrastColor(item.color)
    }))
  ]

  const { sHours,sMinutes,eHours,eMinutes } = getTimeRange(events);
  console.log(sHours,sMinutes,eHours,eMinutes);

  //console.log(`${String(sHours.padStart(2, '0')}:00:00`)

  let ec = createCalendar(
      // HTML element the calendar will be mounted to
      document.getElementById('ec'),
      // Array of plugins
      [TimeGrid],
      // Options object
      {
        view: 'timeGridDay',
        headerToolbar: false,
        allDaySlot: false,
        date: props.date,
        slotDuration: '00:15',

        dayHeaderFormat: function(date) {
          return '';  // 空文字を返して見出しを非表示風にする
        },
        eventContent: function(arg) {
          return {
            html: `<div class="my-event-title">${arg.event.title}</div>`
            // `arg.timeText` を使わないことで「10:00～11:30」の部分を表示しない
          };
        },
        events: events,
        slotMinTime: `${ sHours }:${ sMinutes }:00`, // 最初の予定より前は非表示
        slotMaxTime: `${ eHours }:${ eMinutes }:00`    // 最後の予定より後は非表示
      }
  );

  communication_loading.value = false;
}

onMounted(() => {
  loadData();
})
</script>

<template>
  <div v-show="!communication_loading" id="trackpage">
    <div id="ec" class="calendar"></div>
    <p>本日合計：<span v-if="Math.floor(sum / 60) > 0">{{ Math.floor(sum / 60) }}時間</span><span v-if="sum % 60 > 0">{{sum % 60}}分</span></p>

  </div>

  <div v-if="communication_loading" style="text-align: center; margin-top: 20px;">
    <p>通信中...</p>
  </div>
</template>

<style scoped>
#trackpage{
  height: calc(100dvh - 80px - 34px - 40px - 10px);
  overflow-y: scroll;
}

.calendar {
  height: calc(100dvh - 220px);
  overflow-y: auto;   /* 縦スクロールを有効に */
}

</style>
