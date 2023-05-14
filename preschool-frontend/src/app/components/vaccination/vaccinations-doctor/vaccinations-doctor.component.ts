import { Component, OnInit } from '@angular/core';
import { User } from 'src/app/models/user.model';
import { Vaccination } from 'src/app/models/vaccination.model';
import { VaccinationService } from 'src/app/services/vaccination.service';

@Component({
  selector: 'app-vaccinations-doctor',
  templateUrl: './vaccinations-doctor.component.html',
  styleUrls: ['./vaccinations-doctor.component.css']
})
export class VaccinationsDoctorComponent implements OnInit {

  vaccinations: Array<Vaccination> = [];
  user: User = new User();

  constructor(private vaccinationService: VaccinationService) { }

  ngOnInit(): void {
  }

}
