import { Component, OnInit } from '@angular/core';
import { Vaccination } from 'src/app/models/vaccination.model';
import { HealthcareService } from 'src/app/services/healthcare.service';
import {jsPDF} from "jspdf"
import html2canvas from 'html2canvas'

@Component({
  selector: 'app-vaccinations-my-regular',
  templateUrl: './vaccinations-my-regular.component.html',
  styleUrls: ['./vaccinations-my-regular.component.css']
})
export class VaccinationsMyRegularComponent implements OnInit {

  constructor(private healthcareService: HealthcareService) { }
  vaccinations: Vaccination[] = []


  ngOnInit(): void {
    this.healthcareService.GetMyTakenVaccinationsRegular()
      .subscribe({
        next: (response) => {
          this.vaccinations = response
        }
      })
  }

  openPDF(): void {
    let DATA: any = document.getElementById('htmlData')
    html2canvas(DATA).then((canvas) => {
      let fileWidth = 208;
      let fileHeight = (canvas.height * fileWidth) / canvas.width
      const FILEURI = canvas.toDataURL('image/png')
      let PDF = new jsPDF('p', 'mm', 'a4')
      let position = 0
      PDF.addImage(FILEURI, 'PNG', 0, position, fileWidth, fileHeight)
      PDF.save('Vakcinacije.pdf')
    })
  }

}
