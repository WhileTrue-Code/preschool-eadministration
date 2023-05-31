import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, FormBuilder, Validators, AbstractControl } from '@angular/forms';
import { Appointment } from 'src/app/models/appointment.model';
import { User } from 'src/app/models/user.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-appointments',
  templateUrl: './appointments.component.html',
  styleUrls: ['./appointments.component.css']
})
export class AppointmentsDoctorComponent implements OnInit {

  appointments: Array<Appointment> = []
  user: User = new User()
  options = ["Slobodni", "Zauzeti", "Svi"]


  constructor(private healthcareService: HealthcareService) { }

  ngOnInit(): void {
    this.healthcareService.GetMyAppointmentsDoctor()
      .subscribe({
        next: (data) => {
          this.appointments = data;
        },
        error: (error) => {
          console.log(error)
        }
      })

    this.healthcareService.GetMe()
      .subscribe({
        next: (data) => {
          this.user = data;
        },
        error: (error) => {
          console.log(error)
        }
      })
  }

  search(search_option: string) {

    if (search_option == "Slobodni") {
      this.healthcareService.GetMyAvailableAppointmentsDoctor()
        .subscribe({
          next: (data) => {
            this.appointments = data;
          },
          error: (error) => {
            console.log(error)
          }
        })
    }

    if (search_option == "Zauzeti") {
      this.healthcareService.GetMyTakenAppointmentsDoctor()
        .subscribe({
          next: (data) => {
            this.appointments = data;
          },
          error: (error) => {
            console.log(error)
          }
        })
    }

    if (search_option == "Svi") {
      this.healthcareService.GetMyAppointmentsDoctor()
        .subscribe({
          next: (data) => {
            this.appointments = data;
          },
          error: (error) => {
            console.log(error)
          }
        })
    }

  }

}
