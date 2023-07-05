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
	EmploymentDuration int8 `json:"employmentDuration" bson:"employmentDuration"`
	NetSalary          int  `json:"netSalary" bson:"netSalary"`

	// GrossPay is sum of NetSalary and all contributions and taxes which have to be payed to state of RS
	GrossPay float64 `json:"grossPay" bson:"grossPay"`

	// PersonalIncomeTax is tax which has to be payed to the state of RS
	// to 3600 RSD PIT amounts to 10 percent of net salary amount
	// from 3601 to 15000 PIT 15 percent
	// from 15001 to 30000 PIT 20 percent
	// from 30000 and higher PIT 10 percent
	PersonalIncomeTax float64 `json:"personalIncomeTax" bson:"personalIncomeTax"`

	// PDContribution stands for Pension and Disability contribution and presented by 26% of employee net salary
	PDContribution float64 `json:"pdContribution" bson:"pdContribution"`

	// HIContribution stands for Health Insurance Contribution and presented by 12,3% of employee net salary
	HIContribution float64 `json:"hiContribution" bson:"hiContribution"`

	// UIContribution stands for Unemployment Insurance Contribution and presented by 0,075% of employee net salary
	UIContribution float64 `json:"uiContribution" bson:"uiContribution"`

	// EFContribution stands for Employment Fund Contribution and presented by 0,075% of employee net salary
	EFContribution float64 `json:"efContribution" bson:"efContribution"`

	RegistrationStatus RegistrationStatus `json:"registrationStatus" bson:"registrationStatus"`
}

type EmploymentStatus string

const (
	UNEMPLOYED                 = "unemployed"
	CONTRACT_DEFINITE_PERIOD   = "definite_contract"
	CONTRACT_INDEFINITE_PERIOD = "indefinite_contract"
	CONTRACT_TEMPORARY_WORKS   = "temporary_works"
)

type ChangeEmploymentStatus struct {
	EmploymentStatus   EmploymentStatus `json:"employmentStatus"`
	EmploymentDuration int8             `json:"employmentDuration"`
}

type ResolveRequestRegisterEmployee struct {
	EmployeeID string             `json:"employeeID"`
	CompanyID  int                `json:"companyID"`
	Status     RegistrationStatus `json:"status"`
}

type RegistrationStatus int

const (
	PENDING RegistrationStatus = iota + 1
	ACCEPTED
	DECLINED
)
