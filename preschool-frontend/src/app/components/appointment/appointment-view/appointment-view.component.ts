import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Appointment } from 'src/app/models/appointment.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-appointment-view',
  templateUrl: './appointment-view.component.html',
  styleUrls: ['./appointment-view.component.css']
})
export class AppointmentViewComponent implements OnInit {

  constructor(private route: ActivatedRoute,
              private router: Router,
              private healthcareService: HealthcareService) { }

  appointment: Appointment = new Appointment();
  appointment_id = String(this.route.snapshot.paramMap.get("id"))

  ngOnInit(): void {
    this.healthcareService.GetSingleAppointment(this.appointment_id)
      .subscribe({
        next: (data) => {
          this.appointment = data;
        }
      })
  }

  update() {
    this.healthcareService.SetAppointment(this.appointment_id)
      .subscribe({
        next: () => {
          this.router.navigate(['/Appointments-Regular'])
        },
        error: (error) => {
          console.log(error);
        }
      })
  }

  isTaken(): boolean {
    if (this.appointment.user != null) {
      return true;
    } else {
      return false;
    }
  }

}
