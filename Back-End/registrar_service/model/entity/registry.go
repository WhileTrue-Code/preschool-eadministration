package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	JMBG          string             `json:"jmbg" bson:"JMBG"`
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

type BirthCertificate struct {
	ID            primitive.ObjectID
	Ime           string
	Prezime       string
	ImeOca        string
	JMBGOca       string
	ImeMajke      string
	JMBGMajke     string
	DatumRodjenja string
	MestoRodjenja string
	JMBG          string
	Pol           Pol
}

type ExtractFromTheDeathRegister struct {
	ID            primitive.ObjectID
	Ime           string
	Prezime       string
	ImeOca        string
	JMBGOca       string
	ImeMajke      string
	JMBGMajke     string
	DatumRodjenja string
	MestoRodjenja string
	JMBG          string
	Pol           Pol
	DatimSmrti    string
	MestoSmrti    string
}

type CertificateOfCitizenship struct {
	ID            primitive.ObjectID
	Ime           string
	Prezime       string
	ImeOca        string
	JMBGOca       string
	ImeMajke      string
	JMBGMajke     string
	DatumRodjenja string
	MestoRodjenja string
	JMBG          string
	Pol           Pol
	DatimSmrti    string
	MestoSmrti    string
	Drzava        string
}

type ExcerptFromTheMarriageRegister struct {
	ID                     primitive.ObjectID `json:"id" bson:"_id"`
	ImeMladozenje          string             `json:"ime_mladozenje" bson:"ime_mladozenje"`
	ImeMlade               string             `json:"ime_mlade" bson:"ime_mlade"`
	PrezimeMladozenje      string             `json:"prezime_mladozenje" bson:"prezime_mladozenje"`
	DevojkackoPrezimeMlade string             `json:"devojkacko_prezime_mlade" bson:"devojkacko_prezime_mlade"`
	DatumVencanja          int64              `json:"datum_vencanja" bson:"datum_vencanja"`
	MestoVencanja          string             `json:"mesto_vencanja" bson:"mesto_vencanja"`
	JMBGMladozenje         string             `json:"jmbg_mladozenje" bson:"jmbg_mladozenje"`
	JMBGMlade              string             `json:"jmbg_mlade" bson:"jmbg_mlade"`
	Svedok1                User               `json:"svedok_1" bson:"svedok_1"`
	Svedok2                User               `json:"svedok_2" bson:"svedok_2"`
}
