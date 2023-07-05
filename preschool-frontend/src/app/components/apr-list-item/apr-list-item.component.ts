import { HttpErrorResponse } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { MatSnackBar, MatSnackBarConfig } from '@angular/material/snack-bar';
import { Company } from 'src/app/models/aprAccount.model';
import { AprService } from 'src/app/services/apr.service';

@Component({
  selector: 'app-apr-list-item',
  templateUrl: './apr-list-item.component.html',
  styleUrls: ['./apr-list-item.component.css']
})
export class AprListItemComponent implements OnInit {
  

  constructor(private aprService: AprService,
              private matSnackBar: MatSnackBar) { }

  @Input() apr: Company = new Company()

  ngOnInit(): void {
  }

  redirectToChangeCompany(): void{

  }

  liquidateCompany(): void{
    this.aprService.LiquidateCompany(this.apr.id)
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
