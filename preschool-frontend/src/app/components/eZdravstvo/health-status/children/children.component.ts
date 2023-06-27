import { Component, OnInit } from '@angular/core';
import { Dete } from 'src/app/models/dete';

@Component({
  selector: 'app-children',
  templateUrl: './children.component.html',
  styleUrls: ['./children.component.css']
})
export class ChildrenComponent implements OnInit {

  deca: Array<Dete> = []
  
  constructor() { }

  //Treba getovati svu decu pa dodati dugme za pregled koji izbacuje formu za dodavanje health status reporta

  ngOnInit(): void {
    let novoDete = new Dete();
    novoDete.ime = "Jovan"
    novoDete.prezime = "Stefanovski"
    novoDete.jmbg = "1232123123212"
    novoDete.adresa = "Skolsa 14"
    novoDete.opstina = "Plandiste"
    this.deca.push(novoDete)
  }

}
