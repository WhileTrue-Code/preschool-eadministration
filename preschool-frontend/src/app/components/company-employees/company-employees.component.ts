import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Employee } from 'src/app/models/employee.model';
import { CrosoService } from 'src/app/services/croso.service';
import { CompanyIDService } from 'src/app/services/inputs.service';

@Component({
  selector: 'app-company-employees',
  templateUrl: './company-employees.component.html',
  styleUrls: ['./company-employees.component.css']
})
export class CompanyEmployeesComponent implements OnInit {

  employees: Employee[] = []
  permissionErr: string = ""
  companyID: string = ""

  constructor(private companyIDservice: CompanyIDService,
              private crosoService: CrosoService,
              private router: Router) { }

  ngOnInit(): void {
    this.companyID = this.companyIDservice.getCompanyID()

    if (!this.permission()) {
      this.permissionErr = "Nemate prava pristupa ovoj stranici."
      return
    }

    this.crosoService.GetEmployeesByCompanyID(this.companyID)
    .subscribe({
      next: (response: Employee[]) => {
        console.log(response)
        this.employees = response
      },
      error: (error: HttpErrorResponse) => {
        console.log("error message " + error.message)
      }
    })
  }

  permission(): boolean {
    const token = localStorage.getItem("authToken");
    if (!token || token == "") {
      return false
    }

    return true
  }

  redirectToRegister() {
    this.router.navigate(['/RegisterEmployee'])
  }
}
