import { Component, Input, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Competition } from 'src/app/models/competition.model';
import { Prijava } from 'src/app/models/prijava';
import { CompetitionService } from 'src/app/services/competition.service';

@Component({
  selector: 'app-prijava-main',
  templateUrl: './prijava-main.component.html',
  styleUrls: ['./prijava-main.component.css']
})
export class PrijavaMainComponent implements OnInit {

  @Input() competition: Competition = new Competition();
  prijave: Array<Prijava> = [];
  comp_id = String(this.route.snapshot.paramMap.get("id"))


  constructor(private competitionService: CompetitionService,
    private route: ActivatedRoute) { }

  ngOnInit(): void {
    console.log(this.comp_id)
    this.competitionService.GetApplyesForOneCompetition(this.comp_id)
      .subscribe({
        next: (data) => {
          this.prijave = data;
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
