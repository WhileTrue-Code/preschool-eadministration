package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"time"
)

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

type Credentials struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	JMBG     string             `bson:"jmbg" json:"jmbg" validate:"onlyCharAndNum,required"`
	Password string             `bson:"password" json:"password" validate:"onlyCharAndNum,required"`
	UserType UserType           `bson:"userType" json:"userType" validate:"onlyChar"`
}

type UserType string

const (
	Admin     = "Admin"
	Regular   = "Regular"
	Doctor    = "Doctor"
	Registrar = "Registrar"
)

type Claims struct {
	UserID    primitive.ObjectID `json:"user_id"`
	JMBG      string             `json:"jmbg"`
	Role      UserType           `json:"userType"`
	ExpiresAt time.Time          `json:"expires_at"`
}

func (user *User) FromJSON(reader io.Reader) error {
	d := json.NewDecoder(reader)
	return d.Decode(user)
}
