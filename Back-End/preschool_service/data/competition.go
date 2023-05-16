package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

type Competition struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	DatumObjave     int64              `json:"datum_objave" bson:"datumObjave"`
	PocetakKonkursa int64              `json:"pocetak_konkursa" bson:"pocetakKonkursa"`
	KrajKonkursa    int64              `json:"kraj_konkursa" bson:"krajKonkursa"`
	Grad            string             `json:"grad" bson:"grad"`
	Opstina         string             `json:"opstina" bson:"opstina"`
	Uzrast          string             `json:"uzrast" bson:"uzrast"`
	BrojDece        int64              `json:"broj_dece" bson:"brojDece"`
	//Prijava         *Prijava           `json:"prijava" bson:"prijava"`
	//Vrtic           *Vrtic              `json:"vrtic" bson:"vrtic"`
}

//type Vrtic struct {
//	ID      primitive.ObjectID `json:"id" bson:"_id"`
//	Naziv   string             `json:"naziv" bson:"naziv"`
//	Adresa  string             `json:"adresa" bson:"adresa"`
//	Telefon string             `json:"telefon" bson:"telefon"`
//	Email   string             `json:"email" bson:"email"`
//}

func (p *Competitions) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Competition) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Competition) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}
