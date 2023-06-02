package controller

import (
	"authorization"
	"croso_service/domain"
	"croso_service/errors"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type CrosoController struct {
	Logger  *zap.Logger
	Service domain.CrosoService
}

func NewController(crosoService domain.CrosoService, logger *zap.Logger) *CrosoController {
	return &CrosoController{
		Logger:  logger,
		Service: crosoService,
	}
}

func (controller *CrosoController) Init(router *mux.Router) {
	router.HandleFunc("/register", controller.RegisterCrosoCompany).Methods("POST")
	router.HandleFunc("/company", controller.GetMyCrosos).Methods("GET")
	router.HandleFunc("/employee/register", controller.RequestEmployeeRegistration).Methods("POST")
	router.HandleFunc("/employee/status", controller.PatchEmployeeRegistrationStatus).Methods("PATCH")
	router.HandleFunc("/employees/{companyID}", controller.GetCompanyEmployees).Methods("GET")
	http.Handle("/", router)
	controller.Logger.Info("Controller router endpoints initialized and handle run.")
}

func (controller *CrosoController) RegisterCrosoCompany(writer http.ResponseWriter, req *http.Request) {
	controller.Logger.Info("Started registering new CROSO account.")

	var request domain.RequestForCompanyRegistration
	bytes, err := io.ReadAll(req.Body)
	if err != nil {
		controller.Logger.Error("error in reading bytes of http request Body",
			zap.Error(err),
		)
		http.Error(writer, errors.ERR_BAD_REQUEST_CHECK_DATA, http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(bytes, &request)
	if err != nil {
		controller.Logger.Error("error on unmarshalling body into RequestForCompanyRegistration",
			zap.Error(err),
		)
		http.Error(writer, errors.ERR_BAD_REQUEST_CHECK_DATA, http.StatusInternalServerError)
	}

	authToken := req.Header.Get("Authorization")
	if authToken == "" {
		http.Error(writer, "authorization token doesn't exist in header", http.StatusUnauthorized)
	}
	splitted := strings.Split(authToken, " ")
	claims := authorization.GetMapClaims([]byte(splitted[1]))

	request.FounderID = claims["jmbg"]
	controller.Logger.Info("data request", zap.Any("req", request))

	err = controller.Service.RegisterCrosoAccount(&request)
	if err != nil {
		http.Error(writer, "error in saving CrosoAccount to repository.", http.StatusInternalServerError)
		return
	}

	controller.Logger.Info("CROSO account registered successfully.",
		zap.Any("request", request),
	)
	var response any = "Registracija firme na CROSO servisu uspesna."
	stringResp, _ := json.Marshal(response)
	writer.Write([]byte(stringResp))

}

func (controller *CrosoController) GetMyCrosos(writer http.ResponseWriter, req *http.Request) {
	authToken := req.Header.Get("Authorization")
	token := strings.Split(authToken, " ")[1]

	claims := authorization.GetMapClaims([]byte(token))

	founderID, ok := claims["jmbg"]
	if !ok {
		http.Error(writer, "bad request. (jmbg not presented in token)", http.StatusBadRequest)
		return
	}

	myCrosos := controller.Service.GetMyCrosos(founderID)

	responseBytes, _ := json.Marshal(myCrosos)
	writer.Write(responseBytes)
}

func (controller *CrosoController) RequestEmployeeRegistration(writer http.ResponseWriter, req *http.Request) {
	var request domain.Employee
	reqBytes, err := io.ReadAll(req.Body)
	if err != nil {
		controller.Logger.Error("error in reading request Body to bytes",
			zap.Error(err),
		)
		http.Error(writer, errors.ERR_BAD_REQUEST_CHECK_DATA, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBytes, &request)
	if err != nil {
		controller.Logger.Error("error in unmarshalling RegisterEmployeeRequest",
			zap.Error(err),
		)
		http.Error(writer, errors.ERR_BAD_REQUEST_CHECK_DATA, http.StatusBadRequest)
		return
	}

	if request.EmploymentStatus == domain.CONTRACT_DEFINITE_PERIOD && request.EmploymentDuration <= 0 {
		controller.Logger.Warn("employment status set to definite_period but not set employment duration")
		http.Error(writer, "Morate uneti trajanje ugovora u mesecima za izabrani tip ugovora", http.StatusBadRequest)
		return
	}

	err = controller.Service.RequestRegisterEmployee(&request)
	if err != nil {
		controller.Logger.Error("error in saving request.",
			zap.Error(err),
		)
		http.Error(writer, errors.ERR_SERVER_INTERNAL_MSG, http.StatusInternalServerError)
		return
	}

	var response any = "Zahtev za registraciju novog zapošljenog je podnet uspešno!"
	resBytes, _ := json.Marshal(response)
	writer.Write(resBytes)

}

func (controller *CrosoController) PatchEmployeeRegistrationStatus(writer http.ResponseWriter, req *http.Request) {
	var request domain.ResolveRequestRegisterEmployee
	reqBytes, err := io.ReadAll(req.Body)
	if err != nil {
		controller.Logger.Error("error in reading request Body to bytes",
			zap.Error(err),
		)
		http.Error(writer, errors.ERR_BAD_REQUEST_CHECK_DATA, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBytes, &request)
	if err != nil {
		controller.Logger.Error("error in unmarshalling ResolveRequestRegisterEmployee",
			zap.Error(err),
		)
		http.Error(writer, errors.ERR_BAD_REQUEST_CHECK_DATA, http.StatusBadRequest)
		return
	}

	err = controller.Service.ResolveRequestRegisterEmployee(&request)
	if err != nil {
		controller.Logger.Error("error in patching employee register status",
			zap.Error(err),
		)
		http.Error(writer, errors.ERR_SERVER_INTERNAL_MSG, http.StatusInternalServerError)
		return
	}

	var response any = "Zahtev je uspešno prihvaćen!"
	if request.Status == domain.DECLINED {
		response = "Zahtev je uspešno odbijen!"
	}

	resBytes, _ := json.Marshal(response)
	writer.Write(resBytes)

}

func (controller *CrosoController) GetCompanyEmployees(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	companyID, ok := vars["companyID"]
	if !ok {
		http.Error(writer, errors.ERR_BAD_REQUEST_CHECK_DATA, http.StatusBadRequest)
		return
	}

	employees, err := controller.Service.GetEmployeesByCompanyID(companyID)
	if err != nil {
		controller.Logger.Error("error in getting employees by companyID",
			zap.Error(err),
			zap.String("companyID", companyID),
		)
	}
	var employeesRes interface{} = employees
	bytes, _ := json.Marshal(employeesRes)
	writer.Write(bytes)

}
