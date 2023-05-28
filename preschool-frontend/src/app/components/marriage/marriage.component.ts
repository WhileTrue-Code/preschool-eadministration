import { Component, OnInit } from '@angular/core';
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {Marriage} from "../../models/marriage";

@Component({
  selector: 'app-marriage',
  templateUrl: './marriage.component.html',
  styleUrls: ['./marriage.component.css']
})
export class MarriageComponent implements OnInit {

  constructor() { }

  marriage: Marriage = new Marriage();

  formGroup: FormGroup = new FormGroup(
    {
      ime_mladozenje : new FormControl('', Validators.required),
      ime_mlade : new FormControl(''),
      prezime_mladozenje : new FormControl(''),
      devojkacko_prezime_mlade : new FormControl(''),
      datum_vencanja : new FormControl(''),
      mesto_vencanja : new FormControl(''),
      jmbg_mladozenje : new FormControl(''),
      jmbg_mlade : new FormControl(''),
      svedok_1 : new FormControl(''),
      svedok_2 : new FormControl(''),
    }
  )

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

  }

}
