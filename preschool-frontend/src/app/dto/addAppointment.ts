export class AddAppointment {
    startOfAppointment: number = 0;
    endOfAppointment: number = 0;

    AddAppointment(startOfAppointment: number, endOfAppointment: number) {
        this.startOfAppointment = startOfAppointment;
        this.endOfAppointment = endOfAppointment;
    }
}