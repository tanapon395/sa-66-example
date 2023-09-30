import { GendersInterface } from "./IGender";

export interface UsersInterface {
  ID?: number;
  FirstName?: string;
  LastName?: string;
  Email?: string;
  Phone?: string;
  GenderID?: number;
  Gender?: GendersInterface;
}
