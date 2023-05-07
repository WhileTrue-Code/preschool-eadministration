import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { AppComponent } from "../app.component";
import { environment } from "src/environments/environment";
import { Competition } from "../models/competition.model";
import { AprCompany } from "../models/aprAccount.model";

@Injectable({
    providedIn: 'root'
})
export class AprService {
    private url = "apr"
    constructor(private http: HttpClient) { }

    public RegisterAprCompany(aprCompany: AprCompany): Observable<string> {
        return this.http.post<string>(`${environment.baseApiUrl}/${this.url}/register`, aprCompany);
    }

    public GetAprCompaniesByFounderID(): Observable<AprCompany[]> {
        return this.http.get<AprCompany[]>(`${environment.baseApiUrl}/${this.url}/`);
    }

}