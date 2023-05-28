import { Component, OnInit } from '@angular/core';
import {Credentials} from "../../models/credentials";
import {Router} from "@angular/router";

@Component({
  selector: 'app-choose-service',
  templateUrl: './choose-service.component.html',
  styleUrls: ['./choose-service.component.css']
})
export class ChooseServiceComponent implements OnInit {

  constructor(
    private router: Router,
  ) { }

  credentials = new Credentials();

  ngOnInit(): void {
  }

  selectService(service: string, credentials: Credentials){
    this.credentials.service = service;
    this.router.navigate(['/Login'], {state: {credentials}}).then()
  }

}
