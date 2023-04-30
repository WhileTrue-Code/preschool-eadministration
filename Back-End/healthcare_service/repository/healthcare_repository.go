package repository

import "healthcare_service/model"

type HealthcareRepository interface {
	GetAllAppointments() ([]*model.Appointment, error)
	GetAllAvailableAppointments() ([]*model.Appointment, error)
	CreateNewAppointment(appointment model.Appointment) error

	GetAllVaccinations() ([]*model.Vaccination, error)
	CreateNewVaccination(vaccination model.Vaccination) error
}
