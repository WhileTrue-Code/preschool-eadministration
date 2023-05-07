import { Component, OnInit } from '@angular/core';
import { Appointment } from 'src/app/models/appointment.mode';
import { User } from 'src/app/models/user.model';
import { AppointmentService } from 'src/app/services/appointment.service';

@Component({
  selector: 'app-appointments',
  templateUrl: './appointments.component.html',
  styleUrls: ['./appointments.component.css']
})
export class AppointmentsComponent implements OnInit {

  appointments: Array<Appointment> = [];
  user: User = new User();

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
