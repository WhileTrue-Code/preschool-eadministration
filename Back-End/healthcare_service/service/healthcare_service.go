package service

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"healthcare_service/model"
	"healthcare_service/repository"
	"log"
	"os"
	"time"
)

type HealthcareService struct {
	store          repository.HealthcareRepository
	natsConnection *nats.Conn
}

func NewHealthcareService(store repository.HealthcareRepository, natsConnection *nats.Conn) *HealthcareService {
	return &HealthcareService{
		store:          store,
		natsConnection: natsConnection,
	}
}

func (service *HealthcareService) CreateNewAppointment(appointment model.Appointment, jmbg string) error {

	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	log.Println(os.Getenv("GET_USER_BY_JMBG"))
	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var doctor model.User
	err = json.Unmarshal(response.Data, &doctor)
	if err != nil {
		log.Println("Error in Unmarshaling json")
		return err
	}
	log.Println(doctor)
	//appointment.Doctor = doctor
	return nil
}
