package domain

type AprRepository interface {
	SaveAprAccount(aprAccount *AprAccount) error
	FindAprAccountsByFounderID(founderID string) ([]AprAccount, error)
<<<<<<< Updated upstream
	FindAprAccountsByCompanyID(companyID int) (found AprAccount, err error)
	FindCompanyByFounderIDAndCompanyID(founderID string, companyID int) (company AprAccount, err error)
	DoesExistAprWithID(ID int) (exists bool)
	PatchCompany(newCompany AprAccount) (err error)
=======
	DoesExistAprWithID(ID int) (exists bool)
>>>>>>> Stashed changes
}
