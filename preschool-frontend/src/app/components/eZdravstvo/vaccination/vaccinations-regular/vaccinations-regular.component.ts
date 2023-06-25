import { Component, OnInit } from '@angular/core';
import { Vaccination } from 'src/app/models/vaccination.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-vaccinations-regular',
  templateUrl: './vaccinations-regular.component.html',
  styleUrls: ['./vaccinations-regular.component.css']
})
export class VaccinationsRegularComponent implements OnInit {

  vaccinations: Array<Vaccination> = []

  constructor(private healthcareService: HealthcareService) { }

  ngOnInit(): void {
    this.healthcareService.GetAllAvailableVaccinations()
      .subscribe({
        next: (data) => {
          this.vaccinations = data
        }, 
        error: (error) => {
          console.log(error)
        }
      })
  }

}
