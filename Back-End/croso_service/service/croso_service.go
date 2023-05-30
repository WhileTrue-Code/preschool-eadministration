package service

import (
	"croso_service/domain"
	"croso_service/errors"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type CrosoServiceImpl struct {
	Repo   domain.CrosoRepository
	Nats   *nats.Conn
	Logger *zap.Logger
}

var (
	GET_COMPANY = "COMPANY_GET_BY_FOUNDER_COMPANY_ID"
)

func NewAprService(aprRepo domain.CrosoRepository, nats *nats.Conn, logger *zap.Logger) domain.CrosoService {
	return &CrosoServiceImpl{
		Repo:   aprRepo,
		Nats:   nats,
		Logger: logger,
	}
}

func (service *CrosoServiceImpl) RegisterCrosoAccount(request *domain.RequestForCompanyRegistration) (err error) {
	// logic from NATS to get company account from APR service

	requestBytes, _ := json.Marshal(*request)

	msg, err := service.Nats.Request(os.Getenv(GET_COMPANY), []byte(requestBytes), 5*time.Second)
	if err != nil {
		service.Logger.Error("an error occured on requesting data on NATS.",
			zap.Any("requestData", *request),
			zap.Error(err),
		)
		return
	}

	var account domain.CrosoAccount
	err = json.Unmarshal(msg.Data, &account)
	if err != nil {
		service.Logger.Error("error in unmarshal data responded from NATS.",
			zap.Error(err),
		)
		return
	}

	return service.Repo.SaveCrosoAccount(&account)
}

func (service *CrosoServiceImpl) RequestRegisterEmployee(employee *domain.Employee) error {
	employee.ID = primitive.NewObjectID()
	employee.RegistrationTimestamp = time.Now().Unix()
	return service.Repo.SaveEmployee(employee)
}

func (service *CrosoServiceImpl) ResolveRequestRegisterEmployee(request *domain.ResolveRequestRegisterEmployee) (err error) {

	employee := service.Repo.GetEmployee(bson.M{
		"companyID":  request.CompanyID,
		"employeeID": request.EmployeeID,
	})

	if employee == nil {
		return fmt.Errorf(errors.ERR_EMPLOYEE_NOT_FOUND)
	}

	employee.RegistrationStatus = request.Status

	err = service.Repo.UpdateEmployee(employee)

	return
}

func (service *CrosoServiceImpl) GetEmployeesByCompanyID(companyID string) (employees []domain.Employee, err error) {
	return service.Repo.FindEmployeesWithCompanyID(companyID)
}
