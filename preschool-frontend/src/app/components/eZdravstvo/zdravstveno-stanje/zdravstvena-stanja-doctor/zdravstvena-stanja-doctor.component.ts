import { Component, OnInit } from '@angular/core';
import { ZdravstvenoStanje } from 'src/app/models/zdravstvenoStanje.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-zdravstvena-stanja-doctor',
  templateUrl: './zdravstvena-stanja-doctor.component.html',
  styleUrls: ['./zdravstvena-stanja-doctor.component.css']
})
export class ZdravstvenaStanjaDoctorComponent implements OnInit {

  zdravstvenaStanja: Array<ZdravstvenoStanje> = []

  constructor(private healthcareService: HealthcareService) { }

  ngOnInit(): void {
    this.healthcareService.GetAllZdravstvenaStanja()
      .subscribe({
        next: (response) => {
          this.zdravstvenaStanja = response
        },
        error: (error) => {
          console.log(error)
        }
      })
  }

}
