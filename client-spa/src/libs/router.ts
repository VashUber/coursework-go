import { RouteRecordRaw, createRouter, createWebHistory } from "vue-router";
import { useUser } from "~/composables/user";
import Default from "~/layouts/Default.vue";

const Index = () => import("~/views/Index.vue");
const Signin = () => import("~/views/Signin.vue");
const Signup = () => import("~/views/Signup.vue");
const Profile = () => import("~/views/Profile.vue");
const Ticket = () => import("~/views/Ticket.vue");

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
        meta: {
          guestOnly: true,
        },
      },
      {
        path: "/signup",
        name: "signup",
        component: Signup,
        meta: {
          guestOnly: true,
        },
      },
      {
        path: "/profile",
        name: "profile",
        component: Profile,
        meta: {
          userOnly: true,
        },
      },
      {
        path: "/ticket",
        name: "ticket",
        component: Ticket,
        meta: {
          userOnly: true,
        },
      },
    ],
  },
];

export const router = createRouter({
  routes,
  history: createWebHistory(),
});

const { user } = useUser();

router.beforeEach((to, from, next) => {
  if (user.value && to.meta.guestOnly) {
    return next("/");
  }

  if (!user.value && to.meta.userOnly) {
    return next("/");
  }

  return next();
});
