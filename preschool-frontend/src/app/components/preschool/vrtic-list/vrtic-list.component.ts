import { Component, Input, OnInit } from '@angular/core';
import { Vrtic } from 'src/app/models/vrtic';

@Component({
  selector: 'app-vrtic-list',
  templateUrl: './vrtic-list.component.html',
  styleUrls: ['./vrtic-list.component.css']
})
export class VrticListComponent implements OnInit {

  @Input() vrtici: Vrtic[] = [];


  constructor() { }

  ngOnInit(): void {
  }

}
