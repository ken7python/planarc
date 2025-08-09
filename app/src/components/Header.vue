<script setup lang="ts">
  import { useRoute } from 'vue-router'
  import homeIcon from '@/assets/icons/home.svg';
  import logoutIcon from '@/assets/icons/logout.svg';
  import backIcon from '@/assets/icons/back.svg';

  import { CONST } from '@/logic/const.js';
  import { user } from '@/logic/user.js';

  const route = useRoute()

  function back() {
    window.history.back();
  }


</script>

<template>
  <div id="header">
    <!-- HOME -->
    <div id="HOME" v-if="route.path === CONST.homeLink">
      <div class="title">
        <homeIcon class="icon"></homeIcon>
        <h2 class="link">HOME</h2>
        <span class="spacer"></span>
      </div>

      <logoutIcon class="icon logout" @click="user.logout"></logoutIcon>
    </div>

    <!-- AUTH -->
    <div id="AUTH" v-if="route.path === CONST.authLink">
      <div class="title">
        <homeIcon class="icon"></homeIcon>
        <h2 class="link">ユーザ登録・ログイン</h2>
        <span class="spacer"></span>
      </div>
    </div>

    <!-- LOGIN -->
    <div id="LOGIN" v-if="route.path === CONST.loginLink">
      <backIcon class="icon back" @click="back"></backIcon>
      <div class="title">
        <homeIcon class="icon"></homeIcon>
        <h2 class="link">ログイン</h2>
        <span class="spacer"></span>
      </div>
    </div>

    <!-- SIGNUP -->
    <div id="Signup" v-if="route.path === CONST.signupLink">
      <backIcon class="icon back" @click="back"></backIcon>
      <div class="title">
        <homeIcon class="icon"></homeIcon>
        <h2 class="link">ユーザ登録</h2>
        <span class="spacer"></span>
      </div>
    </div>

    <!-- DAILY -->
    <div id="DAILY" v-if="route.path === CONST.dailyLink">

      <div class="title">
        <homeIcon class="icon"></homeIcon>
        <h2 class="link">DAILY</h2>
        <span class="spacer"></span>
      </div>
    </div>

    <!-- FLAG -->
    <div id="FLAG" v-if="route.path === CONST.flagLink">
      <div class="title">
        <homeIcon class="icon"></homeIcon>
        <h2 class="link">未完了リスト</h2>
        <span class="spacer"></span>
      </div>
    </div>

  </div>
</template>

<style scoped>
  #header {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 80px;
    background-color: #1C409A;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .title {
    display: flex;
    align-items: center;
    gap: 8px;
    height: 80px;
    /* 横方向中央寄せは親のflexで制御 */
    padding-left: 0;
  }
  .icon {
    color: #FFFFFF;
    width: 35px;
    height: 35px;
    margin-top: 0;
    display: inline-block;
  }
  .link {
    color: #FFFFFF;
    font-size: 24px;
    font-weight: bold;
    margin: 0;
    line-height: 1;
    display: flex;
    align-items: center;
    height: 80px;
  }

  .logout {
    position: absolute;
    right: 20px;
    top: 20px;
    width: 40px;
    height: 40px;
    cursor: pointer;
  }

  .back {
    position: absolute;
    left: 20px;
    top: 25px;
    cursor: pointer;
    width: 30px;
    height: 30px;
  }

</style>