import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Employee } from 'src/app/models/employee.model';
import { CrosoService } from 'src/app/services/croso.service';

@Component({
  selector: 'app-employee-requests',
  templateUrl: './employee-requests.component.html',
  styleUrls: ['./employee-requests.component.css']
})
export class EmployeeRequestsComponent implements OnInit {

  employees: Employee[] = new Array<Employee>;

  constructor(private crosoService: CrosoService) { }

  ngOnInit(): void {
    this.crosoService.GetPendingEmployees()
    .subscribe({
      next: (employees: Employee[]) => {
        console.log('employees' + employees)
        this.employees = employees;
      },
      error: (error: HttpErrorResponse) => {
        console.log('an error occured')
      }
    })

  }

}
