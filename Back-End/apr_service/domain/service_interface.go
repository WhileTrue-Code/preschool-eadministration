package domain

import "github.com/nats-io/nats.go"

type AprService interface {
	RegisterAprAccount(aprAccount *AprAccount) error
	FindAprByFounderID(founderID string) ([]AprAccount, error)
	SubscribeToNats(connection *nats.Conn)
}
