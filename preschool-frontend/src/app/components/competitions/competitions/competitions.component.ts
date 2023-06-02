import { Component, OnInit } from '@angular/core';
import { Competition } from 'src/app/models/competition.model';
import { CompetitionService } from 'src/app/services/competition.service';

@Component({
  selector: 'app-competitions',
  templateUrl: './competitions.component.html',
  styleUrls: ['./competitions.component.css']
})
export class CompetitionsComponent implements OnInit {

  competitions: Array<Competition> = [];

  constructor(private competitionService: CompetitionService) { }

  ngOnInit(): void {
    this.competitionService.GetAllCompetitions()
      .subscribe({
        next: (data) => {
          this.competitions = data;
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

  notLoggedIn(): boolean {
    if (localStorage.getItem("authToken") === null) {
      return true
    }
    else {
      return false
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
}
