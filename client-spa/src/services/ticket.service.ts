import dayjs from "dayjs";
import { $http } from "~/libs/axios";
import { ITicket } from "types/ticket";

export const ticketService = {
  buyTicket: async (body: { ticket: number; club: number }) => {
    return $http.post("/api/ticket/buy-ticket", body);
  },
  getUserTicket: async () => {
    const data = (
      await $http.get<{
        ticket: ITicket;
      }>("/api/ticket/get-user-ticket")
    ).data;

    if (!data.ticket) {
      return null;
    }

    return {
      ...data.ticket,
      start_date: dayjs(data.ticket.start_date).format("DD.MM.YYYY"),
      expired_date: dayjs(data.ticket.expired_date).format("DD.MM.YYYY"),
    };
  },
};
