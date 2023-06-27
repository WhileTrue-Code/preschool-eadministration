import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Company } from 'src/app/models/aprAccount.model';
import { CrosoService } from 'src/app/services/croso.service';

@Component({
  selector: 'app-my-crosos',
  templateUrl: './my-crosos.component.html',
  styleUrls: ['./my-crosos.component.css']
})
export class MyCrososComponent implements OnInit {
  companies: Company[] = []
  permissionErr: string = ""
  constructor(private crosoService: CrosoService,
              private router: Router) { }
  

  ngOnInit(): void {
    this.crosoService.GetCrosoCompaniesByFounderID().
    subscribe({
      next: (response: Company[]) => {
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
    this.router.navigate(['/RegisterApr'])
  }

}
