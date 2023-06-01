import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { User } from 'src/app/models/user.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-add-person-registry',
  templateUrl: './add-person-registry.component.html',
  styleUrls: ['./add-person-registry.component.css']
})
export class AddPersonRegistryComponent implements OnInit {

  constructor(
    private healthcareService: HealthcareService,
    private formBuilder: FormBuilder,
    private router: Router) { }

  submitted = false;

  formGroup: FormGroup = new FormGroup({
    ime: new FormControl(''),
    prezime: new FormControl(''),
    ime_oca: new FormControl(''),
    jmbg_oca: new FormControl(''),
    ime_majke: new FormControl(''),
    jmbg_majke: new FormControl(''),
    datum_rodjenja: new FormControl(''),
    mesto_rodjenja: new FormControl(''),
    jmbg: new FormControl(''),
    pol: new FormControl(''),
    drzava: new FormControl(''),
  })

  ngOnInit(): void {
    this.formGroup = this.formBuilder.group({
      ime: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(30)]],
      prezime: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(30)]],
      ime_oca: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(30)]],
      jmbg_oca: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(30)]],
      ime_majke: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(30)]],
      jmbg_majke: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(30)]],
      datum_rodjenja: ['', [Validators.required]],
      mesto_rodjenja: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(30)]],
      jmbg: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(30)]],
      pol: ['', [Validators.required]],
      drzava: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(30)]],
    })
  }

  get group(): { [key: string]: AbstractControl } {
    return this.formGroup.controls;
  }

  onSubmit() {
    this.submitted = true;

    if (this.formGroup.invalid) {
      return;
    }

    let user = new User()

    let ime = this.formGroup.get("ime")?.value
    console.log(ime)

    let prezime = this.formGroup.get("prezime")?.value
    console.log(prezime)

    let ime_oca = this.formGroup.get("ime_oca")?.value
    console.log(ime_oca)

    let jmbg_oca = this.formGroup.get("jmbg_oca")?.value
    console.log(jmbg_oca)

    let ime_majke = this.formGroup.get("ime_majke")?.value
    console.log(ime_majke)

    let jmbg_majke = this.formGroup.get("jmbg_majke")?.value
    console.log(jmbg_majke)

    var datum: Date = new Date(this.formGroup.get('datum_rodjenja')?.value)
    let datum_rodjenja = Number(datum.getTime()) / 1000
    console.log(datum_rodjenja)

    let mesto_rodjenja = this.formGroup.get("mesto_rodjenja")?.value
    console.log(mesto_rodjenja)

    let jmbg = this.formGroup.get("jmbg")?.value
    console.log(jmbg)

    let pol = this.formGroup.get("pol")?.value
    console.log(pol)

    let drzava = this.formGroup.get("drzava")?.value
    console.log(drzava)
  }

  drzave = new Array("Srbija", "Austrija", "Hrvatska", "Bosna", "Makedonija", "Bugarska", "Rumunija", "Crna Gora")
  polovi = new Array("Muski", "Zenski")

}
