import { Component, Input, OnInit } from '@angular/core';
import { Company } from 'src/app/models/aprAccount.model';

@Component({
  selector: 'app-apr-list-item',
  templateUrl: './apr-list-item.component.html',
  styleUrls: ['./apr-list-item.component.css']
})
export class AprListItemComponent implements OnInit {
  

  constructor() { }

  @Input() apr: Company = new Company()

  ngOnInit(): void {
  }

}
