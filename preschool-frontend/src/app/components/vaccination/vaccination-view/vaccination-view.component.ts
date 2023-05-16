import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Vaccination } from 'src/app/models/vaccination.model';
import { VaccinationService } from 'src/app/services/vaccination.service';

@Component({
  selector: 'app-vaccination-view',
  templateUrl: './vaccination-view.component.html',
  styleUrls: ['./vaccination-view.component.css']
})
export class VaccinationViewComponent implements OnInit {

  constructor(private route: ActivatedRoute,
              private router: Router,
              private vaccinationService: VaccinationService) { }

  vaccination: Vaccination = new Vaccination()
  vaccination_id = String(this.route.snapshot.paramMap.get("id"))

  ngOnInit(): void {
    this.vaccinationService.GetSingleVaccination(this.vaccination_id)
      .subscribe({
        next: (data) => {
          this.vaccination = data
        },
        error: (error) => {
          console.log(error)
        }
      })
  }

  update() {
    this.vaccinationService.SetVaccination(this.vaccination_id)
      .subscribe({
        next: () => {
          console.log("PUT")
          this.router.navigate(['/Vaccinations-Regular'])
        },
        error: (error) => {
          console.log(error);
        }
      })
  }

  isTaken(): boolean {
    if (this.vaccination.user != null) {
      return true;
    } else {
      return false;
    }
  }

}
