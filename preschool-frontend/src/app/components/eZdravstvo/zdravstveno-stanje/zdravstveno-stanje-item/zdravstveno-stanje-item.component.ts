import { Component, Input, OnInit } from '@angular/core';
import { ZdravstvenoStanje } from 'src/app/models/zdravstvenoStanje.model';

@Component({
  selector: 'app-zdravstveno-stanje-item',
  templateUrl: './zdravstveno-stanje-item.component.html',
  styleUrls: ['./zdravstveno-stanje-item.component.css']
})
export class ZdravstvenoStanjeItemComponent implements OnInit {

  @Input() zdravstvenoStanje: ZdravstvenoStanje = new ZdravstvenoStanje();

  constructor() { }

  ngOnInit(): void {
  }

}
