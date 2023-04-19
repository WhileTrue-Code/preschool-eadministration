package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"registrar_service/model/entity"
	"registrar_service/repository"
)

type RegistrarService struct {
	store repository.RegistrarRepository
}

func NewRegistrarService(store repository.RegistrarRepository) *RegistrarService {
	return &RegistrarService{
		store: store,
	}
}

func (service *RegistrarService) CreateNewBirthCertificate(user entity.User) (int, error) {

	isExist := service.store.IsUserExist(user.JMBG)
	log.Println(isExist)
	if isExist {
		return 1, nil
	}

	user.ID = primitive.NewObjectID()
	user.Preminuo = false
	err := service.store.CreateNewBirthCertificate(user)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func (service *RegistrarService) FindOneCertificate(jmbg string, certificateType int) {

	//user := service.store.FindOneUser(jmbg)
	//
	//if certificateType == 1 {
	//
	//}

}

func (service *RegistrarService) FindOneUser(jmbg string) *entity.User {
	return service.store.FindOneUser(jmbg)
}

func (service *RegistrarService) CreateNewMarriage(marriage entity.ExcerptFromTheMarriageRegister) {
	marriage.ID = primitive.NewObjectID()
	service.store.CreateNewMarriage(marriage)
}
