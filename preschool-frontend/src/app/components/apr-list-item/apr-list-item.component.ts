import { HttpErrorResponse } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { MatSnackBar, MatSnackBarConfig } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { Company } from 'src/app/models/aprAccount.model';
import { AprService } from 'src/app/services/apr.service';
import { InputsService } from 'src/app/services/inputs.service';

@Component({
  selector: 'app-apr-list-item',
  templateUrl: './apr-list-item.component.html',
  styleUrls: ['./apr-list-item.component.css']
})
export class AprListItemComponent implements OnInit {
  

  constructor(private aprService: AprService,
              private inputsService: InputsService,
              private router: Router,
              private matSnackBar: MatSnackBar) { }

  @Input() apr: Company = new Company()

  ngOnInit(): void {
  }

  redirectToChangeCompany(): void{
    this.inputsService.setCompany(this.apr)
    this.router.navigate(['/ChangeCompanyDetails'])
  }

  liquidateCompany(): void{
    this.aprService.LiquidateCompany(this.apr.companyID)
    .subscribe({
      next: (response: string) => {
        this.openSnackBar(response, 1500);
      },
      error: (error: HttpErrorResponse) => {
        console.log("error happened: " + error.error)
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
