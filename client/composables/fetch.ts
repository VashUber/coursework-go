import { appendHeader } from "h3";

export const $http = async (url: string) => {
  const event = useRequestEvent();
  const headerCookies = useRequestHeaders(["cookie"]).cookie || "";

  const res = await $fetch.raw(url, {
    credentials: "include",
    headers: {
      cookie: headerCookies,
    },
  });

  const cookies = (res.headers.get("set-cookie") || "").split(",");

  for (const cookie of cookies) {
    appendHeader(event, "set-cookie", cookie);
  }

  return res._data;
};
