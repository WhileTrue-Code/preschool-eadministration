import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, FormBuilder, Validators, AbstractControl } from '@angular/forms';
import { Appointment } from 'src/app/models/appointment.model';
import { User } from 'src/app/models/user.model';
import { AppointmentService } from 'src/app/services/appointment.service';

@Component({
  selector: 'app-appointments',
  templateUrl: './appointments.component.html',
  styleUrls: ['./appointments.component.css']
})
export class AppointmentsDoctorComponent implements OnInit {

  appointments: Array<Appointment> = [];
  user: User = new User();
  options = ["Slobodni", "Zauzeti", "Svi"]


  constructor(private appointmentService: AppointmentService) { }

  ngOnInit(): void {
    this.appointmentService.GetMyAppointmentsDoctor()
      .subscribe({
        next: (data) => {
          this.appointments = data;
        },
        error: (error) => {
          console.log(error)
        }
      })

    this.appointmentService.GetMe()
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

    console.log(search_option)

    if (search_option == "Slobodni") {
      this.appointmentService.GetMyAvailableAppointmentsDoctor()
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
      this.appointmentService.GetMyTakenAppointmentsDoctor()
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
      this.appointmentService.GetMyAppointmentsDoctor()
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

  isLoggedIn(): boolean {
    if (localStorage.getItem("authToken") != null) {
      return true;
    }
    else {
      return false;
    }
  }

  notLoggedIn(): boolean {
    if (localStorage.getItem("authToken") === null) {
      return true
    }
    else {
      return false
    }
  }

}
