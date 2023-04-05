import { RouteRecordRaw, createRouter, createWebHistory } from "vue-router";

const routes: RouteRecordRaw[] = [];

export const router = createRouter({
  routes,
  history: createWebHistory(),
});
