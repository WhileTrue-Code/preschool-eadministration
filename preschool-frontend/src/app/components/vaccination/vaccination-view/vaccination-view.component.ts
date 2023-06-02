import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Vaccination } from 'src/app/models/vaccination.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-vaccination-view',
  templateUrl: './vaccination-view.component.html',
  styleUrls: ['./vaccination-view.component.css']
})
export class VaccinationViewComponent implements OnInit {

  constructor(private route: ActivatedRoute,
              private router: Router,
              private healthcareService: HealthcareService) { }

  vaccination: Vaccination = new Vaccination()
  vaccination_id = String(this.route.snapshot.paramMap.get("id"))

  ngOnInit(): void {
    this.healthcareService.GetSingleVaccination(this.vaccination_id)
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
    this.healthcareService.SetVaccination(this.vaccination_id)
      .subscribe({
        next: () => {
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
