import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { AppComponent } from "../app.component";
import { environment } from "src/environments/environment";
import { Competition } from "../models/competition.model";

@Injectable({
    providedIn: 'root'
})
export class CompetitionService {
    private url = "preschool/competitions"
    constructor(private http: HttpClient) { }

    public GetAllCompetitions(): Observable<Competition[]>{
        return this.http.get<Competition[]>(`${environment.baseApiUrl}/${this.url}/all`);
    }

    public AddCompetition(competition: Competition): Observable<Competition> {
        return this.http.post<Competition>(`${environment.baseApiUrl}/${this.url}/add`, competition);
    }
}