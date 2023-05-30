import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Competition } from 'src/app/models/competition.model';
import { CompetitionService } from 'src/app/services/competition.service';

@Component({
  selector: 'app-competition-add',
  templateUrl: './competition-add.component.html',
  styleUrls: ['./competition-add.component.css']
})
export class CompetitionAddComponent implements OnInit {

  constructor(
    private fb: FormBuilder,
    private competitionService: CompetitionService,
    private router: Router,
    private route: ActivatedRoute) { }

    competitionForm: FormGroup = new FormGroup({
      datum_objave: new FormControl(''),
      pocetak_konkursa: new FormControl(''),
      kraj_konkursa: new FormControl(''),
      grad: new FormControl(''),
      opstina: new FormControl(''),
      uzrast: new FormControl(''),
      broj_dece: new FormControl(''),
    });
    submitted = false;


    onSubmit(){
      this.submitted = true;
      if (this.competitionForm.invalid) {
        return;
      }

      var DatumObjave: Date = new Date(this.competitionForm.get('datum_objave')?.value)
      var PocetakKonkursa: Date = new Date(this.competitionForm.get('pocetak_konkursa')?.value)
      var KrajKonkursa: Date = new Date(this.competitionForm.get('kraj_konkursa')?.value)

      let dodajCompetition: Competition = new Competition();
      dodajCompetition.datum_objave = Number(DatumObjave.getTime()) / 1000
      dodajCompetition.pocetak_konkursa = Number(PocetakKonkursa.getTime()) / 1000
      dodajCompetition.kraj_konkursa = Number(KrajKonkursa.getTime()) / 1000
      dodajCompetition.uzrast = this.competitionForm.get("uzrast")?.value;
      dodajCompetition.broj_dece = this.competitionForm.get("broj_dece")?.value;
  

      this.route.params.subscribe(params => {

        const vrtic_id = params['id']
        this.competitionService.AddCompetition(dodajCompetition, vrtic_id)
        .subscribe({
          next: (data) => {
            this.router.navigate(['/Competitions']);
          },
          error: (error) => {
            console.log(error);
          },
          complete: () => {
            this.router.navigate(['/Competitions'])
          }
        })

      })
  
    }
  

  ngOnInit(): void {
  }

}
