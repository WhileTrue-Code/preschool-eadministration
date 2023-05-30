import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Vrtic } from 'src/app/models/vrtic';
import { VrticService } from 'src/app/services/vrtic.service';

@Component({
  selector: 'app-vrtic-view',
  templateUrl: './vrtic-view.component.html',
  styleUrls: ['./vrtic-view.component.css']
})
export class VrticViewComponent implements OnInit {

  constructor(private route:ActivatedRoute, private vrticService:VrticService) { }

  vrtic_id = String(this.route.snapshot.paramMap.get("id"))
  vrtic:Vrtic = new Vrtic;

  

  ngOnInit(): void {
    console.log(this.vrtic_id)
    this.vrticService.GetSingleVrtic(this.vrtic_id)
    .subscribe({
      next:(data) => {
        this.vrtic=data
      }

    })
  }

  isAdmin(): boolean {
    if (localStorage.getItem("customRole") == "Admin") {
      return true;
    }
    else {
      return false;
    }
  }

}
