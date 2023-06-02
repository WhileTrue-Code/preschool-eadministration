import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { CompanyID } from 'src/app/models/companyID.model';
import { CrosoService } from 'src/app/services/croso.service';

@Component({
  selector: 'app-register-croso',
  templateUrl: './register-croso.component.html',
  styleUrls: ['./register-croso.component.css']
})
export class RegisterCrosoComponent implements OnInit {

  formGroup: FormGroup = new FormGroup({
    companyID: new FormControl(''),
  });

  constructor(private formBuilder: FormBuilder,
              private crosoService: CrosoService,
              private router: Router) { }

  submitted = false;

  get registerCrosoForm(): { [key: string]: AbstractControl } {
    return this.formGroup.controls;
  }

  ngOnInit(): void {
    this.formGroup = this.formBuilder.group({
      companyID: ['', [Validators.required, this.exactValueValidator(8), Validators.pattern('[-_0-9]*')]],
    });
  }

  onSubmit() {
    this.submitted = true;

    if (this.formGroup.invalid) {
      return;
    }

    let companyID: CompanyID = new CompanyID()

    companyID.companyID = this.formGroup.get("companyID")?.value

    this.crosoService.RegisterCrosoCompany(companyID)
    .subscribe({
      next: (response:string) => {
        alert(response)
        console.log(response)
        
        setTimeout(() => {
          this.router.navigate(['/MyCrosos']);
        }, 1500)
        
      },
      error: (error: HttpErrorResponse) => {
        console.log(error)
      }
    });


  }

  exactValueValidator(expectedValue: number) {
    return function(control: FormControl) {
      const value = control.value.toString().length;
  
      if (value != expectedValue) {
        return { exactValue: true };
      }
      return null;
    };
  }

}
