import { getLocaleNumberSymbol } from '@angular/common';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-welcome',
  templateUrl: './welcome.component.html',
  styleUrls: ['./welcome.component.css']
})
export class WelcomeComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

  isLoggedIn(): boolean {
    let token = localStorage.getItem("authToken")
    if(token && token != "") {
      return true
    }

    return false
  }

}
