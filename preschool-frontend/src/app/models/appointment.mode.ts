import { User } from "./user.model";

export class Appointment {
    id: number = 0;
	dayOfAppointment: number = 0;
    startOfAppointment: Date = new Date;
    endOfAppointment: Date = new Date;
    user: User = new User;
    doctor: User = new User;
}