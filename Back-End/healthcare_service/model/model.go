package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Appointment struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id"`
	StartOfAppointment int64              `json:"startOfAppointment" bson:"startOfAppointment"`
	EndOfAppointment   int64              `json:"endOfAppointment" bson:"endOfAppointment"`
	User               *User              `json:"user" bson:"user"`
	Doctor             *User              `json:"doctor" bson:"doctor"`
}

type Vaccination struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id"`
	StartOfVaccination int64              `json:"startOfVaccination" bson:"startOfVaccination"`
	EndOfVaccination   int64              `json:"endOfVaccination" bson:"endOfVaccination"`
	VaccineType        VaccineType        `json:"vaccineType" bson:"vaccineType"`
	User               *User              `json:"user" bson:"user"`
	Doctor             *User              `json:"doctor" bson:"doctor"`
}

type User struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Ime           string             `json:"ime" bson:"ime"`
	Prezime       string             `json:"prezime" bson:"prezime"`
	ImeOca        string             `json:"ime_oca" bson:"imeOca"`
	JMBGOca       string             `json:"jmbg_oca" bson:"JMBGOca"`
	ImeMajke      string             `json:"ime_majke" bson:"imeMajke"`
	JMBGMajke     string             `json:"jmbg_majke" bson:"JMBGMajke"`
	DatumRodjenja int64              `json:"datum_rodjenja" bson:"datumRodjenja"`
	MestoRodjenja string             `json:"mesto_rodjenja" bson:"mestoRodjenja"`
	JMBG          string             `json:"jmbg" bson:"JMBG" unique:"true"`
	Pol           Pol                `json:"pol" bson:"pol"`
	Preminuo      bool               `json:"preminuo" bson:"Preminuo"`
	DatimSmrti    int64              `json:"datim_smrti" bson:"DatimSmrti"`
	MestoSmrti    string             `json:"mesto_smrti" bson:"MestoSmrti"`
	Drzava        string             `json:"drzava" bson:"Drzava"`
}

type Pol string

const (
	Muski  = "Muski"
	Zenski = "Zenski"
)

type VaccineType string

const (
	BCG = "BCG"
	HB  = "HB"
	DTP = "DTP"
	IPV = "IPV"
	HIB = "HIB"
	PCV = "PCV"
)

type ZdravstvenoStanje struct {
	ID                      primitive.ObjectID `json:"id" bson:"_id"`
	Jmbg                    string             `json:"jmbg" bson:"jmbg"`
	ZdravstveniProblemi     string             `json:"zdravstveni_problemi" bson:"zdravstveniProblemi"`
	SpecificnaIshrana       string             `json:"specificna_ishrana" bson:"specificnaIshrana"`
	DomZdravljaUKomJeKarton string             `json:"dom_zdravlja_u_kom_je_karton" bson:"domZdravljaUKomJeKarton"`
	SmetnjeURazvoju         string             `json:"smetnje_u_razvoju" bson:"smetnjeURazvoju"`
	SpecificniPodaci        string             `json:"specificni_podaci" bson:"specificniPodaci"`
}
