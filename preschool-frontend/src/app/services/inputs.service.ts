import { Injectable } from '@angular/core';
import { Employee } from '../models/employee.model';
import { VirtualAction } from 'rxjs';
import { Company } from '../models/aprAccount.model';

@Injectable({
    providedIn: 'root'
})
export class InputsService {
  private companyID: string = "";
  private company: Company = new Company();
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

  setCompany(value: Company){
    this.company = value;
  }

  getCompany(): Company {
    return this.company
  }
}