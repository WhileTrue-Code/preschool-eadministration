import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AddAppointment } from 'src/app/dto/addAppointment';
import { AppointmentService } from 'src/app/services/appointment.service';

@Component({
  selector: 'app-appointment-add',
  templateUrl: './appointment-add.component.html',
  styleUrls: ['./appointment-add.component.css']
})
export class AppointmentAddComponent implements OnInit {

  appointmentFormGroup: FormGroup = new FormGroup({
    dayOfAppointment: new FormControl(''),
    startOfAppointment: new FormControl(''),
    endOfAppointment: new FormControl('')
  });

  dateRange = new FormGroup({
    start: new FormControl(),
    end: new FormControl()
  });

  constructor(private appointmentService: AppointmentService,
              private router: Router,
              private formBuilder: FormBuilder
              ) { }

  submitted = false;

  ngOnInit(): void {
    this.appointmentFormGroup = this.formBuilder.group({
      dayOfAppointment: ['', [Validators.required]],
      startOfAppointment: ['', [Validators.required]],
      endOfAppointment: ['', [Validators.required]]
    });
  }

  get appointmentGroup(): { [key: string]: AbstractControl } {
    return this.appointmentFormGroup.controls;
  }

  onSubmit() {
    console.log("sent")
    // this.submitted = true;

    // if (this.appointmentFormGroup.invalid) {
    //   return;
    // }

    // let addAppointment: AddAppointment = new AddAppointment();

    // var DayOfAppointment: Date = new Date(this.appointmentFormGroup.get('dayOfAppointment')?.value)
    // var StartOfAppointment: Date = new Date(this.appointmentFormGroup.get('startOfAppointment')?.value)
    // var EndOfAppointment: Date = new Date(this.appointmentFormGroup.get('endOfAppointment')?.value)

    // addAppointment.dayOfAppointment = Number(DayOfAppointment.getTime()) / 1000
    // addAppointment.startOfAppointment = Number(StartOfAppointment.getTime()) / 1000
    // addAppointment.endOfAppointment = Number(EndOfAppointment.getTime()) / 1000

    // this.appointmentService.AddAppointment(addAppointment)
    //   .subscribe({
    //     next: (data) => {
    //       console.log("sent")
    //       this.router.navigate(['/Appointments']);
    //     },
    //     error: (error) => {
    //       console.log(error);
    //     }
    //   })

  }

}
