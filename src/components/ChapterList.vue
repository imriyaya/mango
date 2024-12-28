<script setup lang="ts">
import {Manga} from "@/types.ts";

const props = defineProps({
  manga: {
    type: Manga,
    required: true
  }
})
</script>

<template>
  <div>
    <ol>
      <li v-for="chapter in manga.chapters" class="m-1 bg-white overflow-y-scroll">
        <RouterLink :to="`/manga/${manga._id}/${chapter.id}`" class="text-emerald-500 underline font-bold"><span class="text-xs">No{{chapter.id}}.</span> {{chapter.title}}</RouterLink>
      </li>
    </ol>
  </div>
</template>

<style scoped>

</style>