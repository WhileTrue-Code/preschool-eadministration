import { Vrtic } from "./vrtic";

export class Competition {
    id: string = "";
    datum_objave: number = 0;
    pocetak_konkursa: number = 0;
    kraj_konkursa: number = 0;
    uzrast: string = "";
    broj_dece: number = 0;
    vrtic: Vrtic = new Vrtic;
    status: string = "";


    Competition(id: string, datum_objave: number, pocetak_konkursa: number, kraj_konkursa: number, uzrast:string, broj_dece:number,vrtic:Vrtic, status:string) {
        this.id = id;
        this.datum_objave = datum_objave;
        this.pocetak_konkursa = pocetak_konkursa;
        this.kraj_konkursa = kraj_konkursa;
        this.uzrast = uzrast;
        this.broj_dece = broj_dece;
        this.vrtic = vrtic;
        this.status = status;
    }
}