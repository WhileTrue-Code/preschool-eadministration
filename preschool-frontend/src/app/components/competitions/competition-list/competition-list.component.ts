import { Component, Input, OnInit } from '@angular/core';
import { Competition } from 'src/app/models/competition.model';

@Component({
  selector: 'app-competition-list',
  templateUrl: './competition-list.component.html',
  styleUrls: ['./competition-list.component.css']
})
export class CompetitionListComponent implements OnInit {

  @Input() competitions: Competition[] = [];

  searchQuery: string = ''

  search(): void {
    this.competitions = this.competitions.filter(item => {
      return item.vrtic.grad.toLowerCase().includes(this.searchQuery.toLowerCase()) ||
        item.vrtic.opstina.toLocaleLowerCase().includes(this.searchQuery.toLowerCase()) ||
        item.vrtic.naziv.toLocaleLowerCase().includes(this.searchQuery.toLowerCase())

    });
  }


  constructor() { }

  ngOnInit(): void {
  }

}
