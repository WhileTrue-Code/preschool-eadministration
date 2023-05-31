import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup, Validators} from "@angular/forms";
import {Credentials} from "../../models/credentials";
import {AuthService} from "../../services/auth.service";
import { Router } from '@angular/router';
import {StoreServiceService} from "../../services/store-service.service";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  constructor(
    private authService: AuthService,
    private formBuilder: FormBuilder,
    private router: Router,
    private storeService: StoreServiceService
  ) { }

  credentials = new Credentials();
  notFound = false

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
    this.credentials.jmbg  = this.formGroup.get('jmbg')?.value
    this.credentials.password =  this.formGroup.get('password')?.value
    this.authService.Login(this.credentials).subscribe(
      ({
        next: (response) => {
          if (response != null){
            if (response == "JMBG not exist!"){
              localStorage.clear()
            }else if (response == "Password doesn't match!"){
              localStorage.clear()
            }else{
              localStorage.setItem('authToken', response)
              this.router.navigate(['/chose-service']).then();
            }
          }
        },
        error: (error) => {
          localStorage.clear()
          console.log(error.status)
          console.error(error)
          if (error.status = 403) {
            this.notFound = true;
          }
        }
      })
    );
  }
}
