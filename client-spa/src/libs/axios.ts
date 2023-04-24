import axios from "axios";
import { useNotifications } from "~/composables/notification";

const { setNotification } = useNotifications();

export const $http = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  withCredentials: true,
  timeout: 1000,
});

$http.interceptors.response.use(
  (res) => res,
  (err) => {
    setNotification({
      message: err.response.data ? err.response.data : `${err.response.status} error code`,
      success: false,
    });
  }
);
