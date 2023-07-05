import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar, MatSnackBarConfig } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { Company } from 'src/app/models/aprAccount.model';
import { AprService } from 'src/app/services/apr.service';
import { InputsService } from 'src/app/services/inputs.service';

@Component({
  selector: 'app-change-company-details',
  templateUrl: './change-company-details.component.html',
  styleUrls: ['./change-company-details.component.css']
})
export class ChangeCompanyDetailsComponent implements OnInit {

  formGroup: FormGroup = new FormGroup({
    name: new FormControl(''),
    address: new FormControl(''),
  });

  submitted: boolean = false;

  company: Company = new Company()

  companyName: string = "";
  companyAddress: string = "";

  constructor(private formBuilder: FormBuilder,
              private aprService: AprService,
              private inputsService: InputsService,
              private router: Router,
              private matSnackBar: MatSnackBar) { }

  ngOnInit(): void {
    this.company = this.inputsService.getCompany();
    this.companyName = this.company.name;
    this.companyAddress = this.company.address;

    this.formGroup = this.formBuilder.group({
      name: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(20), Validators.pattern('[-_a-zA-Z0-9 ]*')]],
      address: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(20), Validators.pattern('[-_a-zA-Z0-9 ]*')]],
    });
  }

  get changeAprForm(): { [key: string]: AbstractControl } {
    return this.formGroup.controls;
  }

  onSubmit(){
    this.submitted = true;

    if (this.formGroup.invalid) {
      return;
    }

    this.company.name = this.companyName;
    this.company.address = this.companyAddress;

    this.aprService.UpdateCompany(this.company)
    .subscribe({
      next: (response: string) => {
        this.openSnackBar(response, 1500);
        this.router.navigate(['MyAprs']);
      },
      error: (error: HttpErrorResponse) => {
        console.log("error occured: " + error);
      }
    })
  }

  openSnackBar(msg: string, duration: number) {
    let config = new MatSnackBarConfig()
    let defaultDuration = 1500;
    if (duration <= 0) {
      duration = defaultDuration;
    }
    
    config.duration = duration;
    this.matSnackBar.open(msg, "ok", config)
  }

  onNameChange(event: any){
    if(event.target.value) {
      this.companyName = event.target.value
    }
    
  }

  onAddressChange(event: any){
    if(event.target.value) {
      this.companyAddress = event.target.value
    }
  }

}
