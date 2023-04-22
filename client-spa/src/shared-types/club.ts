export interface IClub {
  ID: number;
  thumb: string;
  image: string;
  info: string;
  name: string;
  address: {
    street: string;
    home: string;
  };
  schedule: string[];
}
