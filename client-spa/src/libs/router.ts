import { RouteRecordRaw, createRouter, createWebHistory } from "vue-router";
import Default from "~/layouts/Default.vue";

const Index = () => import("~/views/Index.vue");

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    component: Default,
    children: [
      {
        path: "/",
        name: "home",
        component: Index,
      },
    ],
  },
];

export const router = createRouter({
  routes,
  history: createWebHistory(),
});
