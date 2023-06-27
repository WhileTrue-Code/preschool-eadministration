import { Component, Input, OnInit } from '@angular/core';
import { Dete } from 'src/app/models/dete';

@Component({
  selector: 'app-child-item',
  templateUrl: './child-item.component.html',
  styleUrls: ['./child-item.component.css']
})
export class ChildItemComponent implements OnInit {

  @Input() dete: Dete = new Dete();

  constructor() { }

  ngOnInit(): void {
  }

}
