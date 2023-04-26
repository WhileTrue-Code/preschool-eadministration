package service

import (
	"github.com/nats-io/nats.go"
	"healthcare_service/repository"
)

type HealthcareService struct {
	store          repository.HealthcareRepository
	natsConnection *nats.Conn
}

func NewHealthcareService(store repository.HealthcareRepository, natsConnection *nats.Conn) *HealthcareService {
	return &HealthcareService{
		store:          store,
		natsConnection: natsConnection,
	}
}
