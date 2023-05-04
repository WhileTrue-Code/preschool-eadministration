package domain

type AprAccount struct {
	ID                        string `bson:"_id" json:"id"`
	CompanyID                 string `bson:"companyID" json:"companyID"`
	Name                      string `bson:"name" json:"name"`
	Address                   string `bson:"address" json:"address"`
	FounderID                 string `bson:"founderID" json:"founderID"`
	StartCapital              int    `bson:"startCapital" json:"startCapital"`
	AuthorizedPersonFirstName string `bson:"authorizedPersonFirstName" json:"authorizedPersonFirstName"`
	AuthorizedPersonLastName  string `bson:"authorizedPersonLastName" json:"authorizedPersonLastName"`
	LastUpdateDate            int    `bson:"lastUpdateDate" json:"lastUpdateDate"`
}
