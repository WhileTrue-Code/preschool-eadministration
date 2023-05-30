export class Vrtic {
  id: string = '';
  naziv: string = '';
  adresa: string = '';
  telefon: string = '';
  email: string = '';
  grad: string = '';
  opstina: string = '';

  Vrtic(
    id: string,
    naziv: string,
    adresa: string,
    telefon: string,
    email: string,
    grad: string,
    opstina: string) 
  {
    this.id = id;
    this.naziv = naziv;
    this.adresa = adresa;
    this.telefon = telefon;
    this.email = email;
    this.grad = grad;
    this.opstina = opstina;
  }
}
