package domain

<<<<<<< Updated upstream
import (
	"model"
)

type AprAccount model.CompanyAccount
=======
import "go.mongodb.org/mongo-driver/bson/primitive"

type AprAccount struct {
	ID                        primitive.ObjectID `bson:"_id" json:"id"`
	CompanyID                 int                `bson:"companyID" json:"companyID"`
	Name                      string             `bson:"name" json:"name"`
	Address                   string             `bson:"address" json:"address"`
	FounderID                 string             `bson:"founderID" json:"founderID"`
	StartCapital              int                `bson:"startCapital" json:"startCapital"`
	AuthorizedPersonFirstName string             `bson:"authorizedPersonFirstName" json:"authorizedPersonFirstName"`
	AuthorizedPersonLastName  string             `bson:"authorizedPersonLastName" json:"authorizedPersonLastName"`
	LastUpdateDate            int                `bson:"lastUpdateDate" json:"lastUpdateDate"`
}
>>>>>>> Stashed changes
