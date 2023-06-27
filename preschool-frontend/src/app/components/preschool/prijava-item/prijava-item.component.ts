import { Component,Input, OnInit } from '@angular/core';
import { Prijava } from 'src/app/models/prijava';

@Component({
  selector: 'app-prijava-item',
  templateUrl: './prijava-item.component.html',
  styleUrls: ['./prijava-item.component.css']
})
export class PrijavaItemComponent implements OnInit {


  @Input() prijava: Prijava = new Prijava();
  
  
  constructor() { }

  ngOnInit(): void {
  }

  // isUpisan(): boolean {
  //   if (this.prijava.status = "Upisan") {
  //     return true;
  //   }
  //   else {
  //     return false;
  //   }
  // }
}
