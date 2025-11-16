import { GendersInterface } from "./IGender";

export interface UsersInterface {
  ID?: number;
  StudentID?: string;
  FirstName?: string;
  LastName?: string;
  Email?: string;
  Phone?: string;
  GenderID?: number;
  Gender?: GendersInterface;
  LinkedIn?: string;
  Profile?: string;
}
