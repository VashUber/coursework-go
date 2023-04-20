import { $http } from "~/libs/axios";
import type { ITicketPreview } from "types/ticket-preview";

export const ticketsPreviewService = {
  getAllTicketsPreview: async () => {
    return (
      await $http.get<{
        tickets: ITicketPreview[];
      }>("/api/tickets-preview/get-all")
    ).data.tickets;
  },
};
