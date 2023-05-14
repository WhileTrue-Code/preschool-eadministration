import { Component, Input, OnInit } from '@angular/core';
import { Appointment } from 'src/app/models/appointment.model';

@Component({
  selector: 'app-appointment-item',
  templateUrl: './appointment-item.component.html',
  styleUrls: ['./appointment-item.component.css']
})
export class AppointmentItemComponent implements OnInit {

  @Input() appointment: Appointment = new Appointment();

  constructor() { }

  ngOnInit(): void {
  }

  isTaken(): boolean {
    if (this.appointment.user != null) {
      return true;
    } else {
      return false;
    }
  }

}
