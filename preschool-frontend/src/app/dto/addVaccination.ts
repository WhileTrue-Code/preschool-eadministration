export class AddVaccination {
    startOfVaccination: number = 0;
    endOfVaccination: number = 0;
    vaccineType: string = "";

    AddVaccination(startOfVaccination: number, endOfVaccination: number, vaccineType: string) {
        this.startOfVaccination = startOfVaccination;
        this.endOfVaccination = endOfVaccination;
        this.vaccineType = vaccineType;
    }
}