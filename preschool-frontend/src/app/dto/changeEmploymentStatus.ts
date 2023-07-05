export class ChangeEmploymentStatus {
    employmentStatus: string = '';
    employmentDuration: number = 0;

    AddVaccination(employmentStatus: string, employmentDuration: number) {
        this.employmentStatus = employmentStatus;
        this.employmentDuration = employmentDuration;
    }
}