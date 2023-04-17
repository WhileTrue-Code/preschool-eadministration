package repository_impl

import (
	"go.mongodb.org/mongo-driver/mongo"
	"registrar_service/repository"
)

type AuthRepositoryImpl struct {
	registrar *mongo.Collection //izvodi
}

const (
	DATABASE   = "credentials"
	COLLECTION = "user_credentials"
)

func NewAuthRepositoryImpl(client *mongo.Client) repository.AuthRepository {
	registrar := client.Database(DATABASE).Collection(COLLECTION)
	return &AuthRepositoryImpl{
		registrar: registrar,
	}
}
