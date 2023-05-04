export class AddAppointment {
	dayOfAppointment: number = 0;
    startOfAppointment: number = 0;
    endOfAppointment: number = 0;

    AddAppointment(dayOfAppointment: number, startOfAppointment: number, endOfAppointment: number) {
        this.dayOfAppointment = dayOfAppointment;
        this.startOfAppointment = startOfAppointment;
        this.endOfAppointment = endOfAppointment;
    }
}