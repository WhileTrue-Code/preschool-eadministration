import { Component, Input, OnInit } from '@angular/core';
import { Employee } from '../../models/employee.model';
import { MatSnackBar, MatSnackBarConfig } from '@angular/material/snack-bar';
import { CrosoService } from '../../services/croso.service';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { InputsService } from '../../services/inputs.service';
import { ChangeEmploymentStatus } from '../../dto/changeEmploymentStatus';
import { HttpErrorResponse, HttpHeaderResponse } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-employee-change-employment-status',
  templateUrl: './employee-change-employment-status.component.html',
  styleUrls: ['./employee-change-employment-status.component.css']
})
export class EmployeeChangeEmploymentStatusComponent implements OnInit {

  form: FormGroup = new FormGroup({
    employmentStatus: new FormControl(''),
    employmentDuration: new FormControl(''),
  });

  employee: Employee = new Employee();

  selectedEmploymentStatus = "";
  employmentDuration = 0;

  validationErr: string = "";
  
  constructor(private crosoService: CrosoService,
              private matSnackBar: MatSnackBar,
              private inputsService: InputsService,
              private router: Router,
              /*private formBuilder: FormBuilder*/) { }

  ngOnInit(): void {
    this.employee = this.inputsService.getEmployee();
    this.inputsService.setCompanyID(this.employee.companyID.toString())
    this.selectedEmploymentStatus = this.employee.employmentStatus;
    this.employmentDuration = this.employee.employmentDuration;
  }

  onSubmit() {
    if (this.selectedEmploymentStatus == "definite_contract" && this.employmentDuration == 0){
      this.validationErr = "Za izabrani tip ugovora je potrebno izabrati trajanje ugovora izrazen u broju meseci.";
      return
    }

    this.validationErr = ''

    let changeEmploymentStatus = new ChangeEmploymentStatus();

    changeEmploymentStatus.employmentStatus = this.selectedEmploymentStatus;
    changeEmploymentStatus.employmentDuration = this.employmentDuration;

    this.crosoService.ChangeEmploymentStatus(this.employee.id, changeEmploymentStatus)
    .subscribe({
      next: (response: string) => {
        setTimeout(() => {
          this.router.navigate(['/CompanyEmployees'])
          this.openSnackBar(response, 2000)
        }, 500)
        
        
      },
      error: (error: HttpErrorResponse) => {

      }
    })

  }

  onEmploymentStatusSelected(event: any) {
    let selected = event.value;
    if (selected) {
     
      this.selectedEmploymentStatus = selected;
      if (selected == "definite_contract" && this.employmentDuration == 0){
        this.validationErr = "Za izabrani tip ugovora je potrebno izabrati trajanje ugovora izrazen u broju meseci.";
      }else {
        this.employmentDuration = 0;
        this.validationErr = "";
      }
    }
  }

  onEmploymentDurationInput(event: any) {
    if(event.target.value){
      this.employmentDuration = parseInt(event.target.value)
      if (this.selectedEmploymentStatus == "definite_contract" && this.employmentDuration == 0){
        this.validationErr = "Za izabrani tip ugovora je potrebno izabrati trajanje ugovora izrazen u broju meseci.";
      }else {
        this.validationErr = "";
      }
    }
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

  

}
