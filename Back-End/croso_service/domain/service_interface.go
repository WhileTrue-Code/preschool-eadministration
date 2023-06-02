package domain

type CrosoService interface {
	RegisterCrosoAccount(request *RequestForCompanyRegistration) error
	GetMyCrosos(founderID string) []CrosoAccount
	RequestRegisterEmployee(request *Employee) error
	ResolveRequestRegisterEmployee(request *ResolveRequestRegisterEmployee) error
	GetEmployeesByCompanyID(companyID string) (employees []Employee, err error)
}
