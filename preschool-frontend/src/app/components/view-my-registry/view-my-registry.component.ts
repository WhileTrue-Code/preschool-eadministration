import { Component, OnInit } from '@angular/core';
import {StoreServiceService} from "../../services/store-service.service";
import {MarriageService} from "../../services/marriage.service";
import {MatSnackBar} from "@angular/material/snack-bar";
import {jsPDF} from "jspdf"
import html2canvas from 'html2canvas';
@Component({
  selector: 'app-view-my-registry',
  templateUrl: './view-my-registry.component.html',
  styleUrls: ['./view-my-registry.component.css']
})
export class ViewMyRegistryComponent implements OnInit {

  constructor(
    private storeService: StoreServiceService,
    private registrarService: MarriageService,
    private _snackBar: MatSnackBar
  ) { }

  response: any


  ngOnInit(): void {
    var convertedDateElement = document.getElementById('convertedDate');
    var timestamp = this.response.datum_smrti;
    var date = new Date(timestamp);
    var formattedDate = date.toLocaleDateString();
    if(convertedDateElement){
      convertedDateElement.textContent = formattedDate;
    }
  }

  GetCertificate(type: String) {
    this.registrarService.GetCertificate(type).subscribe(
      {
        next: (value) => {
          if(value==null){
            console.log(JSON.stringify(value))
            this.openSnackBar("Korisnik nije preminuo!", "OK")
          }else{
            this.response = value
            console.log(JSON.stringify(value))
          }
        }
      }
    )
  }

  openPDF(): void {
    let DATA: any = document.getElementById('htmlData');
    html2canvas(DATA).then((canvas) => {
      let fileWidth = 208;
      let fileHeight = (canvas.height * fileWidth) / canvas.width;
      const FILEURI = canvas.toDataURL('image/png');
      let PDF = new jsPDF('p', 'mm', 'a4');
      let position = 0;
      PDF.addImage(FILEURI, 'PNG', 0, position, fileWidth, fileHeight);
      PDF.save('Izvod.pdf');
    });
  }

  openSnackBar(message: string, action: string) {
    this._snackBar.open(message, action,  {
      duration: 3500,
      verticalPosition: "top",
    });
  }


}
