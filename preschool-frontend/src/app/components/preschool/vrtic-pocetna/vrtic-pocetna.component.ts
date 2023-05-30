import { Component, OnInit } from '@angular/core';
import { Vrtic } from 'src/app/models/vrtic';
import { VrticService } from 'src/app/services/vrtic.service';

@Component({
  selector: 'app-vrtic-pocetna',
  templateUrl: './vrtic-pocetna.component.html',
  styleUrls: ['./vrtic-pocetna.component.css']
})

export class VrticPocetnaComponent implements OnInit {

  vrtici: Array<Vrtic> = [];

  constructor(private vrticService: VrticService) { }

  ngOnInit(): void {
    this.vrticService.GetAllVrtici()
      .subscribe({
        next: (data) => {
          this.vrtici = data;
        },
        error: (error) => {
          console.log(error)
        }
      })
  }

  isLoggedIn(): boolean {
    if (localStorage.getItem("authToken") != null) {
      return true;
    }
    else {
      return false;
    }
  }

  isAdmin(): boolean {
    if (localStorage.getItem("customRole") == "Admin") {
      return true;
    }
    else {
      return false;
    }
  }

  notLoggedIn(): boolean {
    if (localStorage.getItem("authToken") === null) {
      return true
    }
    else {
      return false
    }
  }
}
