import { Component, OnInit } from '@angular/core';
import { Appointment } from 'src/app/models/appointment.mode';
import { AppointmentService } from 'src/app/services/appointment.service';

@Component({
  selector: 'app-appointments-regular',
  templateUrl: './appointments-regular.component.html',
  styleUrls: ['./appointments-regular.component.css']
})
export class AppointmentsRegularComponent implements OnInit {

  appointments: Array<Appointment> = [];

  constructor(private appointmentService: AppointmentService) { }

  ngOnInit(): void {
    this.appointmentService.GetAllAvailableAppointments()
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
