package repository_impl

import (
	"go.mongodb.org/mongo-driver/mongo"
	"healthcare_service/repository"
)

const (
	DATABASE   = "healthcare"
	COLLECTION = "hospital"
)

type HealthcareRepositoryImpl struct {
	healthcare *mongo.Collection //izvodi
}

func NewAuthRepositoryImpl(client *mongo.Client) repository.HealthcareRepository {
	healthcare := client.Database(DATABASE).Collection(COLLECTION)
	return &HealthcareRepositoryImpl{
		healthcare: healthcare,
	}
}
