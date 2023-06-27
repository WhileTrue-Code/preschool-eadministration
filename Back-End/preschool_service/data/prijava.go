package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

type Prijava struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	CompetitionID primitive.ObjectID `json:"competition_id" bson:"_idCompetition"`
	Bodovi        int                `json:"bodovi" bson:"bodovi"`
	Dete          Dete               `json:"dete" bson:"dete"`
	Status        string             `json:"status" bson:"status"`
}

//type Status string

const (
	Prijavljen = "Prijavljen"
	Odbijen    = "Odbijen"
	Upisan     = "Upisan"
)

type Dete struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	JMBG              string             `json:"jmbg" bson:"JMBG"`
	DatumRodjenja     int64              `json:"datum_rodjenja" bson:"datumRodjenja"`
	Ime               string             `json:"ime" bson:"ime"`
	Prezime           string             `json:"prezime" bson:"prezime"`
	Opstina           string             `json:"opstina" bson:"opstina"`
	Adresa            string             `json:"adresa" bson:"adresa"`
	ZdravstvenoStanje *ZdravstvenoStanje `json:"zdravstveno_stanje" bson:"zdravstvenoStanje"`
	//JMBGOca           string             `json:"jmbg_oca" bson:"JMBGOca"`
	//JMBGMajke         string             `json:"jmbg_majke" bson:"JMBGMajke"`
}

type ZdravstvenoStanje struct {
	ZdravstveniProblemi     string `json:"zdravstveni_problemi" bson:"zdravstveniProblemi"`
	SpecificnaIshrana       string `json:"specificna_ishrana" bson:"specificnaIshrana"`
	DomZdravljaUKomJeKarton string `json:"dom_zdravlja_u_kom_je_karton" bson:"domZdravljaUKomJeKarton"`
	SmetnjeURazvoju         string `json:"smetnje_u_razvoju" bson:"smetnjeURazvoju"`
	SpecificniPodaci        string `json:"specificni_podaci" bson:"specificniPodaci"`
}

func (p *Prijave) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Prijava) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
