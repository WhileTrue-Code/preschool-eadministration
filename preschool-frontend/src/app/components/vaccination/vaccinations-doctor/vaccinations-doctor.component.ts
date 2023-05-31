import { Component, OnInit } from '@angular/core';
import { User } from 'src/app/models/user.model';
import { Vaccination } from 'src/app/models/vaccination.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-vaccinations-doctor',
  templateUrl: './vaccinations-doctor.component.html',
  styleUrls: ['./vaccinations-doctor.component.css']
})
export class VaccinationsDoctorComponent implements OnInit {

  vaccinations: Array<Vaccination> = []
  user: User = new User()
  options = ["Slobodni", "Zauzeti", "Svi"]

  constructor(private healthcareService: HealthcareService) { }

  ngOnInit(): void {
    this.healthcareService.GetMyVaccinationsDoctor()
      .subscribe({
        next: (data) => {
          this.vaccinations = data
        },
        error: (error) => {
          console.log(error)
        }
      })
  }

  search(search_option: string) {
    
    if (search_option == "Slobodni") {
      this.healthcareService.GetMyAvailableVaccinationsDoctor()
        .subscribe({
          next: (data) => {
            this.vaccinations = data;
          },
          error: (error) => {
            console.log(error)
          }
        })
    }

    if (search_option == "Zauzeti") {
      this.healthcareService.GetMyTakenVaccinationsDoctor()
        .subscribe({
          next: (data) => {
            this.vaccinations = data;
          },
          error: (error) => {
            console.log(error)
          }
        })
    }

    if (search_option == "Svi") {
      this.healthcareService.GetMyVaccinationsDoctor()
        .subscribe({
          next: (data) => {
            this.vaccinations = data;
          },
          error: (error) => {
            console.log(error)
          }
        })
    }

  }

}
