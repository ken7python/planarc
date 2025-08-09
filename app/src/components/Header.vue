<script setup lang="ts">
  import { useRoute } from 'vue-router'
  import homeIcon from '@/assets/icons/home.svg';
  import logoutIcon from '@/assets/icons/logout.svg';

  import { CONST } from '@/logic/const.js';
  import { user } from '@/logic/user.js';

  const route = useRoute()


</script>

<template>
  <div id="header">
    <div class="title">
      <homeIcon class="icon"></homeIcon>
      <h2 class="link">HOME</h2>
      <span class="spacer"></span>
    </div>
    <logoutIcon class="icon logout" @click="user.logout"></logoutIcon>
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

</style>