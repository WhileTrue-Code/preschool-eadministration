package domain

import (
	"model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CrosoAccount model.CompanyAccount

type RequestForCompanyRegistration struct {
	FounderID string `json:"founderID"`
	CompanyID string `json:"companyID"`
}

type Employee struct {
	ID                    primitive.ObjectID `json:"id" bson:"_id"`
	FirstName             string             `json:"firstName" bson:"firstName"`
	LastName              string             `json:"lastName" bson:"lastName"`
	Address               string             `json:"address" bson:"address"`
	EmployeeID            string             `json:"employeeID" bson:"employeeID"`
	IDCardNumber          string             `json:"idCardNumber" bson:"idCardNumber"`
	PassportNumber        string             `json:"passportNumber" bson:"passportNumber"`
	HealthCardNumber      string             `json:"healthCardNumber" bson:"healthCardNumber"`
	CompanyID             int                `json:"companyID" bson:"companyID"`
	RegistrationTimestamp int64              `json:"timestamp" bson:"timestamp"`
	EmploymentStatus      EmploymentStatus   `json:"employmentStatus" bson:"employmentStatus"`
	// EmploymentDuration represents employment duration in number
	// of months if is EmloymentStatus set to CONTRACT_DEFINITE_PERIOD
	EmploymentDuration int8               `json:"employmentDuration" bson:"employmentDuration"`
	RegistrationStatus RegistrationStatus `json:"registrationStatus" bson:"registrationStatus"`
}

type EmploymentStatus string

const (
	UNEMPLOYED                 = "unemployed"
	CONTRACT_DEFINITE_PERIOD   = "definite_contract"
	CONTRACT_INDEFINITE_PERIOD = "indefinite_contract"
	CONTRACT_TEMPORARY_WORKS   = "temporary_works"
)

type ResolveRequestRegisterEmployee struct {
	EmployeeID string
	CompanyID  int
	Status     RegistrationStatus
}

type RegistrationStatus int

const (
	PENDING RegistrationStatus = iota + 1
	ACCEPTED
	DECLINED
)
