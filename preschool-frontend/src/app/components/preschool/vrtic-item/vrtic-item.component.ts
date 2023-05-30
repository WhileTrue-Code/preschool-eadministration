import { Vrtic } from 'src/app/models/vrtic';
import { Component, Input, OnInit } from '@angular/core';


@Component({
  selector: 'app-vrtic-item',
  templateUrl: './vrtic-item.component.html',
  styleUrls: ['./vrtic-item.component.css']
})
export class VrticItemComponent implements OnInit {

  @Input() vrtic: Vrtic = new Vrtic();


  constructor() { }

  ngOnInit(): void {
  }

}
