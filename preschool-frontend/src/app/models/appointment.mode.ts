import { User } from "./user.model";

export class Appointment {
    id: number = 0;
	dayOfAppointment: number = 0;
    startOfAppointment: number = 0;
    endOfAppointment: number = 0;
    user: User = new User;
    doctor: User = new User;

    Appointment(id: number, dayOfAppointment: number, startOfAppointment: number, endOfAppointment: number, user: User, doctor: User) {
        this.id = id;
        this.dayOfAppointment = dayOfAppointment;
        this.startOfAppointment = startOfAppointment;
        this.endOfAppointment = endOfAppointment;
        this.user = user;
        this.doctor = doctor;
    }
}