import { Injectable } from "@angular/core";
import { Competition } from "../models/competition.model";
import { HttpClient } from "@angular/common/http";
import { Vrtic } from "../models/vrtic";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment";


@Injectable({
    providedIn: 'root'
})
export class VrticService {
    private url = "preschool"
    constructor(private http: HttpClient) { }

    public GetAllVrtici(): Observable<Vrtic[]>{
        return this.http.get<Vrtic[]>(`${environment.baseApiUrl}/${this.url}/vrtici/all`);
    }

    public AddVrtic(vrtic: Vrtic): Observable<Vrtic> {
        return this.http.post<Vrtic>(`${environment.baseApiUrl}/${this.url}/vrtic/add`, vrtic);
    }

    public GetSingleVrtic(vrtic_id: string): Observable<Vrtic> {
        return this.http.get<Vrtic>(`${environment.baseApiUrl}/${this.url}/vrtic/` + vrtic_id);
    }
}