import { HttpErrorResponse } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { MatSnackBar, MatSnackBarConfig } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { Employee } from 'src/app/models/employee.model';
import { CrosoService } from 'src/app/services/croso.service';
import { InputsService } from 'src/app/services/inputs.service';

@Component({
  selector: 'app-employee-list-item',
  templateUrl: './employee-list-item.component.html',
  styleUrls: ['./employee-list-item.component.css']
})
export class EmployeeListItemComponent implements OnInit {

  @Input() employee: Employee = new Employee()
  status: string = ""
  constructor(private crosoService: CrosoService,
              private inputsService: InputsService,
              private router: Router,
              private matSnackBar: MatSnackBar) { }

  ngOnInit(): void {
    // this.status = this.employee.GetRsEmploymentStatus()
    if (this.employee.employmentStatus === "definite_contract") {
      this.status = "na odredjeno"
    }else if (this.employee.employmentStatus === "indefinite_contract"){
      this.status = "na neodredjeno"
    }else if (this.employee.employmentStatus === "temporary_works"){
      this.status = "privremeno povremeni rad"
    }else{
      this.status = "nezaposlen/odjavljen"
    }
  }

  redirectToChangeEmploymentStatus(){
    this.inputsService.setEmployee(this.employee)
    this.router.navigate(['/ChangeEmploymentStatus'])
  }

  cancelEmployment(){
    this.crosoService.CancelEmployment(this.employee.id)
    .subscribe({
      next: (response: string) => {
        this.openSnackBar(response, 1000);
        this.inputsService.setCompanyID(this.employee.companyID.toString());
      },
      error: (error: HttpErrorResponse) => {
        this.openSnackBar(error.error, 1200)
      }
    })
  }

  openSnackBar(msg: string, duration: number) {
    let config = new MatSnackBarConfig()
    let defaultDuration = 1000;
    if (duration <= 0) {
      duration = defaultDuration;
    }
    
    config.duration = duration;
    this.matSnackBar.open(msg, "ok", config)
  }

}
