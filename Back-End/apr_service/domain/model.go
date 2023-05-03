package domain

type AprAccount struct {
	Name                      string `json:"name"`
	Address                   string `json:"address"`
	CompanyID                 string `json:"companyID"`
	Password                  string `json:"password"`
	FounderFirstName          string `json:"founderFirstName"`
	FounderLastName           string `json:"founderLastName"`
	FounderEmail              string `json:"founderEmail"`
	FounderPhone              string `json:"founderPhone"`
	FounderID                 string `json:"founderID"`
	StartCapital              int    `json:"startCapital"`
	AuthorizedPersonFirstName string `json:"authorizedPersonFirstName"`
	AuthorizedPersonLastName  string `json:"authorizedPersonLastName"`
	LastUpdateDate            int    `json:"lastUpdateDate"`
}
