<script setup lang="ts">
import {useRoute} from "vue-router";
import {objectIdRegex} from "@/utils/regex.ts";
import {computed, onMounted, onUnmounted, ref, type Ref, toRaw} from "vue";
import {Manga, type MangaListResponse} from "@/types.ts";
import {$fetch} from "ofetch";
import {API_HOST} from "@/utils/api.ts";
import MangaPage from "@/components/MangaPage.vue";
import {useViewingPage} from "@/stores/viewingPage.ts";

const mangaId = useRoute().params.manga_id as string
const isValidMangaId = objectIdRegex.test(mangaId)

const chapterId = Number(useRoute().params.chapter_id)

const manga: Ref<Manga | null> = ref(null);
const chapter = computed(() => manga.value?.chapters.find(chapter => chapter.id == chapterId))
const page: Ref<number> = ref(1)
const pageElement = ref<HTMLDivElement>({} as HTMLDivElement)
const isDragging = ref<boolean>(false)
const dragStartPosition = ref<{x: number, y: number}>({x: 0, y: 0})

$fetch<Manga>(API_HOST + `/manga/${mangaId}`)
    .then(response => {
      manga.value = response
    })

function incrementPage() {
  if (page.value != chapter.value?.page) page.value++
}

function decrementPage() {
  if (page.value != 1) page.value--
}

function onMouseDown(event: any) {
  isDragging.value = true
  dragStartPosition.value = {x: event.offsetX, y: event.offsetY}
}

function onMouseMove(event: any) {
  if (isDragging.value) {
    const pos = dragStartPosition.value
    console.log(`x: ${event.offsetX - pos.x} y: ${pos.y - event.offsetY}`)
    pageElement.value.style.left = event.offsetX - pos.x + "px"
  }
}

function onMouseUp(event: any) {
  isDragging.value = false
}

onMounted(() => useViewingPage().switch(true))
onUnmounted(() => useViewingPage().switch(false))
</script>

<template>
  <div v-if="isValidMangaId && chapter && manga">
    <div class="h-screen flex left-10" ref="pageElement" @mousedown="onMouseDown($event)" @mouseup="onMouseUp($event)" @mousemove="onMouseMove($event)">
      <MangaPage class="fixed pointer-events-none top-1 -right-full object-contain" v-if="page != 1" :id="manga._id" :chapter="chapterId" :page="page - 1"/>
      <MangaPage class="top-1 pointer-events-none mx-auto object-contain " :id="manga._id" :chapter="chapterId" :page="page"/>
      <MangaPage class="fixed pointer-events-none top-1 -left-full object-contain" v-if="page != chapter.page" :id="manga._id" :chapter="chapterId" :page="page + 1"/>
    </div>
    <button class="bottom-5 right-1/3 fixed z-10 text-center bg-blue-500 opacity-55" @click="incrementPage">Next</button>
    <button class="bottom-5 right-2/3 fixed z-10 text-center bg-blue-500 opacity-55" @click="decrementPage">Previous</button>
  </div>
  <div v-else>
    <h1 class="text-2xl font-bold text-center">ページURLが間違っています。</h1>
  </div>
</template>

<style scoped>

</style>