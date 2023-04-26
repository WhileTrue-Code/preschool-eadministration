package domain

type AprService interface {
	RegisterAprAccount(aprAccount *AprAccount) error
	FindAprByFounderID(founderID string) ([]AprAccount, error)
}
