import { Dete } from "./dete";

export class Prijava {
    id: string = "";
    bodovi: string = "";
    dete: Dete = new Dete;


    Prijava(id: string, bodovi: string, dete: Dete) {
        this.id = id;
        this.bodovi = bodovi
        this.dete = dete
    }
}