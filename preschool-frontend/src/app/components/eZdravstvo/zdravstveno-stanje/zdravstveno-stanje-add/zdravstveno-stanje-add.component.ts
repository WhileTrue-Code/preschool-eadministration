import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ZdravstvenoStanje } from 'src/app/models/zdravstvenoStanje.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-zdravstveno-stanje-add',
  templateUrl: './zdravstveno-stanje-add.component.html',
  styleUrls: ['./zdravstveno-stanje-add.component.css']
})
export class ZdravstvenoStanjeAddComponent implements OnInit {

  constructor(
    private healthcareService: HealthcareService,
    private formBuilder: FormBuilder,
    private router: Router
  ) { }

  submitted = false;
  stanjeAlreadyExists = false

  formGroup: FormGroup = new FormGroup({
    jmbg: new FormControl(''),
    zdravstveni_problemi: new FormControl(''),
    specificna_ishrana: new FormControl(''),
    dom_zdravlja_u_kom_je_karton: new FormControl(''),
    smetnje_u_razvoju: new FormControl(''),
    specificni_podaci: new FormControl(''),
  })

  ngOnInit(): void {
    this.formGroup = this.formBuilder.group({
      jmbg: ['', [Validators.required, Validators.minLength(1), Validators.maxLength(30)]],
      zdravstveni_problemi: ['', Validators.maxLength(30)],
      specificna_ishrana: ['', Validators.maxLength(30)],
      dom_zdravlja_u_kom_je_karton: ['', Validators.maxLength(30)],
      smetnje_u_razvoju: ['', Validators.maxLength(30)],
      specificni_podaci: ['', Validators.maxLength(30)]
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

    let zdravstvenoStanje = new ZdravstvenoStanje();
    zdravstvenoStanje.jmbg = this.formGroup.get("jmbg")?.value
    zdravstvenoStanje.zdravstveni_problemi = this.formGroup.get("zdravstveni_problemi")?.value
    zdravstvenoStanje.specificna_ishrana = this.formGroup.get("specificna_ishrana")?.value
    zdravstvenoStanje.dom_zdravlja_u_kom_je_karton = this.formGroup.get("dom_zdravlja_u_kom_je_karton")?.value
    zdravstvenoStanje.smetnje_u_razvoju = this.formGroup.get("smetnje_u_razvoju")?.value
    zdravstvenoStanje.specificni_podaci = this.formGroup.get("specificni_podaci")?.value

    console.log(zdravstvenoStanje)
    
    this.healthcareService.NewZdravstvenoStanje(zdravstvenoStanje)
      .subscribe({
        next: () => {
          this.router.navigate(['/Zdravstvena-Stanja-Doctor'])
        },
        error: (error) => {
          console.log(error)
          if (error.status == 202) {
            this.stanjeAlreadyExists = true
          }
        }
      })
  }

}
