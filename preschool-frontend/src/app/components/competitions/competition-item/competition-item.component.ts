import { Component, Input, OnInit } from '@angular/core';
import { Competition } from 'src/app/models/competition.model';

@Component({
  selector: 'app-competition-item',
  templateUrl: './competition-item.component.html',
  styleUrls: ['./competition-item.component.css']
})
export class CompetitionItemComponent implements OnInit {

  @Input() competition: Competition = new Competition();

  constructor() { }

  ngOnInit(): void {
  }

}
