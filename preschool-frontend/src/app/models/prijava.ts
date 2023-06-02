import { Dete } from "./dete";

export class Prijava {
    id: string = "";
    bodovi: number = 0;
    dete: Dete = new Dete;


    Prijava(id: string, bodovi: number, dete: Dete) {
        this.id = id;
        this.bodovi = bodovi
        this.dete = dete
    }
}