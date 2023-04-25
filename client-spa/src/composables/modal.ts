import { ref } from "vue";

export type TBuyTicketProps = {
  ticket: number;
};

type TModals = {
  type: "BuyTicket";
  props: TBuyTicketProps;
} | null;

const currModal = ref<TModals>(null);
const setModal = (modal: TModals) => {
  currModal.value = modal;
};
const close = () => {
  setModal(null);
};

export const useModal = () => {
  return {
    currModal,
    close,
    setModal,
  };
};
