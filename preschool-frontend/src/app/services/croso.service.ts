import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { AppComponent } from "../app.component";
import { environment } from "src/environments/environment";
import { Competition } from "../models/competition.model";
import { Company } from "../models/aprAccount.model";
import { CompanyID } from "../models/companyID.model";
import { Employee } from "../models/employee.model";
import { ChangeEmploymentStatus } from "../dto/changeEmploymentStatus";
import { ResolveStatus } from "../dto/resolveStatus";

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

    public GetPendingEmployees(): Observable<Employee[]> {
        return this.http.get<Employee[]>(`${environment.baseApiUrl}/${this.url}/employee/pending`);
    }

    public ResolveRegisterStatus(resolveStatus: ResolveStatus): Observable<string> {
        return this.http.patch<string>(`${environment.baseApiUrl}/${this.url}/employee/status`, resolveStatus);
    }

    public ChangeEmploymentStatus(id: string, changeEmploymentStatus: ChangeEmploymentStatus): Observable<string> {
        return this.http.patch<string>(`${environment.baseApiUrl}/${this.url}/employees/${id}/employmentStatus`, changeEmploymentStatus);
    }

    public CancelEmployment(id: string): Observable<string> {
        return this.http.patch<string>(`${environment.baseApiUrl}/${this.url}/employees/${id}/cancelEmployment`, null);
    }

}