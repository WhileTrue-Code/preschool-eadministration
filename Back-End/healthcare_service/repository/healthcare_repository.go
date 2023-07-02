package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"healthcare_service/model"
)

type HealthcareRepository interface {
	GetAllAppointments() ([]*model.Appointment, error)
	GetMyAppointmentsDoctor(id primitive.ObjectID) ([]*model.Appointment, error)
	GetMyAvailableAppointmentsDoctor(id primitive.ObjectID) ([]*model.Appointment, error)
	GetMyTakenAppointmentsDoctor(id primitive.ObjectID) ([]*model.Appointment, error)
	GetAllAvailableAppointments() ([]*model.Appointment, error)
	GetAppointmentByID(id primitive.ObjectID) (*model.Appointment, error)
	CreateNewAppointment(appointment *model.Appointment) error
	SetAppointment(appointment *model.Appointment) error
	DeleteAppointmentByID(id primitive.ObjectID) error

	GetAllVaccinations() ([]*model.Vaccination, error)
	GetMyVaccinationsDoctor(id primitive.ObjectID) ([]*model.Vaccination, error)
	GetMyAvailableVaccinationsDoctor(id primitive.ObjectID) ([]*model.Vaccination, error)
	GetMyTakenVaccinationsDoctor(id primitive.ObjectID) ([]*model.Vaccination, error)
	GetAllAvailableVaccinations() ([]*model.Vaccination, error)
	GetMyTakenVaccinationsRegular(id primitive.ObjectID) ([]*model.Vaccination, error)
	GetVaccinationByID(id primitive.ObjectID) (*model.Vaccination, error)
	CreateNewVaccination(vaccination *model.Vaccination) error
	SetVaccination(vaccination *model.Vaccination) error
	DeleteVaccinationByID(id primitive.ObjectID) error

	GetAllZdravstvenoStanje() ([]*model.ZdravstvenoStanje, error)
	GetZdravstvenoStanjeByID(id primitive.ObjectID) (*model.ZdravstvenoStanje, error)
	GetZdravstvenoStanjeByJMBG(jmbg string) (*model.ZdravstvenoStanje, error)
	CreateNewZdravstvenoStanje(zdravstvenoStanje *model.ZdravstvenoStanje) error
	DeleteZdravstvenoStanjeByJMBG(jmbg string) error
}
