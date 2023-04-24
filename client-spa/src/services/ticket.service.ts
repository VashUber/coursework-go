import { $http } from "~/libs/axios";

export const ticketService = {
  buyTicket: async (body: { ticket: number; club: number }) => {
    return $http.post("/api/ticket/buy-ticket", body);
  },
  getUserTicket: async () => {
    return $http.get("/api/ticket/get-user-ticket");
  },
};
