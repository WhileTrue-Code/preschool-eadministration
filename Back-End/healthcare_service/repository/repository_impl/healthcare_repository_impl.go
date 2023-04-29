package repository_impl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"healthcare_service/model"
	"healthcare_service/repository"
)

type HealthcareRepositoryImpl struct {
	appointment *mongo.Collection
	vaccination *mongo.Collection
}

const (
	DATABASE               = "healthcare"
	COLLECTION_APPOINTMENT = "appointment"
	COLLECTION_VACCINATION = "vaccination"
)

func NewAuthRepositoryImpl(client *mongo.Client) repository.HealthcareRepository {
	appointment := client.Database(DATABASE).Collection(COLLECTION_APPOINTMENT)
	vaccination := client.Database(DATABASE).Collection(COLLECTION_VACCINATION)

	return &HealthcareRepositoryImpl{
		appointment: appointment,
		vaccination: vaccination,
	}
}

func (repository *HealthcareRepositoryImpl) GetAllAppointments() ([]*model.Appointment, error) {
	filter := bson.D{{}}
	return repository.filterAppointments(filter)
}

func (repository *HealthcareRepositoryImpl) CreateNewAppointment(appointment *model.Appointment) (*model.Appointment, error) {
	result, err := repository.appointment.InsertOne(context.TODO(), appointment)
	if err != nil {
		return nil, err
	}
	appointment.ID = result.InsertedID.(primitive.ObjectID)
	return appointment, nil
}

func (repository *HealthcareRepositoryImpl) filterAppointments(filter interface{}) ([]*model.Appointment, error) {
	cursor, err := repository.appointment.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeAppointment(cursor)
}

func (repository *HealthcareRepositoryImpl) filterOneAppointment(filter interface{}) (appointment *model.Appointment, err error) {
	result := repository.appointment.FindOne(context.TODO(), filter)
	err = result.Decode(&appointment)
	return
}

func (repository *HealthcareRepositoryImpl) GetAllVaccinations() ([]*model.Vaccination, error) {
	filter := bson.D{{}}
	return repository.filterVaccinations(filter)
}

func (repository *HealthcareRepositoryImpl) CreateNewVaccination(vaccination model.Vaccination) error {
	_, err := repository.vaccination.InsertOne(context.Background(), vaccination)
	if err != nil {
		return err
	}
	return nil
}

func (repository *HealthcareRepositoryImpl) filterVaccinations(filter interface{}) ([]*model.Vaccination, error) {
	cursor, err := repository.vaccination.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeVaccination(cursor)
}

func (repository *HealthcareRepositoryImpl) filterOneVaccination(filter interface{}) (appointment *model.Appointment, err error) {
	result := repository.appointment.FindOne(context.TODO(), filter)
	err = result.Decode(&appointment)
	return
}

func decodeAppointment(cursor *mongo.Cursor) (appointments []*model.Appointment, err error) {
	for cursor.Next(context.TODO()) {
		var appointment model.Appointment
		err = cursor.Decode(&appointment)
		if err != nil {
			return
		}
		appointments = append(appointments, &appointment)
	}
	err = cursor.Err()
	return
}

func decodeVaccination(cursor *mongo.Cursor) (vaccinations []*model.Vaccination, err error) {
	for cursor.Next(context.TODO()) {
		var vaccination model.Vaccination
		err = cursor.Decode(&vaccination)
		if err != nil {
			return
		}
		vaccinations = append(vaccinations, &vaccination)
	}
	err = cursor.Err()
	return
}
