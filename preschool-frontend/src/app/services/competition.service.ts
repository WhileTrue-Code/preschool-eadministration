import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { AppComponent } from "../app.component";
import { environment } from "src/environments/environment";
import { Competition } from "../models/competition.model";
import { Prijava } from "../models/prijava";

@Injectable({
    providedIn: 'root'
})
export class CompetitionService {
    private url = "preschool/competitions"
    constructor(private http: HttpClient) { }

    public GetAllCompetitions(): Observable<Competition[]> {
        return this.http.get<Competition[]>(`${environment.baseApiUrl}/${this.url}/all`);
    }

    public AddCompetition(competition: Competition, vrtic_id: string): Observable<Competition> {
        return this.http.post<Competition>(`${environment.baseApiUrl}/preschool/vrtic/${vrtic_id}/competitions/add`, competition);
    }

    public GetSingleCompetition(competition_id: string): Observable<Competition> {
        return this.http.get<Competition>(`${environment.baseApiUrl}/${this.url}/getById/` + competition_id);
    }

    public ApplyForCompetition(prijava: Prijava, competition_id: string): Observable<Prijava> {
        return this.http.post<Prijava>(`${environment.baseApiUrl}/preschool/competitions/${competition_id}/apply`, prijava);
    }

    // public PromeniStatus(competition_id: string, competition: Competition) {
    //     return this.http.put(`${environment.baseApiUrl}/preschool/competitions/${competition_id}/changeStatus`, competition)
    // }

    public updateStanjeCompetition(competition_id: string) {
        return this.http.put(`${environment.baseApiUrl}/preschool/competitions/${competition_id}`, {});
    }


    public GetApplyesForOneCompetition(competition_id: string): Observable<Prijava[]> {
        return this.http.get<Prijava[]>(`${environment.baseApiUrl}/${this.url}/${competition_id}/applyes`);
    }


}