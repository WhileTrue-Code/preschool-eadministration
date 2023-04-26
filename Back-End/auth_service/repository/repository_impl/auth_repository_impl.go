package repository_impl

import (
	domain "auth_service/model/entity"
	"auth_service/repository"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const (
	DATABASE   = "auth"
	COLLECTION = "credentials"
)

type AuthRepositoryImpl struct {
	auth *mongo.Collection //izvodi
}

func NewAuthRepositoryImpl(client *mongo.Client) repository.AuthRepository {
	auth := client.Database(DATABASE).Collection(COLLECTION)
	return &AuthRepositoryImpl{
		auth: auth,
	}
}

func (store *AuthRepositoryImpl) IsJMBGUnique(jmbg string) bool {

	exist := false
	filter := bson.M{"jmbg": jmbg}
	credentials, err := store.filterOne(filter)
	if err != nil {
		log.Printf("AuthRepositoryImpl Error IsJMBGUnique(): %s", err)
	}
	if credentials != nil {
		exist = true
	}
	return exist
}

func (store *AuthRepositoryImpl) SignUp(credentials domain.Credentials) {
	_, err := store.auth.InsertOne(context.Background(), credentials) //insert credentials of user
	if err != nil {
		fmt.Errorf("AuthRepositoryImpl Error SignUp(): %s", err)
	}
}

func (store *AuthRepositoryImpl) GetCredentials(jmbg string) (*domain.Credentials, error) {
	filter := bson.M{"jmbg": jmbg}

	credentials, err := store.filterOne(filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return credentials, nil

}

func (store *AuthRepositoryImpl) filterOne(filter interface{}) (credentials *domain.Credentials, err error) {
	result := store.auth.FindOne(context.TODO(), filter)
	err = result.Decode(&credentials)
	return
}
