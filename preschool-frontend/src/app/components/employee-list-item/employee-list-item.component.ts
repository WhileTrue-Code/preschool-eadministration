import { Component, Input, OnInit } from '@angular/core';
import { Employee } from 'src/app/models/employee.model';

@Component({
  selector: 'app-employee-list-item',
  templateUrl: './employee-list-item.component.html',
  styleUrls: ['./employee-list-item.component.css']
})
export class EmployeeListItemComponent implements OnInit {

  @Input() employee: Employee = new Employee()
  status: string = ""
  constructor() { }

  ngOnInit(): void {
    console.log(this.employee)
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

}
