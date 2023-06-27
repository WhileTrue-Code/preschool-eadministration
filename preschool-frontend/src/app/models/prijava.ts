import { Dete } from "./dete";

export class Prijava {
    id: string = "";
    bodovi: number = 0;
    dete: Dete = new Dete;
    status: string = "";


    Prijava(id: string, bodovi: number, dete: Dete, status: string) {
        this.id = id;
        this.bodovi = bodovi
        this.dete = dete
        this.status = status
    }
}