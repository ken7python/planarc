<script setup lang="ts">
  import { ref } from 'vue';
  let subjectName = ref<string>('');
  let subjectColor = ref<string>('#000000');
</script>

<template>
  <div>
    <div id="addSubject">
      <input type="text" placeholder="科目名を入力" v-model="subjectName" />
      <br>

      <label class="color-field">
        <span class="color-label">色を選択：</span>
        <input type="color" value="#5B8DEF" aria-label="Select accent color" v-model="subjectColor" />
      </label>
      <br>

      <button class="btn" style="margin: 0 auto;">追加</button>
    </div>
    <div id="subjectList">
      <p>科目のリストを追加予定</p>
    </div>
  </div>
</template>

<style scoped>
  #addSubject {
    text-align: center;
  }
  #subjectList {
    text-align: center;
  }

  .color-field {
    display: inline-flex;
    align-items: center;
    gap: 8px;
    font-size: 14px;
  }

  .color-field input[type="color"] {
    width: 40px;
    height: 40px;
    border: none;
    padding: 0;
    border-radius: 6px;
    cursor: pointer;
    appearance: none;
  }

  .color-field input[type="color"]::-webkit-color-swatch {
    border: none;
    border-radius: 4px;
  }

  .color-field input[type="color"]::-moz-color-swatch {
    border: none;
    border-radius: 4px;
  }
</style>