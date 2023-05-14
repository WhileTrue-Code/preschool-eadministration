import { Component, Input, OnInit } from '@angular/core';
import { Appointment } from 'src/app/models/appointment.model';

@Component({
  selector: 'app-appointment-list',
  templateUrl: './appointment-list.component.html',
  styleUrls: ['./appointment-list.component.css']
})
export class AppointmentListComponent implements OnInit {

  @Input() appointments: Appointment[] = [];

  constructor() { }

  ngOnInit(): void {
  }

}
