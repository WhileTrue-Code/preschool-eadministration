import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { AppComponent } from "../app.component";
import { environment } from "src/environments/environment";
import { Competition } from "../models/competition.model";
import { Company } from "../models/aprAccount.model";

@Injectable({
    providedIn: 'root'
})
export class AprService {
    private url = "apr"
    constructor(private http: HttpClient) { }

    public RegisterAprCompany(aprCompany: Company): Observable<string> {
        return this.http.post<string>(`${environment.baseApiUrl}/${this.url}/register`, aprCompany);
    }

    public GetAprCompaniesByFounderID(): Observable<Company[]> {
        return this.http.get<Company[]>(`${environment.baseApiUrl}/${this.url}/`);
    }

    public UpdateCompany(company: Company): Observable<string> {
        return this.http.put<string>(`${environment.baseApiUrl}/${this.url}/`, company);
    }

    public LiquidateCompany(companyID: number): Observable<string> {
        return this.http.put<string>(`${environment.baseApiUrl}/${this.url}/liquidate/${companyID}`, null);
    }

}