package repository_impl

import (
	"go.mongodb.org/mongo-driver/mongo"
	"registrar_service/repository"
)

type AuthRepositoryImpl struct {
	auth *mongo.Collection //izvodi
}

const (
	DATABASE   = "auth"
	COLLECTION = "credentials"
)

func NewAuthRepositoryImpl(client *mongo.Client) repository.AuthRepository {
	auth := client.Database(DATABASE).Collection(COLLECTION)
	return &AuthRepositoryImpl{
		auth: auth,
	}
}
