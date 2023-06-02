import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Employee } from 'src/app/models/employee.model';
import { CrosoService } from 'src/app/services/croso.service';
import { CompanyIDService } from 'src/app/services/inputs.service';

@Component({
  selector: 'app-register-employee',
  templateUrl: './register-employee.component.html',
  styleUrls: ['./register-employee.component.css']
})
export class RegisterEmployeeComponent implements OnInit {
  formGroup: FormGroup = new FormGroup({
    firstName: new FormControl(''),
    lastName: new FormControl(''), 
    address: new FormControl(''),
    jmbg: new FormControl(''),
    idCardNumber: new FormControl(''),
    passportNumber: new FormControl(''),
    employmentStatus:  new FormControl(''),
    employmentDuration:  new FormControl(''),
  });

  companyID: string = "";
  constructor (private companyIDservice: CompanyIDService,
               private crosoService: CrosoService,
               private formBuilder: FormBuilder,
               private router: Router) { }

  submitted = false;

  ngOnInit(): void {
    this.companyID = this.companyIDservice.getCompanyID()
    console.log(this.companyID)

    this.formGroup = this.formBuilder.group({
      firstName: ['', [Validators.required]],
      lastName: ['', [Validators.required]],
      address: ['', [Validators.required]],
      jmbg: ['', []],
      idCardNumber: ['', []],
      passportNumber: ['', []],
      employmentStatus: ['', [Validators.required]],
      employmentDuration: ['', []],
    });
  }

  get registerEmployeeForm(): { [key: string]: AbstractControl } {
    return this.formGroup.controls;
  }
  
  
  onSubmit() {
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
      alert(response)
      setTimeout(() => {
        this.companyIDservice.setCompanyID(this.companyID)
        this.router.navigate(['/CompanyEmployees']);
      }, 800)
    },error: (errResp: HttpErrorResponse)=> {
      console.log(errResp)
    }
  })
    

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
