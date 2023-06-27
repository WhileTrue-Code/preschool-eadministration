import { Component, OnInit } from '@angular/core';
import { Appointment } from 'src/app/models/appointment.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-appointments-regular',
  templateUrl: './appointments-regular.component.html',
  styleUrls: ['./appointments-regular.component.css']
})
export class AppointmentsRegularComponent implements OnInit {

  appointments: Array<Appointment> = [];

  constructor(private healthcareService: HealthcareService) { }

  ngOnInit(): void {
    this.healthcareService.GetAllAvailableAppointments()
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
