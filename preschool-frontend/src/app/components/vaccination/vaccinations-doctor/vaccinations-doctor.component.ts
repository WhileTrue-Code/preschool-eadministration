import { Component, OnInit } from '@angular/core';
import { User } from 'src/app/models/user.model';
import { Vaccination } from 'src/app/models/vaccination.model';
import { VaccinationService } from 'src/app/services/vaccination.service';

@Component({
  selector: 'app-vaccinations-doctor',
  templateUrl: './vaccinations-doctor.component.html',
  styleUrls: ['./vaccinations-doctor.component.css']
})
export class VaccinationsDoctorComponent implements OnInit {

  vaccinations: Array<Vaccination> = []
  user: User = new User()
  options = ["Slobodni", "Zauzeti", "Svi"]

  constructor(private vaccinationService: VaccinationService) { }

  ngOnInit(): void {
    this.vaccinationService.GetMyVaccinationsDoctor()
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
      this.vaccinationService.GetMyAvailableVaccinationsDoctor()
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
      this.vaccinationService.GetMyTakenVaccinationsDoctor()
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
      this.vaccinationService.GetMyVaccinationsDoctor()
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
