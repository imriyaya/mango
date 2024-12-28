<script setup lang="ts">
import {RouterView, useRoute, useRouter} from 'vue-router'
import CommonHeader from "@/components/CommonHeader.vue";
import {useViewingPage} from "@/stores/viewingPage.ts";
</script>

<template>
  <div class="px-3">
    <CommonHeader />
    <div class="mx-auto max-w-[1000px]">
      <RouterView />
    </div>
  </div>
</template>

<style>
@tailwind base;
@tailwind components;
@tailwind utilities;

body {
  @apply bg-neutral-200
}
</style>
