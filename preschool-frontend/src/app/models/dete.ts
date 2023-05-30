export class Dete {
    id: string = "";
    jmbg: string = "";
    datum_rodjenja: number = 0;
    ime: string = "";
    prezime: string = "";
    opstina: string = "";
    adresa: string = "";

    Dete(id: string, jmbg: string, datum_rodjenja: number, ime: string, prezime: string, opstina: string, adresa: string) {
        this.id = id;
        this.jmbg = jmbg;
        this.datum_rodjenja = datum_rodjenja;
        this.ime = ime;
        this.prezime = prezime;
        this.opstina = opstina;
        this.adresa = adresa;
    }
}