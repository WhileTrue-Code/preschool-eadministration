import { Component, Input, OnInit } from '@angular/core';
import { ZdravstvenoStanje } from 'src/app/models/zdravstvenoStanje.model';

@Component({
  selector: 'app-zdravstveno-stanje-list',
  templateUrl: './zdravstveno-stanje-list.component.html',
  styleUrls: ['./zdravstveno-stanje-list.component.css']
})
export class ZdravstvenoStanjeListComponent implements OnInit {

  @Input() zdravstvenaStanja: ZdravstvenoStanje[] = []

  constructor() { }

  ngOnInit(): void {
  }

}
