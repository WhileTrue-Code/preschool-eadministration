import { Injectable } from '@angular/core';
import { Employee } from '../models/employee.model';
import { VirtualAction } from 'rxjs';

@Injectable({
    providedIn: 'root'
})
export class InputsService {
  private companyID: string = "";
  private employee: Employee = new Employee();

  setCompanyID(value: string) {
    this.companyID = value;
  }

  getCompanyID(): string {
    return this.companyID;
  }

  setEmployee(value: Employee){
    this.employee = value;
  }

  getEmployee(): Employee {
    return this.employee
  }
}