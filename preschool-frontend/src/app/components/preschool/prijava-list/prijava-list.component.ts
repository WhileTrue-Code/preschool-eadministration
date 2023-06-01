import { Component, Input, OnInit } from '@angular/core';
import { Prijava } from 'src/app/models/prijava';

@Component({
  selector: 'app-prijava-list',
  templateUrl: './prijava-list.component.html',
  styleUrls: ['./prijava-list.component.css']
})
export class PrijavaListComponent implements OnInit {

  @Input() prijave: Prijava[] = [];


  constructor() { }

  ngOnInit(): void {
  }

}
