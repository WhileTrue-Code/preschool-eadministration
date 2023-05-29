import { Component, OnInit } from '@angular/core';
import {
  FormControl,
  FormGroup,
  FormBuilder,
  Validators, AbstractControl,
} from "@angular/forms";
import {Credentials} from "../../models/credentials";
import {AuthService} from "../../services/auth.service";
import { Router } from '@angular/router';
import {MatSnackBar} from "@angular/material/snack-bar";

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

  constructor(
    private formBuilder: FormBuilder,
    private authService: AuthService,
    private router: Router,
    private _snackBar: MatSnackBar,
  ) {
  }

  formGroup: FormGroup = new FormGroup({
    jmbg: new FormControl(''),
    password: new FormControl(''),
    repeatPassword: new FormControl('')
  });

  ngOnInit(): void {
    this.formGroup = this.formBuilder.group({
      jmbg: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(30), Validators.pattern('[-_a-zA-Z0-9]*')]],
      password: ['', [Validators.required, Validators.minLength(3), Validators.pattern('[-0-9]*')]],
      repeatPassword: ['', [Validators.required, Validators.minLength(3), Validators.pattern('[-0-9]*')]],
    });
  }

  onSubmit() {
    const credentials = new Credentials();
    credentials._id = 0
    credentials.jmbg  = this.formGroup.get('jmbg')?.value
    credentials.password =  this.formGroup.get('password')?.value
    credentials.userType = "Regular"
    if(this.formGroup.get('password')?.value == this.formGroup.get('repeatPassword')?.value){
      this.authService.Registration(credentials).subscribe(
        {
          next: (response) => {
            console.log(response)
            this.openSnackBar("Uspesno ste se registrovali", "OK")
            this.router.navigate(['Login']).then()
          },
          error: (error) => {
            console.log(JSON.stringify(error?.error?.text))
            this.openSnackBar(error?.error?.text, "OK")
          }
        }
      )
    }else {
      this.openSnackBar("Sifre se ne poklapaju", "OK")
    }
  }

  get registerForm(): { [key: string]: AbstractControl } {
    return this.formGroup.controls;
  }

  submitted = false;

  openSnackBar(message: string, action: string) {
    this._snackBar.open(message, action,  {
      duration: 3500,
      verticalPosition: "top",
    });
  }
}
