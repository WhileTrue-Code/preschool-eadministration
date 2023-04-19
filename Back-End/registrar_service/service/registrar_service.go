package service

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
	"registrar_service/model/entity"
	"registrar_service/repository"
)

type RegistrarService struct {
	store          repository.RegistrarRepository
	natsConnection *nats.Conn
}

func NewRegistrarService(store repository.RegistrarRepository, natsConnection *nats.Conn) *RegistrarService {
	return &RegistrarService{
		store:          store,
		natsConnection: natsConnection,
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

func (service *RegistrarService) SubscribeToNats(natsConnection *nats.Conn) {

	_, err := natsConnection.QueueSubscribe(os.Getenv("CHECK_USER_JMBG"), "queue-group", func(message *nats.Msg) {

		var credentials entity.Credentials
		err := json.Unmarshal(message.Data, &credentials)
		if err != nil {
			log.Println("Error in unmarshal JSON!")
			return
		}

		isExist := service.store.IsUserExist(credentials.JMBG)

		dataToSend, err := json.Marshal(isExist)
		if err != nil {
			log.Println("Error in marshaling json")
			return
		}
		reply := dataToSend
		err = natsConnection.Publish(message.Reply, reply)
		if err != nil {
			log.Printf("Error in publish response: %s", err.Error())
			return
		}
	})

	if err != nil {
		log.Printf("Error in receiving message: %s", err.Error())
		return
	}

	log.Printf("Subscribed to channel: %s", os.Getenv("CHECK_USER_JMBG"))

}
