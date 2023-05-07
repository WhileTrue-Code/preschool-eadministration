package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"healthcare_service/model"
)

type HealthcareRepository interface {
	GetAllAppointments() ([]*model.Appointment, error)
	GetMyAppointmentsDoctor(id primitive.ObjectID) ([]*model.Appointment, error)
	GetAllAvailableAppointments() ([]*model.Appointment, error)
	GetAppointmentByID(id primitive.ObjectID) (*model.Appointment, error)
	CreateNewAppointment(appointment *model.Appointment) error
	SetAppointment(appointment *model.Appointment) error
	DeleteAppointmentByID(id primitive.ObjectID) error

	GetAllVaccinations() ([]*model.Vaccination, error)
	CreateNewVaccination(vaccination model.Vaccination) error
}
