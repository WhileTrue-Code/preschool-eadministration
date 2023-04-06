package repository_impl

import (
	"go.mongodb.org/mongo-driver/mongo"
	"registrar_service/repository"
)

type RegistrarRepositoryImpl struct {
	registrar *mongo.Collection //izvodi
}

const (
	DATABASE   = "registrar"
	COLLECTION = "extract_registrar"
)

func NewRegistrarRepositoryImpl(client *mongo.Client) repository.RegistrarRepository {
	registrar := client.Database(DATABASE).Collection(COLLECTION)
	return &RegistrarRepositoryImpl{
		registrar: registrar,
	}
}
