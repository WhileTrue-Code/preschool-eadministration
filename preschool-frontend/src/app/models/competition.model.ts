export class Competition {
    id: number = 0;
    datum_objave: number = 0;
    pocetak_konkursa: number = 0;
    kraj_konkursa: number = 0;
    grad: string = "";
    opstina: string = "";
    uzrast: string = "";
    broj_dece: number = 0;

    Competition(id: number, datum_objave: number, pocetak_konkursa: number, kraj_konkursa: number, grad: string, opstina: string, uzrast:string, broj_dece:number) {
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