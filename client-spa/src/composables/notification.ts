import { ref } from "vue";

export type TNotification = {
  id: number;
  message: string;
  success: boolean;
};

const notifications = ref<TNotification[]>([]);

const setNotification = (notification: Omit<TNotification, "id">) => {
  const id = Date.now();

  notifications.value.push({
    ...notification,
    id,
  });

  setTimeout(() => {
    notifications.value = notifications.value.filter((e) => e.id !== id);
  }, 3000);
};

export const useNotifications = () => {
  return { notifications, setNotification };
};
