import { Component, OnInit } from '@angular/core';
import {
  FormControl,
  FormGroup,
  FormBuilder,
  Validators, AbstractControl,
} from "@angular/forms";
import {Credentials} from "../../models/credentials";
import {AuthService} from "../../services/auth.service";
import {compareSegments} from "@angular/compiler-cli/src/ngtsc/sourcemaps/src/segment_marker";
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

  constructor(
    private formBuilder: FormBuilder,
    private authService: AuthService,
    private router: Router
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
    console.log(JSON.stringify(credentials))
    console.log(credentials.userType + this.formGroup.get('repeatPassword')?.value)
    this.authService.Registration(credentials).subscribe(
      response => {
        console.log(response)
        this.router.navigate(["/Login"])
      }, error => {
        this.router.navigate(["/Login"])
      }
    )
  }

  get registerForm(): { [key: string]: AbstractControl } {
    return this.formGroup.controls;
  }

  submitted = false;
}
