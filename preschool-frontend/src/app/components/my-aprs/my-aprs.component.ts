import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { AprCompany } from 'src/app/models/aprAccount.model';
import { AprService } from 'src/app/services/apr.service';

@Component({
  selector: 'app-my-aprs',
  templateUrl: './my-aprs.component.html',
  styleUrls: ['./my-aprs.component.css']
})
export class MyAprsComponent implements OnInit {

  companies: AprCompany[] = []
  permissionErr: string = ""

  constructor(private aprService: AprService) { }

  ngOnInit(): void {
    

    if (!this.permission()) {
      this.permissionErr = "Nemate prava pristupa ovoj stranici."
      return
    }

    this.aprService.GetAprCompaniesByFounderID()
    .subscribe({
      next: (response: AprCompany[]) => {
        console.log(response)
        this.companies = response
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

  }

}
