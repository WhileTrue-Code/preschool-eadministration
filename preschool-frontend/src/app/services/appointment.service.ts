import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment";
import { Appointment } from "../models/appointment.mode";
import { AddAppointment } from "../dto/addAppointment";
import { User } from "../models/user.model";

@Injectable({
    providedIn: 'root'
})
export class AppointmentService {
    private url = "healthcare";
    constructor(private http: HttpClient) { }

    public GetAllAppointments(): Observable<Appointment[]> {
        return this.http.get<Appointment[]>(`${environment.baseApiUrl}/${this.url}/allAppointments`);
    }

    public GetMe(): Observable<User> {
        return this.http.get<User>(`${environment.baseApiUrl}/${this.url}/getMe`);
    }

    public GetSingleAppointment(appointment_id: string): Observable<Appointment> {
        return this.http.get<Appointment>(`${environment.baseApiUrl}/${this.url}/getAppointmentByID/` + appointment_id);
    }

    public GetMyAppointmentsDoctor(): Observable<Appointment[]> {
        return this.http.get<Appointment[]>(`${environment.baseApiUrl}/${this.url}/myAppointmentsDoctor`);
    }

    public AddAppointment(addAppointment: AddAppointment): Observable<AddAppointment> {
        return this.http.post<AddAppointment>(`${environment.baseApiUrl}/${this.url}/newAppointment`, addAppointment);
    }

    public SetAppointment(id: string) {
        return this.http.put(`${environment.baseApiUrl}/${this.url}/setAppointment/` + id, null);
    }
}