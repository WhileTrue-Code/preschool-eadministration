import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment";
import { Appointment } from "../models/appointment.mode";

@Injectable({
    providedIn: 'root'
})
export class AppointmentService {
    private url = "appointments";
    constructor(private http: HttpClient) { }

    public GetAllAppointments(): Observable<Appointment[]> {
        return this.http.get<Appointment[]>(`${environment.baseApiUrl}/${this.url}/allAppointments`);
    }

    public GetSingleAppointment(appointment_id: string): Observable<Appointment> {
        return this.http.get<Appointment>(`${environment.baseApiUrl}/${this.url}/getAppointmentByID/` + appointment_id);
    }

    public AddAppointment(appointment: Appointment): Observable<Appointment> {
        return this.http.post<Appointment>(`${environment.baseApiUrl}/${this.url}/newAppointment`, appointment);
    }
}