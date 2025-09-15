<script setup lang="ts">
  import { useRoute } from 'vue-router'
  import progressIcon from '@/assets/icons/progress-page.svg';
  import subjectIcon from '@/assets/icons/subject.svg';
  import analyticsIcon from '@/assets/icons/analytics.svg';
  import dailyIcon from '@/assets/icons/calendar_today.svg';
  import flagIcon from '@/assets/icons/flag.svg';

  import logoutIcon from '@/assets/icons/logout.svg';
  import backIcon from '@/assets/icons/back.svg';
  import accountIcon from '@/assets/icons/account_circle.svg';


  import { CONST } from '@/logic/const.js';
  import { user } from '@/logic/user.js';
  import Analysis from "../views/Analysis.vue";

  const route = useRoute()

  function back() {
    window.history.back();
  }


</script>

<template>
  <div id="header">
    <!-- HOME -->
    <div id="HOME" v-if="route.path === CONST.progressLink">
      <div class="title">
        <progressIcon class="icon"></progressIcon>
        <h2 class="link">進捗</h2>
        <span class="spacer"></span>
      </div>

      <logoutIcon class="icon logout" @click="user.logout"></logoutIcon>
    </div>

    <!-- SUBJECT -->
    <div id="FLAG" v-if="route.path === CONST.subjectLink">
      <div class="title">
        <subjectIcon class="icon"></subjectIcon>
        <h2 class="link">科目作成</h2>
        <span class="spacer"></span>
      </div>
    </div>

    <!-- DAILY -->
    <div id="DAILY" v-if="route.path === CONST.dailyLink">

      <div class="title">
        <dailyIcon class="icon"></dailyIcon>
        <h2 class="link">日</h2>
        <span class="spacer"></span>
      </div>
    </div>

    <!-- FLAG -->
    <div id="FLAG" v-if="route.path === CONST.flagLink">
      <div class="title">
        <flagIcon class="icon"></flagIcon>
        <h2 class="link">未完了リスト</h2>
        <span class="spacer"></span>
      </div>
    </div>

    <!-- analytics -->
    <div id="analytics" v-if="route.path === CONST.analysisLink">
      <div class="title">
        <analyticsIcon class="icon"></analyticsIcon>
        <h2 class="link">分析</h2>
        <span class="spacer"></span>
      </div>
    </div>

    <!-- AUTH -->
    <div id="AUTH" v-if="route.path === CONST.authLink">
      <div class="title">
        <accountIcon class="icon"></accountIcon>
        <h2 class="link">ユーザ登録・ログイン</h2>
        <span class="spacer"></span>
      </div>
    </div>

    <!-- LOGIN -->
    <div id="LOGIN" v-if="route.path === CONST.loginLink">
      <backIcon class="icon back" @click="back"></backIcon>
      <div class="title">
        <accountIcon class="icon"></accountIcon>
        <h2 class="link">ログイン</h2>
        <span class="spacer"></span>
      </div>
    </div>

    <!-- SIGNUP -->
    <div id="Signup" v-if="route.path === CONST.signupLink">
      <backIcon class="icon back" @click="back"></backIcon>
      <div class="title">
        <accountIcon class="icon"></accountIcon>
        <h2 class="link">ユーザ登録</h2>
        <span class="spacer"></span>
      </div>
    </div>

  </div>
</template>

<style scoped>
  #header {
    /*
    position: absolute;
    top: 0;
    left: 0;
    */
    width: 100%;
    height: 30px;
    background-color: #1C409A;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .title {
    display: flex;
    align-items: center;
    gap: 8px;
    height: 40px;
    /* 横方向中央寄せは親のflexで制御 */
    padding-left: 0;
  }
  .icon {
    color: #FFFFFF;
    width: 24px;
    height: auto;
    display: inline-block;
  }
  .link {
    color: #FFFFFF;
    font-size: 18px;
    font-weight: bold;
    margin: 0;
    line-height: 0;
    display: flex;
    align-items: center;
    height: 40px;
  }

  .logout {
    position: absolute;
    right: 20px;
    top: 5px;
    width: 20px;
    height: auto;
    cursor: pointer;
  }

  .back {
    position: absolute;
    left: 20px;
    top: 5px;
    cursor: pointer;
    width: 20px;
    height: 20px;
  }

</style>