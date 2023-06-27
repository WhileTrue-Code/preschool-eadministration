import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AddAppointment } from 'src/app/dto/addAppointment';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-appointment-add',
  templateUrl: './appointment-add.component.html',
  styleUrls: ['./appointment-add.component.css']
})
export class AppointmentAddComponent implements OnInit {

  appointmentFormGroup: FormGroup = new FormGroup({
    startOfAppointment: new FormControl(''),
    endOfAppointment: new FormControl('')
  });

  constructor(private healthcareService: HealthcareService,
              private router: Router,
              private formBuilder: FormBuilder) 
              { }

  submitted = false;
  alreadyExists = false;

  ngOnInit(): void {
    this.appointmentFormGroup = this.formBuilder.group({
      startOfAppointment: ['', [Validators.required]],
      endOfAppointment: ['', [Validators.required]]
    });
  }

  get appointmentGroup(): { [key: string]: AbstractControl } {
    return this.appointmentFormGroup.controls;
  }

  removeError() {
    this.alreadyExists = false;
  }

  onSubmit() {
    this.submitted = true;

    if (this.appointmentFormGroup.invalid) {
      return;
    }

    let addAppointment: AddAppointment = new AddAppointment();

    var StartOfAppointment: Date = new Date(this.appointmentFormGroup.get('startOfAppointment')?.value)
    var EndOfAppointment: Date = new Date(this.appointmentFormGroup.get('endOfAppointment')?.value)

    addAppointment.startOfAppointment = Number(StartOfAppointment.getTime()) / 1000;
    addAppointment.endOfAppointment = Number(EndOfAppointment.getTime()) / 1000;

    this.healthcareService.AddAppointment(addAppointment)
      .subscribe({
        next: (data) => {
          this.router.navigate(['/Appointments-Doctor']);
        },
        error: (error) => {
          console.log(error);
          this.alreadyExists = true;
        }
      })

  }

}
