import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup, Validators} from "@angular/forms";
import {Credentials} from "../../models/credentials";
import {AuthService} from "../../services/auth.service";
import {MatSnackBar} from "@angular/material/snack-bar";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  constructor(
    private authService: AuthService,
    private formBuilder: FormBuilder
  ) { }

  formGroup: FormGroup = new FormGroup({
    jmbg: new FormControl(''),
    password: new FormControl(''),
    repeatPassword: new FormControl('')
  });

  ngOnInit(): void {
    this.formGroup = this.formBuilder.group({
      jmbg: ['', [Validators.required, Validators.minLength(3), Validators.maxLength(30), Validators.pattern('[-_a-zA-Z0-9]*')]],
      password: ['', [Validators.required, Validators.minLength(3), Validators.pattern('[-0-9]*')]]
    });
  }

  onSubmit() {
    const credentials = new Credentials();
    credentials._id = 0
    credentials.jmbg  = this.formGroup.get('jmbg')?.value
    credentials.password =  this.formGroup.get('password')?.value
    credentials.userType = ""
    this.authService.Login(credentials).subscribe(
      ({
        next: (response) => {
          if (response != null){
            localStorage.setItem('authToken', response)
            console.log(response)
          }
        },
        error: (error) => {
          localStorage.clear()
          console.log(error.status)
          console.error(error)
        }
      })
    );
  }


}
