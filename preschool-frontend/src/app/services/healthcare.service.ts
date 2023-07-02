import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment";
import { AddAppointment } from "../dto/addAppointment";
import { AddVaccination } from "../dto/addVaccination";
import { Appointment } from "../models/appointment.model";
import { User } from "../models/user.model";
import { Vaccination } from "../models/vaccination.model";
import { ZdravstvenoStanje } from "../models/zdravstvenoStanje.model";

@Injectable({
    providedIn: 'root'
})
export class HealthcareService {
    private url = "healthcare"
    constructor(private http: HttpClient) { }

    public GetAllAppointments(): Observable<Appointment[]> {
        return this.http.get<Appointment[]>(`${environment.baseApiUrl}/${this.url}/allAppointments`);
    }

    public GetAllAvailableAppointments(): Observable<Appointment[]> {
        return this.http.get<Appointment[]>(`${environment.baseApiUrl}/${this.url}/allAvailableAppointments`);
    }

    public GetSingleAppointment(appointment_id: string): Observable<Appointment> {
        return this.http.get<Appointment>(`${environment.baseApiUrl}/${this.url}/getAppointmentByID/` + appointment_id);
    }

    public GetMyAppointmentsDoctor(): Observable<Appointment[]> {
        return this.http.get<Appointment[]>(`${environment.baseApiUrl}/${this.url}/myAppointmentsDoctor`);
    }

    public GetMyAvailableAppointmentsDoctor(): Observable<Appointment[]> {
        return this.http.get<Appointment[]>(`${environment.baseApiUrl}/${this.url}/myAvailableAppointmentsDoctor`);
    }

    public GetMyTakenAppointmentsDoctor(): Observable<Appointment[]> {
        return this.http.get<Appointment[]>(`${environment.baseApiUrl}/${this.url}/myTakenAppointmentsDoctor`);
    }

    public AddAppointment(addAppointment: AddAppointment): Observable<AddAppointment> {
        return this.http.post<AddAppointment>(`${environment.baseApiUrl}/${this.url}/newAppointment`, addAppointment);
    }

    public SetAppointment(id: string) {
        return this.http.put(`${environment.baseApiUrl}/${this.url}/setAppointment/` + id, null);
    }

    public DeleteAppointment(id: string) {
        return this.http.delete(`${environment.baseApiUrl}/${this.url}/deleteAppointmentByID/` + id);
    }

    public GetMe(): Observable<User> {
        return this.http.get<User>(`${environment.baseApiUrl}/${this.url}/getMe`);
    }

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

    public GetMyTakenVaccinationsRegular(): Observable<Vaccination[]> {
        return this.http.get<Vaccination[]>(`${environment.baseApiUrl}/${this.url}/myTakenVaccinationsRegular`);
    }

    public AddVaccination(addVaccination: AddVaccination): Observable<AddVaccination> {
        return this.http.post<AddVaccination>(`${environment.baseApiUrl}/${this.url}/newVaccination`, addVaccination);
    }

    public SetVaccination(id: string) {
        return this.http.put(`${environment.baseApiUrl}/${this.url}/setVaccination/` + id, null);
    }

    public DeleteVaccination(id: string) {
        return this.http.delete(`${environment.baseApiUrl}/${this.url}/deleteVaccinationByID/` + id);
    }

    public GetAllZdravstvenaStanja(): Observable<ZdravstvenoStanje[]> {
        return this.http.get<ZdravstvenoStanje[]>(`${environment.baseApiUrl}/${this.url}/allZdravstvenaStanja`);
    }

    public GetMyZdravstvenoStanje(): Observable<ZdravstvenoStanje> {
        return this.http.get<ZdravstvenoStanje>(`${environment.baseApiUrl}/${this.url}/myZdravstvenoStanje`)
    }

    public NewZdravstvenoStanje(zdravstvenoStanje: ZdravstvenoStanje): Observable<ZdravstvenoStanje> {
        return this.http.post<ZdravstvenoStanje>(`${environment.baseApiUrl}/${this.url}/newZdravstvenoStanje`, zdravstvenoStanje)
    }

    public AddPersonToRegistry(user: User): Observable<User> {
        return this.http.post<User>(`${environment.baseApiUrl}/${this.url}/addPersonToRegistry`, user)
    }
}