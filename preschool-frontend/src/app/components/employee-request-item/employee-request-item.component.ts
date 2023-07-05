import { HttpErrorResponse } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { MatSnackBar, MatSnackBarConfig } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
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

 

}
