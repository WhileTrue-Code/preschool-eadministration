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

func (service *HealthcareService) GetAllAppointments() ([]*model.Appointment, error) {
	return service.repository.GetAllAppointments()
}

func (service *HealthcareService) GetAllAvailableAppointments() ([]*model.Appointment, error) {
	return service.repository.GetAllAvailableAppointments()
}

func (service *HealthcareService) GetAppointmentByID(id primitive.ObjectID) (*model.Appointment, error) {
	return service.repository.GetAppointmentByID(id)
}

func (service *HealthcareService) CreateNewAppointment(appointment *model.Appointment, jmbg string) (int, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	//existingAppointments, err := service.repository.GetAllAppointments()
	//for _, existingAppointment := range existingAppointments {
	//	if appointment.DayOfAppointment == existingAppointment.DayOfAppointment {
	//		return 1, nil
	//	}
	//}
	//if err != nil {
	//	log.Println("Error getting all Appointments", err)
	//	return 0, err
	//}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var doctor model.User
	err = json.Unmarshal(response.Data, &doctor)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return 0, err
	}

	appointment.ID = primitive.NewObjectID()
	appointment.Doctor = &doctor
	appointment.User = nil

	err = service.repository.CreateNewAppointment(appointment)
	if err != nil {
		log.Println("Error in trying to save Appointment")
		return 0, err
	}

	return 0, nil
}

func (service *HealthcareService) SetAppointment(id primitive.ObjectID, jmbg string) error {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var user model.User
	err = json.Unmarshal(response.Data, &user)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return err
	}

	appointment, err := service.repository.GetAppointmentByID(id)
	if err != nil {
		log.Println("Error in finding Appointment By ID")
		return err
	}

	appointment.User = &user

	err = service.repository.SetAppointment(appointment)
	if err != nil {
		log.Println("Error in Updating Appointment")
		return err
	}

	return nil
}

func (service *HealthcareService) DeleteAppointmentByID(id primitive.ObjectID) error {
	return service.repository.DeleteAppointmentByID(id)
}

func (service *HealthcareService) GetAllVaccinations() ([]*model.Vaccination, error) {
	return service.repository.GetAllVaccinations()
}
