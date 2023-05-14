import { User } from "./user.model";

export class Vaccination {
    id: number = 0;
    startOfVaccination: number = 0;
    endOfVaccination: number = 0;
    vaccineType: string = "";
    user: User = new User;
    doctor: User = new User;

    Appointment(id: number, startOfVaccination: number, endOfVaccination: number, vaccineType: string, user: User, doctor: User) {
        this.id = id;
        this.startOfVaccination = startOfVaccination;
        this.endOfVaccination = endOfVaccination;
        this.vaccineType = vaccineType;
        this.user = user;
        this.doctor = doctor;
    }
}