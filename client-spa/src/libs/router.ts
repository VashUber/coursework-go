import { RouteRecordRaw, createRouter, createWebHistory } from "vue-router";
import Default from "~/layouts/Default.vue";

const Index = () => import("~/views/Index.vue");
const Signin = () => import("~/views/Signin.vue");
const Signup = () => import("~/views/Signup.vue");

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
      {
        path: "/signin",
        name: "signin",
        component: Signin,
      },
      {
        path: "/signup",
        name: "signup",
        component: Signup,
      },
    ],
  },
];

export const router = createRouter({
  routes,
  history: createWebHistory(),
});
