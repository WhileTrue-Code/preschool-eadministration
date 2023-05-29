import {User} from "./user.model";

export class Marriage {
  _id = 0;
  ime_mladozenje = '';
  ime_mlade = '';
  prezime_mladozenje = '';
  devojkacko_prezime_mlade = 'devojkacko_prezime_mlade';
  datum_vencanja: number = 0;
  mesto_vencanja = '';
  jmbg_mladozenje = '';
  jmbg_mlade = '';
  svedok_1: User = new User()
  svedok_2: User = new User()
}
