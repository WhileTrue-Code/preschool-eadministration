package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Appointment struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id"`
	DayOfAppointment   primitive.DateTime `json:"dayOfAppointment" bson:"dayOfAppointment"`
	StartOfAppointment primitive.DateTime `json:"startOfAppointment" bson:"startOfAppointment"`
	EndOfAppointment   primitive.DateTime `json:"endOfAppointment" bson:"endOfAppointment"`
	User               User               `json:"user" bson:"user"`
	Doctor             User               `json:"doctor" bson:"doctor"`
}

type Vaccination struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id"`
	DayOfVaccination   primitive.DateTime `json:"dayOfVaccination" bson:"dayOfVaccination"`
	StartOfVaccination primitive.ObjectID `json:"startOfVaccination" bson:"startOfVaccination"`
	EndOfVaccination   primitive.DateTime `json:"endOfVaccination" bson:"endOfVaccination"`
	VaccineType        VaccineType        `json:"vaccineType" bson:"vaccineType"`
	User               User               `json:"user" bson:"user"`
	Doctor             User               `json:"doctor" bson:"doctor"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Firstname string             `bson:"firstName" json:"firstName" validate:"onlyChar"`
	Lastname  string             `bson:"lastName" json:"lastName" validate:"onlyChar"`
	Age       int                `bson:"age" json:"age"`
	Residence string             `bson:"residence" json:"residence" validate:"onlyCharAndSpace"`
	JMBG      string             `bson:"jmbg" json:"jmbg" validate:"onlyCharAndNum,required"`
	Password  string             `bson:"password" json:"password" validate:"onlyCharAndNum,required"`
	UserType  UserType           `bson:"userType" json:"userType" validate:"onlyChar"`
}

type UserType string

const (
	Admin     = "Admin"
	Regular   = "Regular"
	Doctor    = "Doctor"
	Registrar = "Registrar"
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
