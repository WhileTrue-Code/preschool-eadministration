import { Component, OnInit } from '@angular/core';
import { ZdravstvenoStanje } from 'src/app/models/zdravstvenoStanje.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-zdravstveno-stanje-view-my',
  templateUrl: './zdravstveno-stanje-view-my.component.html',
  styleUrls: ['./zdravstveno-stanje-view-my.component.css']
})
export class ZdravstvenoStanjeViewMyComponent implements OnInit {

  constructor(private healthcareService: HealthcareService) { }

  zdravstvenoStanje: ZdravstvenoStanje = new ZdravstvenoStanje();

  ngOnInit(): void {
    this.healthcareService.GetMyZdravstvenoStanje()
      .subscribe({
        next: (response) => {
          this.zdravstvenoStanje = response
        },
        error: (error) => {
          console.log(error)
        }
      })
  }

}
