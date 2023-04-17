package domain

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Firstname string             `bson:"firstName,omitempty" json:"firstName,omitempty" validate:"onlyChar"`
	Lastname  string             `bson:"lastName,omitempty" json:"lastName,omitempty" validate:"onlyChar"`
	Age       int                `bson:"age,omitempty" json:"age,omitempty"`
	Residence string             `bson:"residence,omitempty" json:"residence,omitempty" validate:"onlyCharAndSpace"`
	JMBG      string             `bson:"username" json:"jmbg" validate:"onlyCharAndNum,required"`
	Password  string             `bson:"password" json:"password" validate:"onlyCharAndNum,required"`
	UserType  UserType           `bson:"userType" json:"userType" validate:"onlyChar"`
}

type UserType string

const (
	Admin     = "Admin"
	Citizen   = "Citizen"
	Registrar = "Registrar"
)

type Claims struct {
	UserID    primitive.ObjectID `json:"user_id"`
	Username  string             `json:"username"`
	Role      UserType           `json:"userType"`
	ExpiresAt time.Time          `json:"expires_at"`
}

func (user *User) FromJSON(reader io.Reader) error {
	d := json.NewDecoder(reader)
	return d.Decode(user)
}
