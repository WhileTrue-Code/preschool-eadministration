export class Employee {
    id: string = "";
    firstName: string = "";
	lastName: string = "";
    address: string = "";
    employeeID: string = "";
    companyID: number = 0;
    idCardNumber: string = "";
    passportNumber: string = "";
    employmentStatus: string = "";
    employmentDuration: number = 0;
    netSalary: number = 0;
    grossPay: number = 0;
    personalIncomeTax: number = 0;
    pdContribution: number = 0;
    hiContribution: number = 0;
    uiContribution: number = 0;
    efContribution: number = 0;

    Employee(id: string, firstName: string, lastName: string, employeeID: string, companyID: number, 
        address: string, idCardNumber: string, passportNumber: string, employmentStatus: string, 
        employmentDuration: number, netSalary: number, grossPay: number, personalIncomeTax: number,
        pdContribution: number, hiContribution: number, uiContribution: number, efContribution: number) {
        
        this.id = id;
        this.firstName = firstName;
        this.lastName = lastName;
        this.address = address;
        this.employeeID = employeeID;
        this.companyID = companyID;
        this.idCardNumber = idCardNumber;
        this.passportNumber = passportNumber;
        this.employmentStatus = employmentStatus;
        this.employmentDuration = employmentDuration;
        this.netSalary = netSalary;
        this.grossPay = grossPay;
        this.personalIncomeTax = personalIncomeTax;
        this.pdContribution = pdContribution;
        this.hiContribution = hiContribution;
        this.uiContribution = uiContribution;
        this.efContribution = efContribution;
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