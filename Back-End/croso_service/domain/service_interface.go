package domain

import "github.com/nats-io/nats.go"

type CrosoService interface {
	RegisterCrosoAccount(request *RequestForCompanyRegistration) error
	GetMyCrosos(founderID string) []CrosoAccount
	RequestRegisterEmployee(request *Employee) error
	GetPendingEmployeeRequests() (pending []Employee)
	ResolveRequestRegisterEmployee(request *ResolveRequestRegisterEmployee) error
	GetEmployeesByCompanyID(companyID string) (employees []Employee, err error)
	ChangeEmploymentStatus(id string, changeRequest ChangeEmploymentStatus) (err error)
	CancelEmployment(id string) (err error)
	SubscribeToNats(connection *nats.Conn)
}
