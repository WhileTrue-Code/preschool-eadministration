package domain

import "go.mongodb.org/mongo-driver/bson"

type CrosoRepository interface {
	SaveCrosoAccount(crosoAccount *CrosoAccount) error
	UpdateCompany(company *CrosoAccount) (err error)
	FindCrosoAccountsByFounderID(founderID string) []CrosoAccount
	FindCompanyByCompanyID(company CrosoAccount) (found CrosoAccount, err error)
	SaveEmployee(request *Employee) error
	GetEmployee(filter bson.M) (employee *Employee)
	GetEmployees(filter bson.D) (employees []Employee)
	UpdateEmployee(request *Employee) (err error)
	FindEmployeesWithCompanyID(companyID string) (employees []Employee, err error)
}
