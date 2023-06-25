import { Component, Input, OnInit } from '@angular/core';
import { Vaccination } from 'src/app/models/vaccination.model';

@Component({
  selector: 'app-vaccination-list',
  templateUrl: './vaccination-list.component.html',
  styleUrls: ['./vaccination-list.component.css']
})
export class VaccinationListComponent implements OnInit {

  @Input() vaccinations: Vaccination[] = [];

  constructor() { }

  ngOnInit(): void {
  }

}
