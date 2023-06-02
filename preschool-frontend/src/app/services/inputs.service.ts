import { Injectable } from '@angular/core';

@Injectable({
    providedIn: 'root'
})
export class CompanyIDService {
  private companyID: string = "";

  setCompanyID(value: string) {
    this.companyID = value;
  }

  getCompanyID(): string {
    return this.companyID;
  }
}