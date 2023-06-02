package domain

import "github.com/nats-io/nats.go"

type CrosoService interface {
	RegisterCrosoAccount(request *RequestForCompanyRegistration) error
	RequestRegisterEmployee(request *Employee) error
	ResolveRequestRegisterEmployee(request *ResolveRequestRegisterEmployee) error
	GetEmployeesByCompanyID(companyID string) (employees []Employee, err error)
	SubscribeToNats(connection *nats.Conn)
}
