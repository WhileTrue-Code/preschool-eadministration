import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Appointment } from 'src/app/models/appointment.mode';
import { AppointmentService } from 'src/app/services/appointment.service';

@Component({
  selector: 'app-appointment-view',
  templateUrl: './appointment-view.component.html',
  styleUrls: ['./appointment-view.component.css']
})
export class AppointmentViewComponent implements OnInit {

  constructor(private route: ActivatedRoute,
              private router: Router,
              private appointmentService: AppointmentService) { }

  appointment: Appointment = new Appointment();
  appointment_id = String(this.route.snapshot.paramMap.get("id"))

  ngOnInit(): void {
    this.appointmentService.GetSingleAppointment(this.appointment_id)
      .subscribe({
        next: (data) => {
          this.appointment = data;
        }
      })
  }

  update() {
    this.appointmentService.SetAppointment(this.appointment_id)
      .subscribe({
        next: () => {
          console.log("PUT")
          this.router.navigate(['/Appointments'])
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
