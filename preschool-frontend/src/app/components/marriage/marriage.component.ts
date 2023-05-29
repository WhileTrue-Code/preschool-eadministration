import { Component, OnInit } from '@angular/core';
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {Marriage} from "../../models/marriage";
import {MarriageService} from "../../services/marriage.service";
import {Router} from "@angular/router";
import {MatSnackBar} from "@angular/material/snack-bar";

@Component({
  selector: 'app-marriage',
  templateUrl: './marriage.component.html',
  styleUrls: ['./marriage.component.css']
})
export class MarriageComponent implements OnInit {

  currentDate: string;
  marriage: Marriage = new Marriage();
  formGroup: FormGroup = new FormGroup(
    {
      ime_mladozenje : new FormControl('', [Validators.required]),
      ime_mlade : new FormControl('',[Validators.required]),
      prezime_mladozenje : new FormControl('',[Validators.required]),
      devojkacko_prezime_mlade : new FormControl('',[Validators.required]),
      datum_vencanja : new FormControl('',[Validators.required]),
      mesto_vencanja : new FormControl('',[Validators.required]),
      jmbg_mladozenje : new FormControl('',[Validators.required]),
      jmbg_mlade : new FormControl('',[Validators.required]),
      svedok_1 : new FormControl('',[Validators.required]),
      svedok_2 : new FormControl('',[Validators.required]),
    }
  )
  constructor(
    private marriageService: MarriageService,
    private router: Router,
    private _snackBar: MatSnackBar,
  ) {
    const today = new Date();
    this.currentDate = today.toISOString().split('T')[0];
  }



  ngOnInit(): void {
    this.formGroup.get('datum_vencanja')?.valueChanges.subscribe(
      (value) => {
        const timestamp = new Date(value).getTime()
        this.marriage.datum_vencanja = timestamp
      }
    )
  }

  createMarriage() {
    this.marriage.ime_mladozenje = this.formGroup.get('ime_mladozenje')?.value
    this.marriage.ime_mlade = this.formGroup.get('ime_mlade')?.value
    this.marriage.prezime_mladozenje = this.formGroup.get('prezime_mladozenje')?.value
    this.marriage.devojkacko_prezime_mlade = this.formGroup.get('devojkacko_prezime_mlade')?.value
    this.marriage.mesto_vencanja = this.formGroup.get('mesto_vencanja')?.value
    this.marriage.jmbg_mladozenje = this.formGroup.get('jmbg_mladozenje')?.value
    this.marriage.jmbg_mlade = this.formGroup.get('jmbg_mlade')?.value
    this.marriage.svedok_1.jmbg = this.formGroup.get('svedok_1')?.value
    this.marriage.svedok_2.jmbg = this.formGroup.get('svedok_2')?.value

    console.log(JSON.stringify(this.marriage))

    this.marriageService.CreateMarriage(this.marriage).subscribe({
      next: () => {
        this.openSnackBar("Uspesno ste kreirali vencanje", "OK")
      },
      error: (error) => {
        const errorMessage = error?.error?.text;
        this.openSnackBar(errorMessage, "OK")
      }
      }
    )

  }
  openSnackBar(message: string, action: string) {
    this._snackBar.open(message, action,  {
      duration: 3500,
      verticalPosition: "top",
    });
  }

}
