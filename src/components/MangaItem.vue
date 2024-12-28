<script setup lang="ts">
import {Manga} from "@/types.ts";
import MangaPage from "@/components/MangaPage.vue";

const props = defineProps({
  manga: {
    type: Manga,
    required: true
  }
})
</script>

<template>
  <div class="w-24 lg:w-44 p-1 lg:p-3 rounded-xl cursor-pointer" @click="$router.push(`/manga/${props.manga._id}`);">
    <MangaPage class="w-28 lg:w-44 mx-auto rounded-xl" :id="manga._id" :chapter="1" :page="1"/>
    <p class="text-xs lg:text-sm mx-auto font-bold">{{manga.title}}</p>
  </div>
</template>

<style scoped>

</style>