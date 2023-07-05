import { animate, state, style, transition, trigger } from '@angular/animations';
import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar, MatSnackBarConfig } from '@angular/material/snack-bar';
import { ActivatedRoute, Router } from '@angular/router';
import { Employee } from 'src/app/models/employee.model';
import { CrosoService } from 'src/app/services/croso.service';
import { InputsService } from 'src/app/services/inputs.service';

@Component({
  selector: 'app-register-employee',
  templateUrl: './register-employee.component.html',
  styleUrls: ['./register-employee.component.css'],
  animations: [
    trigger('fadeInOut', [
      transition(':enter', [
        style({ opacity: 0 }),
        animate('500ms', style({ opacity: 1 }))
      ]),
      transition(':leave', [
        animate('500ms', style({ opacity: 0 }))
      ])
    ])
  ], 
})
export class RegisterEmployeeComponent implements OnInit {
  formGroup: FormGroup = new FormGroup({
    firstName: new FormControl(''),
    lastName: new FormControl(''), 
    address: new FormControl(''),
    jmbg: new FormControl(''),
    idCardNumber: new FormControl(''),
    passportNumber: new FormControl(''),
    netSalary: new FormControl(''),
    employmentStatus:  new FormControl(''),
    employmentDuration:  new FormControl(''),
  });

  companyID: string = "";
  constructor (private inputsService: InputsService,
               private crosoService: CrosoService,
               private formBuilder: FormBuilder,
               private router: Router,
               private snackBar: MatSnackBar) { }

  submitted = false;

  isPassportEnabled = 'y'
  isJmbgEnabled = 'y'
  isIdCardEnabled = 'y'
  employmentStatus = 'n'

  ngOnInit(): void {
    this.companyID = this.inputsService.getCompanyID()
    console.log(this.companyID)

    this.formGroup = this.formBuilder.group({
      firstName: ['', [Validators.required]],
      lastName: ['', [Validators.required]],
      address: ['', [Validators.required]],
      jmbg: ['', []],
      idCardNumber: ['', []],
      passportNumber: ['', []],
      netSalary: ['', [Validators.required, Validators.min(40000)]],
      employmentStatus: ['', [Validators.required]],
      employmentDuration: ['', []],
    });
  }

  get registerEmployeeForm(): { [key: string]: AbstractControl } {
    return this.formGroup.controls;
  }
  
  
  onSubmit() {
    this.submitted = true;

    if (this.formGroup.invalid) {
      return;
    }
    
    if (this.formGroup.get('jmbg')!.value == "" && this.formGroup.get('idCardNumber')!.value == ""){
      if(this.formGroup.get('passportNumber')!.value == ""){
        this.submitted = false
        return
      }
    }

    let employee: Employee = new Employee();

    employee.firstName = this.formGroup.get('firstName')!.value
    employee.lastName = this.formGroup.get('lastName')!.value
    employee.address = this.formGroup.get('address')!.value
    employee.employeeID = this.formGroup.get('jmbg')!.value
    employee.companyID = parseInt(this.companyID)
    employee.netSalary = parseInt(this.formGroup.get('netSalary')!.value) 
    employee.idCardNumber = this.formGroup.get('idCardNumber')!.value
    employee.passportNumber = this.formGroup.get('passportNumber')!.value
    employee.employmentStatus = this.formGroup.get('employmentStatus')!.value
    employee.employmentDuration = this.formGroup.get('employmentDuration')!.value
    if (this.formGroup.get('employmentDuration')!.value == "") {
      employee.employmentDuration = 0
    }
    console.log(employee)
    this.crosoService.RequestEmployeeRegistration(employee)
    .subscribe({next:(response) => {
      this.openSnackBar(response)
      setTimeout(() => {
        this.inputsService.setCompanyID(this.companyID)
        this.router.navigate(['/CompanyEmployees']);
      }, 800)
    },error: (error: HttpErrorResponse)=> {
      this.openSnackBar(error.error)
    }
  })
    

  }

  openSnackBar(msg: string){
    let config: MatSnackBarConfig = new MatSnackBarConfig()
    config.duration = 1000 
    this.snackBar.open(msg, "ok", config)
  }

  change(type: string, event: Event) {
    const value = (event.target as HTMLInputElement).value;
    if ((type == "jmbg" || type == "idCardNumber")) {
      if(value == ''){
        this.isPassportEnabled = 'y'
        return
      }
      this.isPassportEnabled = 'n'
    }else if (type == "passportNumber") {
      if(value == ''){
        this.isJmbgEnabled = 'y'
        this.isIdCardEnabled = 'y'
        return
      }
      
      this.isJmbgEnabled = 'n'
      this.isIdCardEnabled = 'n'
    }

    if (type == "employmentStatus"){
      if (value != "definite_contract"){
        this.employmentStatus = 'n'
        return
      }
      this.employmentStatus = 'y'
    }
  }
  // customReqIdValidator() {
  //   let that = this
  //   return function(control: FormControl) {
  //     const value = control.value.toString().length;
  //     that.formGroup
  //     if (value != expectedValue) {
  //       return { exactValue: true };
  //     }
  //     return null;
  //   };
  // }

}
