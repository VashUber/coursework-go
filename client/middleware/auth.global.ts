export default defineNuxtRouteMiddleware((to, from) => {
  const cookies = useCookie("auth");

  console.log("meta", to.meta.auth);
  console.log("cookie", cookies.value);
});
