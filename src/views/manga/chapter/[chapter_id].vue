<script lang="ts" setup>
import {useRoute} from "vue-router";
import {objectIdRegex} from "@/utils/regex.ts";
import {computed, onMounted, onUnmounted, ref, type Ref} from "vue";
import {Chapter, Manga} from "@/types.ts";
import {$fetch} from "ofetch";
import {API_HOST} from "@/utils/api.ts";
import MangaPage from "@/components/MangaPage.vue";
import {useViewingPage} from "@/stores/viewingPage.ts";
import PageState from "@/components/PageState.vue";

const mangaId = useRoute().params.manga_id as string
const isValidMangaId = objectIdRegex.test(mangaId)

const chapterId = Number(useRoute().params.chapter_id)

const manga: Ref<Manga | null> = ref(null);
const chapter = computed<Chapter>(() => manga.value?.chapters.find(chapter => chapter.id == chapterId))
const page: Ref<number> = ref(1)
const pageElement: Ref<HTMLDivElement> = ref<HTMLDivElement>({} as HTMLDivElement)

function incrementPage() {
  if (page.value != chapter.value?.page) page.value++
}

function decrementPage() {
  if (page.value != 1) page.value--
}

function onClick(event: any) {
  const x = event.offsetX
  if (x < pageElement.value.clientWidth / 2) {
    incrementPage()
  } else {
    decrementPage()
  }
}

$fetch<Manga>(API_HOST + `/manga/${mangaId}`)
    .then(response => {
      manga.value = response
    })

onMounted(() => useViewingPage().switch(true))
onUnmounted(() => useViewingPage().switch(false))
</script>

<template>
  <div v-if="isValidMangaId && chapter && manga">
    <div ref="pageElement" class="h-screen flex left-10" @click="onClick($event)">
      <MangaPage v-if="page != 1" loading="eager" :id="manga._id" :chapter="chapterId" :page="page - 1" class="fixed -z-10 top-0 -right-full"/>
      <MangaPage :id="manga._id" :chapter="chapterId" :page="page" class="top-1 pointer-events-none mx-auto object-contain "/>
      <MangaPage v-if="page != chapter.page" loading="eager" :id="manga._id" :chapter="chapterId" :page="page + 1" class="fixed -z-10 top-0 -right-full"/>
    </div>
    <PageState class="fixed z-10 bottom-5 right-1/2" :page="page" :max-page="chapter.page" />
  </div>
  <div v-else>
    <h1 class="text-2xl font-bold text-center">ページURLが間違っています。</h1>
  </div>
</template>

<style scoped>

</style>