export class ResolveStatus {
    companyID: number = 0;
    employeeID: string = "";
    status: number = 0;

    ResolveStatus(companyID: number, employeeID: string, status: number) {
        this.companyID = companyID;
        this.employeeID = employeeID;
        this.status = status;
    }
}