export class UserDied {
    jmbg: string = ""
    datumSmrti: number = 0
    mestoSmrti: string = ""

    UserDied(jmbg: string, datumSmrti: number, mestoSmrti: string) {
        this.jmbg = jmbg
        this.datumSmrti = datumSmrti
        this.mestoSmrti = mestoSmrti
    }
}