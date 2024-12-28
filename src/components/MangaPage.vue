<script setup lang="ts">
import {API_HOST} from "@/utils/api.ts";

defineProps({
  id: {
    type: String,
    required: true
  },
  chapter: {
    type: Number,
    required: true
  },
  page: {
    type: Number,
    required: true
  }
})
</script>

<template>
  <img :src="`${API_HOST}/manga/${id}/${chapter}/${page}`" loading="lazy" alt=""/>
</template>

<style scoped>

</style>