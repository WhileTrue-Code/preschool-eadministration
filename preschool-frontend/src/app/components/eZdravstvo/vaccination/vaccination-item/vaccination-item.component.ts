import { Component, Input, OnInit } from '@angular/core';
import { Vaccination } from 'src/app/models/vaccination.model';

@Component({
  selector: 'app-vaccination-item',
  templateUrl: './vaccination-item.component.html',
  styleUrls: ['./vaccination-item.component.css']
})
export class VaccinationItemComponent implements OnInit {

  @Input() vaccination: Vaccination = new Vaccination(); 

  constructor() { }

  ngOnInit(): void {
  }

  isTaken(): boolean {
    if (this.vaccination.user != null) {
      return true;
    } else {
      return false;
    }
  }

}
