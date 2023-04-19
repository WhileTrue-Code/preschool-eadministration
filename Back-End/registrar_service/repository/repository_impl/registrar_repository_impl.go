package repository_impl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	domain "registrar_service/model/entity"
	"registrar_service/repository"
)

type RegistrarRepositoryImpl struct {
	user_registry *mongo.Collection //izvodi
	marriage      *mongo.Collection //vencani
}

const (
	DATABASE    = "users"
	COLLECTION  = "user_registry"
	COLLECTION1 = "marriage"
)

func NewRegistrarRepositoryImpl(client *mongo.Client) repository.RegistrarRepository {
	registrar := client.Database(DATABASE).Collection(COLLECTION)
	marriage := client.Database(DATABASE).Collection(COLLECTION1)

	return &RegistrarRepositoryImpl{
		user_registry: registrar,
		marriage:      marriage,
	}
}

func (store *RegistrarRepositoryImpl) CreateNewBirthCertificate(user domain.User) error {
	_, err := store.user_registry.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func (store *RegistrarRepositoryImpl) IsUserExist(jmbg string) bool {

	user, err := store.filterOne(bson.M{"JMBG": jmbg})
	if err != nil {
		log.Println(err.Error())
		return false
	}

	if user != nil {
		return true
	} else {
		return false
	}

}

func (store *RegistrarRepositoryImpl) filterOne(filter interface{}) (user *domain.User, err error) {
	result := store.user_registry.FindOne(context.TODO(), filter)
	err = result.Decode(&user)
	return
}

func (store *RegistrarRepositoryImpl) FindOneUser(jmbg string) *domain.User {

	user, err := store.filterOne(bson.M{"JMBG": jmbg})
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return user

}

func (store *RegistrarRepositoryImpl) CreateNewMarriage(marriage domain.ExcerptFromTheMarriageRegister) {
	_, err := store.marriage.InsertOne(context.Background(), marriage)
	if err != nil {
		log.Printf("Error in RegistrarRepositoryImpl CreateNewMarriage(): %s", err.Error())
		return
	}
}
