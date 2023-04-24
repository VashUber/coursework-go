import { IClubSimple } from "./club";

export interface ITicket {
  id: number;
  price: number;
  start_date: string;
  expired_date: string;
  time: number;
  club: IClubSimple;
}
