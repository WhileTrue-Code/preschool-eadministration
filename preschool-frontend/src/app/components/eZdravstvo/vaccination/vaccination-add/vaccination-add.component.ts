 import { HttpResponse, HttpStatusCode } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AddVaccination } from 'src/app/dto/addVaccination';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-vaccination-add',
  templateUrl: './vaccination-add.component.html',
  styleUrls: ['./vaccination-add.component.css']
})
export class VaccinationAddComponent implements OnInit {

  vaccinationFormGroup: FormGroup = new FormGroup({
    startOfVaccination: new FormControl(''),
    endOfVaccination: new FormControl(''),
    vaccineType: new FormControl('')
  });

  constructor(private healthcareService: HealthcareService,
              private router: Router,
              private formBuilder: FormBuilder) 
              { }

  submitted = false
  alreadyExists = false
  vaccineTypes = ["BCG", "HB", "DTP", "IPV", "HIB", "PCV"]

  ngOnInit(): void {
    this.vaccinationFormGroup = this.formBuilder.group({
      startOfVaccination: ['', [Validators.required]],
      endOfVaccination: ['', [Validators.required]],
      vaccineType: ['', [Validators.required]]
    });
  }

  get vaccinationGroup(): { [key: string]: AbstractControl } {
    return this.vaccinationFormGroup.controls
  }

  removeError() {
    this.alreadyExists = false;
  }

  onSubmit() {
    this.submitted = true

    if (this.vaccinationFormGroup.invalid) {
      return;
    }

    let addVaccination: AddVaccination = new AddVaccination()

    var StartOfVaccination: Date = new Date(this.vaccinationFormGroup.get('startOfVaccination')?.value)
    var EndOfVaccination: Date = new Date(this.vaccinationFormGroup.get('endOfVaccination')?.value)
    var VaccineType = this.vaccinationFormGroup.get("vaccineType")?.value

    addVaccination.startOfVaccination = Number(StartOfVaccination.getTime()) / 1000
    addVaccination.endOfVaccination = Number(EndOfVaccination.getTime()) / 1000
    addVaccination.vaccineType = VaccineType

    this.healthcareService.AddVaccination(addVaccination)
      .subscribe({
        next: (data) => {
            this.router.navigate(['/Vaccinations-Doctor'])
        },
        error: (error) => {
          console.log(error)
          if (error.status = 406) {
            this.alreadyExists = true
          }
        }
      })
  }

}
