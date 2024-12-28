<script setup lang="ts">
const props = defineProps({
  page: {
    type: Number,
    required: true
  },
  maxPage: {
    type: Number,
    required: true
  }
})
</script>

<template>
  <p class="text-xs font-bold rounded-xl p-1 bg-neutral-300 opacity-30">{{page}}/{{maxPage}} ({{Math.floor(page/maxPage * 100)}}%)</p>
</template>

<style scoped>

</style>