import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';
import { Vrtic } from 'src/app/models/vrtic';
import { VrticService } from 'src/app/services/vrtic.service';

@Component({
  selector: 'app-vrtic-add',
  templateUrl: './vrtic-add.component.html',
  styleUrls: ['./vrtic-add.component.css']
})
export class VrticAddComponent implements OnInit {

  constructor(
    private fb: FormBuilder,
    private vrticService: VrticService,
    private router: Router) { }

  vrticForm: FormGroup = new FormGroup({
    naziv: new FormControl(''),
    adresa: new FormControl(''),
    telefon: new FormControl(''),
    email: new FormControl(''),
    grad: new FormControl(''),
    opstina: new FormControl(''),
  });
  submitted = false;


  onSubmit() {
    this.submitted = true;
    if (this.vrticForm.invalid) {
      return;
    }


    let dodajVrtic: Vrtic = new Vrtic();
    dodajVrtic.naziv = this.vrticForm.get("naziv")?.value;
    dodajVrtic.adresa = this.vrticForm.get("adresa")?.value;
    dodajVrtic.telefon = this.vrticForm.get("telefon")?.value;
    dodajVrtic.email = this.vrticForm.get("email")?.value;
    dodajVrtic.grad = this.vrticForm.get("grad")?.value;
    dodajVrtic.opstina = this.vrticForm.get("opstina")?.value;


    this.vrticService.AddVrtic(dodajVrtic)
      .subscribe({
        next: (data) => {
          this.router.navigate(['/PocetnaVrtic']);
        },
        error: (error) => {
          console.log(error);
        },
        complete: () => {
          this.router.navigate(['/PocetnaVrtic'])
        }
      })


  }


  ngOnInit(): void {
  }

}