package service

import (
	"github.com/nats-io/nats.go"
	"healthcare_service/model"
	"healthcare_service/repository"
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

func (service *HealthcareService) CreateNewAppointment(appointment model.Appointment, doctorID string) error {

	//primitiveID, err := primitive.ObjectIDFromHex(doctorID)
	//if err != nil {
	//	log.Println("Primitive ID parsing error.")
	//	return err
	//}

	//user, err := service.store.Get(primitiveID)
	//if err != nil {
	//
	//}

	//appointment.ID = primitive.NewObjectID()
	//appointment.Doctor = user

	//err = service.store.CreateNewAppointment(appointment)
	//if err != nil {
	//	return err
	//}
	return nil
}
