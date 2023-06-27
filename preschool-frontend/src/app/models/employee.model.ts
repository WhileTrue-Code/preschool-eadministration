export class Employee {
    firstName: string = "";
	lastName: string = "";
    address: string = "";
    employeeID: string = "";
    companyID: number = 0;
    idCardNumber: string = "";
    passportNumber: string = "";
    employmentStatus: string = "";
    employmentDuration: number = 0;

    Employee(firstName: string, lastName: string, employeeID: string, companyID: number, 
        address: string, idCardNumber: string, passportNumber: string, employmentStatus: string, 
        employmentDuration: number) {
        this.firstName = firstName;
        this.lastName = lastName;
        this.address = address;
        this.employeeID = employeeID;
        this.companyID = companyID;
        this.idCardNumber = idCardNumber;
        this.passportNumber = passportNumber;
        this.employmentStatus = employmentStatus;
        this.employmentDuration = employmentDuration;
    }

    GetRsEmploymentStatus(): string {
        console.log("tu sam")
        if (this.employmentStatus === "definite_contract") {
            return "NA ODREDJENO"
        }else if (this.employmentStatus === "indefinite_contract"){
            return "NA NEODREDJENO"
        }else if (this.employmentStatus === "temporary_works"){
            return "PRIVREMENO POVREMENI RAD"
        }else{
            return "NEZAPOSLJEN/ODJAVLJEN"
        }
    }
}