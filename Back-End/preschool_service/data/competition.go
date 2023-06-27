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
	Uzrast          string             `json:"uzrast" bson:"uzrast"`
	BrojDece        int                `json:"broj_dece" bson:"brojDece"`
	Vrtic           *Vrtic             `json:"vrtic" bson:"vrtic"`
	Status          string             `json:"status" bson:"status"`
}

//type Status string

const (
	Zatvoren = "Zatvoren"
	Otvoren  = "Otvoren"
)

type Vrtic struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Naziv   string             `json:"naziv" bson:"naziv"`
	Adresa  string             `json:"adresa" bson:"adresa"`
	Telefon string             `json:"telefon" bson:"telefon"`
	Email   string             `json:"email" bson:"email"`
	Grad    string             `json:"grad" bson:"grad"`
	Opstina string             `json:"opstina" bson:"opstina"`
}

func (p *Competitions) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Competition) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
func (p *Vrtici) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Vrtic) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Competition) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}
