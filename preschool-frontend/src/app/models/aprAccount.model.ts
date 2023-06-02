export class Company {
    id: string = "";
	companyID: number = 0;
    name: string = "";
    address: string = "";
    startCapital: number = 0;
    authorizedPersonFirstName: string = "";
    authorizedPersonLastName: string = "";

    Company(id: string, companyID: number, name: string, address: string, startCapital: number, 
        authorizedPersonFirstName: string, authorizedPersonLastName: string) {
        this.id = id;
        this.companyID = companyID;
        this.name = name;
        this.address = address;
        this.startCapital = startCapital;
        this.authorizedPersonFirstName = authorizedPersonFirstName;
        this.authorizedPersonLastName = authorizedPersonLastName;
    }
}