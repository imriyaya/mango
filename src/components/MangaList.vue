<script setup lang="ts">
import {API_HOST} from "@/utils/api.ts";
import {$fetch} from "ofetch";
import type {MangaListResponse} from "@/types.ts";
import {ref, type Ref} from "vue";
import MangaItem from "@/components/MangaItem.vue";
import {Icon} from "@iconify/vue";

let mangaList: Ref<MangaListResponse | null> = ref(null);

$fetch<MangaListResponse>(API_HOST + "/manga/list?skip=0")
    .then(response => {
      mangaList.value = response
    })
</script>

<template>
  <div class="p-2 lg:p-5 bg-neutral-100 rounded-xl">
    <p class="text-xl lg:text-2xl">漫画リスト</p>
    <div v-if="mangaList" class="flex overflow-x-scroll">
      <MangaItem class="m-1 lg:m-3" v-for="manga in mangaList?.result" :manga="manga" />
    </div>
    <div v-else>
      <Icon class="w-8 h-8" icon="svg-spinners:ring-resize"/>
    </div>
  </div>
</template>

<style scoped>

</style>