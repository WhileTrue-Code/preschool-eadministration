package repository

import "healthcare_service/model"

type HealthcareRepository interface {
	CreateNewAppointment(appointment model.Appointment) error
	CreateNewVaccination(vaccination model.Vaccination) error
}
