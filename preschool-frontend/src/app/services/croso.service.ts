import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { AppComponent } from "../app.component";
import { environment } from "src/environments/environment";
import { Competition } from "../models/competition.model";
import { Company } from "../models/aprAccount.model";
import { CompanyID } from "../models/companyID.model";
import { Employee } from "../models/employee.model";

@Injectable({
    providedIn: 'root'
})
export class CrosoService {
    private url = "croso"
    constructor(private http: HttpClient) { }

    public RegisterCrosoCompany(companyID: CompanyID): Observable<string> {
        return this.http.post<string>(`${environment.baseApiUrl}/${this.url}/register`, companyID);
    }

    public GetCrosoCompaniesByFounderID(): Observable<Company[]> {
        return this.http.get<Company[]>(`${environment.baseApiUrl}/${this.url}/company`);
    }

    public RequestEmployeeRegistration(employee: Employee): Observable<string> {
        return this.http.post<string>(`${environment.baseApiUrl}/${this.url}/employee/register`, employee);
    }

    public GetEmployeesByCompanyID(companyID: string): Observable<Employee[]> {
        return this.http.get<Employee[]>(`${environment.baseApiUrl}/${this.url}/employees/${companyID}`);
    }

}