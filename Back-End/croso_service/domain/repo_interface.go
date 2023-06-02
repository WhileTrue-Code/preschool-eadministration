package domain

import "go.mongodb.org/mongo-driver/bson"

type CrosoRepository interface {
	SaveCrosoAccount(crosoAccount *CrosoAccount) error
	FindCrosoAccountsByFounderID(founderID string) []CrosoAccount
	SaveEmployee(request *Employee) error
	GetEmployee(filter bson.M) (employee *Employee)
	GetEmployees(filter bson.D) (employees []Employee)
	UpdateEmployee(request *Employee) (err error)
	FindEmployeesWithCompanyID(companyID string) (employees []Employee, err error)
}
