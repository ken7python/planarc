<script setup lang="ts">

</script>

<template>
  <div id="note">
    <textarea placeholder="メモ"></textarea>
    <br>
    <button class="btn" style="margin: 0 auto;">AIからのコメント</button>
    <br>
    <div id="comment">
      <p>AIからのコメントがここに表示されます。</p>
    </div>
  </div>
</template>

<style scoped>
  #note {
    text-align: center;
  }

  textarea {
    width: 86%;
    height: 22vh;
    background-color: #FFFFFF;
    padding: 10px;
    font-size: 16px;
    margin: 4px 0;
    border: 1px solid #E2E8F8;
    border-radius: 4px;
  }
</style>