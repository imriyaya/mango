import { createRouter, createWebHistory } from 'vue-router'
import {objectIdRegex} from "@/utils/regex.ts";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: () => import('../views/index.vue'),
    },
    {
      path: '/manga/:manga_id',
      component: () => import('../views/manga/[manga_id].vue')
    },
    {
      path: '/manga/:manga_id/:chapter_id',
      component: () => import('../views/manga/chapter/[chapter_id].vue')
    }
  ],
})

export default router
