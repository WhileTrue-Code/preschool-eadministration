export class Competition {
    id: number = 0;
    datum_objave: Date = new Date;
    pocetak_konkursa: Date = new Date;
    kraj_konkursa: Date = new Date;
    grad: string = "";
    opstina: string = "";
    uzrast: string = "";
    broj_dece: number = 0;

    Competition(id: number, datum_objave: Date, pocetak_konkursa: Date, kraj_konkursa: Date, grad: string, opstina: string, uzrast:string, broj_dece:number) {
        this.id = id;
        this.datum_objave = datum_objave;
        this.pocetak_konkursa = pocetak_konkursa;
        this.kraj_konkursa = kraj_konkursa;
        this.grad = grad;
        this.opstina = opstina;
        this.uzrast = uzrast;
        this.broj_dece = broj_dece;
    }
}