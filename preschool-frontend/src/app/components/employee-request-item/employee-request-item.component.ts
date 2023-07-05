import { HttpErrorResponse } from '@angular/common/http';
import { ResourceLoader } from '@angular/compiler';
import { Component, Input, OnInit } from '@angular/core';
import { MatSnackBar, MatSnackBarConfig } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { ResolveStatus } from 'src/app/dto/resolveStatus';
import { Employee } from 'src/app/models/employee.model';
import { CrosoService } from 'src/app/services/croso.service';

@Component({
  selector: 'app-employee-request-item',
  templateUrl: './employee-request-item.component.html',
  styleUrls: ['./employee-request-item.component.css']
})
export class EmployeeRequestItemComponent implements OnInit {

  @Input() employee: Employee = new Employee();
  status = "";

  constructor(private crosoService: CrosoService,
              private matSnackBar: MatSnackBar,
              private router: Router) { }

  ngOnInit(): void {
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

  patchStatus(status: number) {
    let resolveStatus = new ResolveStatus();
    resolveStatus.companyID = this.employee.companyID;
    resolveStatus.employeeID = this.employee.employeeID;
    resolveStatus.status = status;

    this.crosoService.ResolveRegisterStatus(resolveStatus)
    .subscribe({
      next: (response: string) => {
        this.openSnackBar(response, 1500);
      },
      error: (error: HttpErrorResponse) => {
        console.log('error occured')
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

 

}
