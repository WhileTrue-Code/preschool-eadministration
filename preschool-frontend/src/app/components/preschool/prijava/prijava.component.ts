import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Competition } from 'src/app/models/competition.model';
import { Prijava } from 'src/app/models/prijava';
import { CompetitionService } from 'src/app/services/competition.service';

@Component({
  selector: 'app-prijava',
  templateUrl: './prijava.component.html',
  styleUrls: ['./prijava.component.css']
})
export class PrijavaComponent implements OnInit {

  constructor(
    private fb: FormBuilder,
    private competitionService: CompetitionService,
    private router: Router,
    private route: ActivatedRoute) { }

    vrticForm: FormGroup = new FormGroup({
      jmbg: new FormControl(''),
      datum_rodjenja: new FormControl(''),
      ime: new FormControl(''),
      prezime: new FormControl(''),
      opstina: new FormControl(''),
      adresa: new FormControl(''),
    });
    submitted = false;


    onSubmit(){
      this.submitted = true;
      if (this.vrticForm.invalid) {
        return;
      }

      var DatumRodjenja: Date = new Date(this.vrticForm.get('datum_rodjenja')?.value)

      let dodajPrijavu: Prijava = new Prijava();
      dodajPrijavu.dete.jmbg = this.vrticForm.get("jmbg")?.value;
      dodajPrijavu.dete.datum_rodjenja = Number(DatumRodjenja.getTime()) / 1000
      dodajPrijavu.dete.ime = this.vrticForm.get("ime")?.value;
      dodajPrijavu.dete.prezime = this.vrticForm.get("prezime")?.value;
      dodajPrijavu.dete.opstina = this.vrticForm.get("opstina")?.value;
      dodajPrijavu.dete.adresa = this.vrticForm.get("adresa")?.value;

  

      this.route.params.subscribe(params => {

        const competition_id = params['id']
        this.competitionService.ApplyForCompetition(dodajPrijavu, competition_id)
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
