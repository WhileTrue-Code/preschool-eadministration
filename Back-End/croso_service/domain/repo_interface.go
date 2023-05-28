package domain

import "go.mongodb.org/mongo-driver/bson"

type CrosoRepository interface {
	SaveCrosoAccount(crosoAccount *CrosoAccount) error
	FindCrosoAccountsByFounderID(founderID string) ([]CrosoAccount, error)
	SaveEmployee(request *Employee) error
	GetEmployee(filter bson.M) (employee *Employee)
	UpdateEmployee(request *Employee) (err error)
}
