import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Route } from '@angular/router';
import { Competition } from 'src/app/models/competition.model';
import { CompetitionService } from 'src/app/services/competition.service';

@Component({
  selector: 'app-competition-view',
  templateUrl: './competition-view.component.html',
  styleUrls: ['./competition-view.component.css']
})
export class CompetitionViewComponent implements OnInit {


  constructor(private route:ActivatedRoute, private competitionService:CompetitionService) { }

  competition_id = String(this.route.snapshot.paramMap.get("id"))
  competition:Competition = new Competition;

  

  ngOnInit(): void {
    console.log(this.competition_id)
    this.competitionService.GetSingleCompetition(this.competition_id)
    .subscribe({
      next:(data) => {
        this.competition=data
      }

    })
  }

}
