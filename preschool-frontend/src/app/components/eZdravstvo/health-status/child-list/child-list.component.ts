import { Component, Input, OnInit } from '@angular/core';
import { Dete } from 'src/app/models/dete';

@Component({
  selector: 'app-child-list',
  templateUrl: './child-list.component.html',
  styleUrls: ['./child-list.component.css']
})
export class ChildListComponent implements OnInit {

  @Input() deca: Dete[] = [];

  constructor() { }

  ngOnInit(): void {
  }

}
