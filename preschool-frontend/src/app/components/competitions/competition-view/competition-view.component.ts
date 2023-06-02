import { Component, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { ActivatedRoute, Route, Router } from '@angular/router';
import { Competition } from 'src/app/models/competition.model';
import { CompetitionService } from 'src/app/services/competition.service';

@Component({
  selector: 'app-competition-view',
  templateUrl: './competition-view.component.html',
  styleUrls: ['./competition-view.component.css']
})
export class CompetitionViewComponent implements OnInit {


  constructor(private router:Router,private route:ActivatedRoute, private competitionService:CompetitionService, private fb: FormBuilder) { }

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

  izmenaCompetitionForm = this.fb.group({
    status: ['']
  });

  deleteCompetition(competition_id: string){
    this.competitionService.DeleteCompetition(competition_id).subscribe((data) =>{
      console.log(data)
      
    })
    
  }

  isAdmin(): boolean {
    if (localStorage.getItem("customRole") == "Admin") {
      return true;
    }
    else {
      return false;
    }
  }

  isRegularUser(): boolean {
    if (localStorage.getItem("customRole") == "Regular") {
      return true;
    }
    else {
      return false;
    }
  }




  // changeStatus(){
  //   let izmeniCompetition: any = {};
  //   this.route.params.subscribe(params =>{
  //     const compID = params['id']
  //     izmeniCompetition.status = this.izmenaCompetitionForm.get("status")?.value;
  //     this.competitionService.PromeniStatus(compID, izmeniCompetition).subscribe(data =>{
  //       if(data){
  //         this.router.navigate(["/Competitions"])
  //       }
  //     })
  //   })

  // }

}
