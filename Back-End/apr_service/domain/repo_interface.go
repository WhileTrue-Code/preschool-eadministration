package domain

type AprRepository interface {
	SaveAprAccount(aprAccount *AprAccount) error
	FindAprAccountsByFounderID(founderID string) ([]AprAccount, error)
	FindCompanyByFounderIDAndCompanyID(founderID string, companyID int) (company AprAccount, err error)
	DoesExistAprWithID(ID int) (exists bool)
}
