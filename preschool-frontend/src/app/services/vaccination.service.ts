import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Vaccination } from "../models/vaccination.model";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment";
import { AddVaccination } from "../dto/addVaccination";

@Injectable({
    providedIn: 'root'
})
export class VaccinationService {
    private url = "healthcare";
    constructor(private http: HttpClient) { }

    public GetAllVaccinations(): Observable<Vaccination[]> {
        return this.http.get<Vaccination[]>(`${environment.baseApiUrl}/${this.url}/allVaccinations`);
    }

    public GetAllAvailableVaccinations(): Observable<Vaccination[]> {
        return this.http.get<Vaccination[]>(`${environment.baseApiUrl}/${this.url}/allAvailableVaccinations`);
    }

    public GetSingleVaccination(id: string): Observable<Vaccination> {
        return this.http.get<Vaccination>(`${environment.baseApiUrl}/${this.url}/getVaccinationByID/` + id);
    }

    public GetMyVaccinationsDoctor(): Observable<Vaccination[]> {
        return this.http.get<Vaccination[]>(`${environment.baseApiUrl}/${this.url}/myVaccinationsDoctor`);
    }

    public GetMyAvailableVaccinationsDoctor(): Observable<Vaccination[]> {
        return this.http.get<Vaccination[]>(`${environment.baseApiUrl}/${this.url}/myAvailableVaccinationsDoctor`);
    }

    public GetMyTakenVaccinationsDoctor(): Observable<Vaccination[]> {
        return this.http.get<Vaccination[]>(`${environment.baseApiUrl}/${this.url}/myTakenVaccinationsDoctor`);
    }

    public AddVaccination(addAppointment: AddVaccination): Observable<string> {
        return this.http.post<string>(`${environment.baseApiUrl}/${this.url}/newVaccination`, addAppointment);
    }

    public SetVaccination(id: string) {
        return this.http.put(`${environment.baseApiUrl}/${this.url}/setVaccination/` + id, null);
    }

}