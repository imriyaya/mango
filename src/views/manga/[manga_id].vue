<script setup lang="ts">
import {useRoute} from "vue-router";
import {objectIdRegex} from "@/utils/regex.ts";
import {ref, type Ref} from "vue";
import {Manga, type MangaListResponse} from "@/types.ts";
import {$fetch} from "ofetch";
import {API_HOST} from "@/utils/api.ts";
import MangaPage from "@/components/MangaPage.vue";
import ChapterList from "@/components/ChapterList.vue";

const mangaId = useRoute().params.manga_id as string
const isValid = objectIdRegex.test(mangaId)

let manga: Ref<Manga | null> = ref(null);

$fetch<Manga>(API_HOST + `/manga/${mangaId}`)
    .then(response => {
      manga.value = response
    })
</script>

<template>
  <div v-if="isValid && manga">
    <div class="p-5 rounded-xl columns-2 bg-neutral-100">
      <div>
        <MangaPage class="w-64 rounded-xl" :id="manga?._id as string" :chapter="1" :page="1" />
      </div>
      <div>
        <h1 class="text-xl"><span class="text-sm">タイトル:</span> {{ manga?.title }}</h1>
        <h1 class="text-xl"><span class="text-sm">チャプター数:</span> {{ manga?.chapters.length }}</h1>
        <ChapterList :manga="manga" />
        <button class="p-1 bg-blue-500 text-white font-bold rounded-xl">この漫画を読む</button>
      </div>
    </div>
  </div>
  <div v-else>
    <h1 class="text-2xl font-bold text-center">ページURLが間違っています。</h1>
  </div>
</template>

<style scoped>

</style>