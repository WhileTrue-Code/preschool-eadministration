package domain

type AprRepository interface {
	SaveAprAccount(aprAccount *AprAccount) error
	FindAprAccountsByFounderID(founderID string) ([]AprAccount, error)
}