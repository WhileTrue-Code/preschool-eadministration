package service

import (
	"encoding/json"
	"fmt"
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

func (service *RegistrarService) FindOneCertificateByType(jmbg string, certificateType int) (*entity.BirthCertificate, *entity.ExtractFromTheDeathRegister, *entity.CertificateOfCitizenship) {

	user := service.store.FindOneUser(jmbg)

	if certificateType == 1 {

		var certificate entity.BirthCertificate
		certificate.Ime = user.Ime
		certificate.Prezime = user.Prezime
		certificate.ImeOca = user.ImeOca
		certificate.JMBGOca = user.JMBGOca
		certificate.ImeMajke = user.ImeMajke
		certificate.JMBGMajke = user.JMBGMajke
		certificate.DatumRodjenja = user.DatumRodjenja
		certificate.MestoRodjenja = user.MestoRodjenja
		certificate.JMBG = user.JMBG
		certificate.Pol = user.Pol

		//slanje nazad

		return &certificate, nil, nil

	} else if certificateType == 2 {

		if user.Preminuo == true {
			var certificate entity.ExtractFromTheDeathRegister
			certificate.Ime = user.Ime
			certificate.Prezime = user.Prezime
			certificate.ImeOca = user.ImeOca
			certificate.JMBGOca = user.JMBGOca
			certificate.ImeMajke = user.ImeMajke
			certificate.JMBGMajke = user.JMBGMajke
			certificate.DatumRodjenja = user.DatumRodjenja
			certificate.MestoRodjenja = user.MestoRodjenja
			certificate.JMBG = user.JMBG
			certificate.Pol = user.Pol
			certificate.DatimSmrti = user.DatimSmrti
			certificate.MestoSmrti = user.MestoSmrti

			//slanje nazad

			return nil, &certificate, nil

		}

	} else if certificateType == 3 {

		var certificate entity.CertificateOfCitizenship
		certificate.Ime = user.Ime
		certificate.Prezime = user.Prezime
		certificate.ImeOca = user.ImeOca
		certificate.JMBGOca = user.JMBGOca
		certificate.ImeMajke = user.ImeMajke
		certificate.JMBGMajke = user.JMBGMajke
		certificate.DatumRodjenja = user.DatumRodjenja
		certificate.MestoRodjenja = user.MestoRodjenja
		certificate.JMBG = user.JMBG
		certificate.Pol = user.Pol
		certificate.Drzava = user.Drzava

		//slanje nazad

		return nil, nil, &certificate
	}

	return nil, nil, nil

}

func (service *RegistrarService) UpdateCertificate(died entity.UserDied) error {

	user := service.store.FindOneUser(died.JMBG)

	if user == nil {
		return fmt.Errorf("user not exist in database")
	}

	user.Preminuo = true
	user.MestoSmrti = died.MestoSmrti
	user.DatimSmrti = died.DatimSmrti

	err := service.store.UpdateCertificate(*user)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return err
	}
	return nil
}

func (service *RegistrarService) FindOneUser(jmbg string) *entity.User {
	return service.store.FindOneUser(jmbg)
}

func (service *RegistrarService) CreateNewMarriage(marriage entity.ExcerptFromTheMarriageRegister) {
	marriage.ID = primitive.NewObjectID()
	service.store.CreateNewMarriage(marriage)
}

func (service *RegistrarService) SubscribeToNats(natsConnection *nats.Conn) {

	_, err := natsConnection.QueueSubscribe(os.Getenv("CHECK_USER_JMBG"), "queue-registrar-group", func(message *nats.Msg) {

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

	_, err = natsConnection.QueueSubscribe(os.Getenv("GET_USER_BY_JMBG"), "queue-registrar-group", func(message *nats.Msg) {
		var jmbg string
		err := json.Unmarshal(message.Data, &jmbg)
		if err != nil {
			log.Println("Error in unmarshal JSON!")
			return
		}

		user := service.FindOneUser(jmbg)

		dataToSend, err := json.Marshal(user)
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
		log.Println("Error in receiving message: %s", err.Error())
	}

	log.Printf("Subscribed to channel: %s", os.Getenv("GET_USER_BY_JMBG"))

	_, err = natsConnection.QueueSubscribe(os.Getenv("CREATE_USER"), "queue-registrar-group", func(message *nats.Msg) {
		var user entity.User
		err := json.Unmarshal(message.Data, &user)
		if err != nil {
			log.Println("Error in unmarshal JSON!")
			return
		}

		user.ID = primitive.NewObjectID()
		err = service.store.CreateNewBirthCertificate(user)
		if err != nil {
			user.ID = primitive.NilObjectID
			log.Println("Error in Nats")
		}

		dataToSend, err := json.Marshal(user)
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
		log.Println("Error in receiving message: %s", err.Error())
	}

	log.Printf("Subscribed to channel: %s", os.Getenv("CREATE_USER"))
}

func (service *RegistrarService) GetChildren(jmbg string) []entity.User {
	return service.store.GetChildren(jmbg, service.FindOneUser(jmbg).Pol)
}

//func (service *RegistrarService) ReturnUserForHealthcare(natsConnection *nats.Conn) {
//	_, err := natsConnection.QueueSubscribe(os.Getenv("CHECK_USER_JMBG"), "queue-group", func(message *nats.Msg) {
//		var jmbg string
//		err := json.Unmarshal(message.Data, &jmbg)
//		if err != nil {
//			log.Println("Error in unmarshal JSON!")
//			return
//		}
//	})
//	if err != nil {
//		log.Println("Error in receiving message: %s", err.Error())
//	}
//}
