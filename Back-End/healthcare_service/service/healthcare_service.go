package service

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"healthcare_service/model"
	"healthcare_service/repository"
	"log"
	"os"
	"time"
)

type HealthcareService struct {
	repository     repository.HealthcareRepository
	natsConnection *nats.Conn
}

func NewHealthcareService(repository repository.HealthcareRepository, natsConnection *nats.Conn) *HealthcareService {
	return &HealthcareService{
		repository:     repository,
		natsConnection: natsConnection,
	}
}

func (service *HealthcareService) CreateNewAppointment(appointment model.Appointment, jmbg string) (int, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var doctor model.User
	err = json.Unmarshal(response.Data, &doctor)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return 1, err
	}

	appointment.ID = primitive.NewObjectID()
	appointment.Doctor = doctor
	appointment.User = nil
	appointment.StartOfAppointment = time.Now()
	appointment.EndOfAppointment = time.Now()

	err = service.repository.CreateNewAppointment(appointment)
	if err != nil {
		log.Println("Error in trying to save Appointment")
		return 0, err
	}

	return 0, nil
}

func (service *HealthcareService) GetAllAppointments() ([]*model.Appointment, error) {
	return service.repository.GetAllAppointments()
}

func (service *HealthcareService) GetAllAvailableAppointments() ([]*model.Appointment, error) {
	return service.repository.GetAllAvailableAppointments()
}

func (service *HealthcareService) GetAllVaccinations() ([]*model.Vaccination, error) {
	return service.repository.GetAllVaccinations()
}
