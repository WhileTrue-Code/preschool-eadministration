import { Component, Inject, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Company } from 'src/app/models/aprAccount.model';
import { InputsService } from 'src/app/services/inputs.service';

@Component({
  selector: 'app-croso-list-item',
  templateUrl: './croso-list-item.component.html',
  styleUrls: ['./croso-list-item.component.css']
})
export class CrosoListItemComponent implements OnInit {
  
  @Input() croso: Company = new Company()
  
  constructor(private router: Router,
              private inputsService: InputsService) { }
  
  ngOnInit(): void {
  }

  redirectToRegisterEmployee() {
    this.inputsService.setCompanyID(this.croso.companyID.toString())
    this.router.navigate(['RegisterEmployee'])
  }

  redirectToCompanyEmployees() {
    this.inputsService.setCompanyID(this.croso.companyID.toString())
    this.router.navigate(['CompanyEmployees'])
  }

}
