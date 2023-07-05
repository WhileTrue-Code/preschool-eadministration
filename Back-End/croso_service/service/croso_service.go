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
	GET_COMPANY         = "COMPANY_GET_BY_FOUNDER_COMPANY_ID"
	UPDATE_COMPANY      = "UPDATE_COMPANY"
	GET_EMPLOYEE_STATUS = "GET_EMPLOYEE_STATUS_BY_ID"
	CHECK_USER_JMBG     = "CHECK_USER_JMBG"
)

func NewAprService(crosoRepo domain.CrosoRepository, nats *nats.Conn, logger *zap.Logger) domain.CrosoService {
	return &CrosoServiceImpl{
		Repo:   crosoRepo,
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

	if account.CompanyID == 0 {
		service.Logger.Warn("company not found in APR db.", zap.String("companyID", request.CompanyID))
		return fmt.Errorf(errors.ERR_RS_COMPANY_NOT_EXIST_IN_APR)
	}

	service.Logger.Info("id of struct responsed from nats", zap.String("idPrimitive", account.ID.Hex()))

	return service.Repo.SaveCrosoAccount(&account)
}

func (service *CrosoServiceImpl) GetMyCrosos(founderID string) []domain.CrosoAccount {

	return service.Repo.FindCrosoAccountsByFounderID(founderID)
}

func (service *CrosoServiceImpl) RequestRegisterEmployee(employee *domain.Employee) (err error) {

	employeeId := map[string]string{
		"jmbg": employee.EmployeeID,
	}
	bytes, _ := json.Marshal(employeeId)

	msg, err := service.Nats.Request(os.Getenv(CHECK_USER_JMBG), bytes, 5*time.Second)
	if err != nil {
		service.Logger.Error("unable to get msg from nats.",
			zap.Error(err),
		)
		return
	}
	var doesExist bool
	json.Unmarshal(msg.Data, &doesExist)
	service.Logger.Info("got doesExist from nats",
		zap.Any("jmbg", employeeId),
		zap.Bool("doesExist", doesExist),
	)

	if !doesExist {
		return fmt.Errorf(errors.ERR_RS_USER_NOT_EXIST)
	}

	found := service.Repo.GetEmployee(bson.M{"employeeID": employee.EmployeeID})
	if found != nil {
		if found.EmploymentStatus != domain.UNEMPLOYED {
			return fmt.Errorf(errors.ERR_RS_EMPLOYEE_ALREADY_EMPLOYEED)
		}
		employee.ID = found.ID
		employee.RegistrationStatus = domain.PENDING
		employee.RegistrationTimestamp = time.Now().Unix()
		calculateTaxesAndContributions(found)
		return service.Repo.UpdateEmployee(found)
	}

	employee.ID = primitive.NewObjectID()
	employee.RegistrationStatus = domain.PENDING
	employee.RegistrationTimestamp = time.Now().Unix()
	calculateTaxesAndContributions(employee)
	return service.Repo.SaveEmployee(employee)
}

func (service *CrosoServiceImpl) GetPendingEmployeeRequests() (pending []domain.Employee) {
	return service.Repo.GetEmployees(bson.D{{Key: "registrationStatus", Value: domain.PENDING}})
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

func (service *CrosoServiceImpl) GetEmployeeByIDCardID(id string) (employee *domain.Employee, err error) {
	filter := bson.M{"employeeID": id}
	return service.Repo.GetEmployee(filter), nil
}

func (service *CrosoServiceImpl) SubscribeToNats(connection *nats.Conn) {
	_, err := connection.QueueSubscribe(os.Getenv(GET_EMPLOYEE_STATUS), GET_EMPLOYEE_STATUS,
		func(message *nats.Msg) {
			log := service.Logger.Named("[NATS/GET_EMPLOYEE_STATUS]")
			var request map[string]string
			err := json.Unmarshal(message.Data, &request)
			if err != nil {
				log.Error("error in unmarshalling json")
				return
			}

			employeeID, ok := request["employeeID"]
			if !ok {
				log.Error("bad request got from nats.")
				return
			}

			employee, err := service.GetEmployeeByIDCardID(employeeID)
			if err != nil {
				if err.Error() != errors.ERR_EMPLOYEE_NOT_FOUND {
					log.Error("error in finding company",
						zap.Error(err),
						zap.String("employeeID", employeeID),
					)
					return
				}

			}

			response := map[string]bool{
				"employed": false,
			}
			if employee != nil && employee.EmploymentStatus != domain.UNEMPLOYED {
				response["employed"] = true
			}

			responseBytes, err := json.Marshal(response)
			if err != nil {
				log.Error("error in marshalling json response",
					zap.Error(err),
				)
				return
			}

			err = connection.Publish(message.Reply, responseBytes)
			if err != nil {
				log.Error("error in publishing response",
					zap.Error(err),
				)
				return
			}

		},
	)
	if err != nil {
		service.Logger.Error("error in subscribing to NATS queue",
			zap.Error(err),
		)
		return
	}

	_, err = connection.QueueSubscribe(os.Getenv(UPDATE_COMPANY), UPDATE_COMPANY, func(message *nats.Msg) {
		var request domain.CrosoAccount
		err := json.Unmarshal(message.Data, &request)
		if err != nil {
			service.Logger.Error("error in unmarshalling GetCompany struct")
			return
		}

		found, err := service.Repo.FindCompanyByCompanyID(request)
		if err != nil {
			return
		}

		request.ID = found.ID

		updated := true
		err = service.Repo.UpdateCompany(&request)
		if err != nil {
			service.Logger.Error("error in finding company",
				zap.Error(err),
				zap.Any("company", request),
			)
			updated = false
		}

		response, err := json.Marshal(updated)
		if err != nil {
			service.Logger.Error("error in marshalling Company struct",
				zap.Error(err),
			)
			return
		}

		err = connection.Publish(message.Reply, response)
		if err != nil {
			service.Logger.Error("error in publishing response",
				zap.Error(err),
			)
			return
		}

		service.Logger.Info("PUBLISHED MESSAGE")

	})
	if err != nil {
		service.Logger.Error("error in subscribing to NATS queue",
			zap.Error(err),
		)
		return
	}

	service.Logger.Sugar().Infof("Subscribed to channels")
}

func calculateTaxesAndContributions(employee *domain.Employee) {
	employee.PersonalIncomeTax = calculatePersonalIncomeTax(employee.NetSalary)
	employee.PDContribution = calculatePensionDisabilityContribution(employee.NetSalary)
	employee.HIContribution = calculateHealthInsuranceContribution(employee.NetSalary)
	employee.UIContribution = calculateUnemploymentInsuranceContribution(employee.NetSalary)
	employee.EFContribution = calculateEmploymentFundContribution(employee.NetSalary)
	employee.GrossPay = float64(employee.NetSalary) + employee.PersonalIncomeTax +
		employee.PDContribution + employee.HIContribution + employee.UIContribution +
		employee.EFContribution
}

func calculatePersonalIncomeTax(netSalary int) (personalIncomeTax float64) {
	var taxPercent float64
	if 3601 >= netSalary && netSalary <= 15000 {
		taxPercent = 0.15
	} else if 15001 >= netSalary && netSalary <= 30000 {
		taxPercent = 0.2
	} else {
		taxPercent = 0.1
	}

	personalIncomeTax = float64(netSalary) * taxPercent
	return
}

func calculatePensionDisabilityContribution(netSalary int) (pdContribution float64) {
	return float64(netSalary) * 0.26
}

func calculateHealthInsuranceContribution(netSalary int) (hiContribution float64) {
	return float64(netSalary) * 0.123
}

func calculateUnemploymentInsuranceContribution(netSalary int) (uiContribution float64) {
	return float64(netSalary) * 0.0075
}

func calculateEmploymentFundContribution(netSalary int) (efContribution float64) {
	return float64(netSalary) * 0.0075
}

func (service *CrosoServiceImpl) ChangeEmploymentStatus(id string, changeRequest domain.ChangeEmploymentStatus) (err error) {

	primitiveID, _ := primitive.ObjectIDFromHex(id)

	employee := service.Repo.GetEmployee(bson.M{"_id": primitiveID})
	if employee == nil {
		err = fmt.Errorf(errors.ERR_EMPLOYEE_NOT_FOUND)
		return
	}

	employee.EmploymentStatus = changeRequest.EmploymentStatus
	employee.EmploymentDuration = changeRequest.EmploymentDuration

	err = service.Repo.UpdateEmployee(employee)

	return
}

func (service *CrosoServiceImpl) CancelEmployment(id string) (err error) {

	primitiveID, _ := primitive.ObjectIDFromHex(id)

	employee := service.Repo.GetEmployee(bson.M{"_id": primitiveID})
	if employee == nil {
		err = fmt.Errorf(errors.ERR_EMPLOYEE_NOT_FOUND)
		return
	}

	employee.EmploymentStatus = domain.UNEMPLOYED
	employee.CompanyID = 0
	employee.EmploymentDuration = 0

	err = service.Repo.UpdateEmployee(employee)

	return
}
