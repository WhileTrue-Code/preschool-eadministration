import { Component, Input, OnInit } from '@angular/core';
import { AprCompany } from 'src/app/models/aprAccount.model';

@Component({
  selector: 'app-apr-list-item',
  templateUrl: './apr-list-item.component.html',
  styleUrls: ['./apr-list-item.component.css']
})
export class AprListItemComponent implements OnInit {
  

  constructor() { }

  @Input() apr: AprCompany = new AprCompany()

  ngOnInit(): void {
  }

}
