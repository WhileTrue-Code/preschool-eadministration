package repository_impl

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"healthcare_service/model"
	"healthcare_service/repository"
)

type HealthcareRepositoryImpl struct {
	healthcare *mongo.Collection
}

const (
	DATABASE   = "healthcare"
	COLLECTION = "hospital"
)

func NewAuthRepositoryImpl(client *mongo.Client) repository.HealthcareRepository {
	healthcare := client.Database(DATABASE).Collection(COLLECTION)

	return &HealthcareRepositoryImpl{
		healthcare: healthcare,
	}
}

func (store *HealthcareRepositoryImpl) CreateNewAppointment(appointment model.Appointment) error {
	_, err := store.healthcare.InsertOne(context.Background(), appointment)
	if err != nil {
		return err
	}
	return nil
}

func (store *HealthcareRepositoryImpl) CreateNewVaccination(vaccination model.Vaccination) error {
	_, err := store.healthcare.InsertOne(context.Background(), vaccination)
	if err != nil {
		return err
	}
	return nil
}
