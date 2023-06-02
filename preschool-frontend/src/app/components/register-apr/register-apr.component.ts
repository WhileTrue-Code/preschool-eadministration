import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { Company } from 'src/app/models/aprAccount.model';
import { AprService } from 'src/app/services/apr.service';

@Component({
  selector: 'app-register-apr',
  templateUrl: './register-apr.component.html',
  styleUrls: ['./register-apr.component.css']
})
export class RegisterAprComponent implements OnInit {

  formGroup: FormGroup = new FormGroup({
    name: new FormControl(''),
    address: new FormControl(''), 
    startCapital: new FormControl(''),
    authorizedPersonFn: new FormControl(''),
    authorizedPersonLn: new FormControl(''),
  });

  constructor(private aprService: AprService,
              private formBuilder: FormBuilder,
              private router: Router) { }

  submitted = false;

  ngOnInit(): void {
    this.formGroup = this.formBuilder.group({
      name: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(20), Validators.pattern('[-_a-zA-Z0-9 ]*')]],
      address: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(20), Validators.pattern('[-_a-zA-Z0-9 ]*')]],
      startCapital: ['', [Validators.required, Validators.min(100)]],
      authorizedPersonFn: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(30), Validators.pattern('[-_a-zA-Z]*')]],
      authorizedPersonLn: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(30), Validators.pattern('[-_a-zA-Z]*')]],
    });

  }

  get registerAprForm(): { [key: string]: AbstractControl } {
    return this.formGroup.controls;
  }

  onSubmit() {
    this.submitted = true;

    if (this.formGroup.invalid) {
      return;
    }

    let aprCompany: Company = new Company();

    aprCompany.name = this.formGroup.get("name")?.value
    aprCompany.address = this.formGroup.get("address")?.value
    aprCompany.startCapital = parseInt(this.formGroup.get("startCapital")?.value, 10)
    aprCompany.authorizedPersonFirstName = this.formGroup.get("authorizedPersonFn")?.value
    aprCompany.authorizedPersonLastName = this.formGroup.get("authorizedPersonLn")?.value
    // var that = this
    this.aprService.RegisterAprCompany(aprCompany)
      .subscribe({
        next: (response:string) => {
          alert(response)
          console.log(response)
          
          setTimeout(() => {
            this.router.navigate(['/MyAprs']);
          }, 1500)
          
        },
        error: (error: HttpErrorResponse) => {
          console.log(error)
          if (error.status == 400) {
          }else if (error.status == 302) {
          }
        }
      });
    
  }

}
