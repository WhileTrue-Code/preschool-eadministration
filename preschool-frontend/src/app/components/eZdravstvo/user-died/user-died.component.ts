import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { UserDied } from 'src/app/models/userDied.model';
import { MarriageService } from 'src/app/services/marriage.service';

@Component({
  selector: 'app-user-died',
  templateUrl: './user-died.component.html',
  styleUrls: ['./user-died.component.css']
})
export class UserDiedComponent implements OnInit {

  constructor(
    private registratService: MarriageService,
    private formBuilder: FormBuilder,
    private router: Router
  ) { }

  submitted = false
  jmbgDoesNotExist = false

  formGroup: FormGroup = new FormGroup({
    jmbg: new FormControl(''),
    datumSmrti: new FormControl(''),
    mestoSmrti: new FormControl('')
  })



  ngOnInit(): void {
    this.formGroup = this.formBuilder.group({
      jmbg: ['', [Validators.required]],
      datumSmrti: ['', [Validators.required]],
      mestoSmrti: ['', [Validators.required]]
    })
  }

  get group(): { [key: string]: AbstractControl } {
    return this.formGroup.controls
  }

  onSubmit() {
    this.submitted = true

    if (this.formGroup.invalid) {
      return
    }

    let userDied = new UserDied()
    userDied.jmbg = this.formGroup.get("jmbg")?.value
    let DatumSmrti: Date = new Date(this.formGroup.get("datumSmrti")?.value)
    userDied.datumSmrti = Number(DatumSmrti.getTime()) / 1000
    userDied.mestoSmrti = this.formGroup.get("mestoSmrti")?.value

    this.registratService.UpdateCertificate(userDied)
      .subscribe({
        next: (data) => {
          this.router.navigate(["/Zdravstvena-Stanja-Doctor"])
          console.log("DATA")
          console.log(data)
        },
        error: (error) => {
          console.log(error)
          if (error.status = 500) {
            this.jmbgDoesNotExist = true
          }
        }
      })
  }

}
