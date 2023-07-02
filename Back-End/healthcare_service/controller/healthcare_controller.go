package controller

import (
	"authorization"
	"encoding/json"
	"github.com/casbin/casbin"
	"github.com/cristalhq/jwt/v4"
	"github.com/gorilla/mux"
	"healthcare_service/model"
	"healthcare_service/service"
	"log"
	"net/http"
	"os"
)

type HealthcareController struct {
	service *service.HealthcareService
}

var jwtKey = []byte(os.Getenv("SECRET_KEY"))
var verifier, _ = jwt.NewVerifierHS(jwt.HS256, jwtKey)

func NewHealthcareController(service *service.HealthcareService) *HealthcareController {
	return &HealthcareController{
		service: service,
	}
}

func (controller *HealthcareController) Init(router *mux.Router) {
	authEnforcer, err := casbin.NewEnforcerSafe("./auth_model.conf", "./policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/allAppointments", controller.GetAllAppointments).Methods("GET")
	router.HandleFunc("/myAppointmentsDoctor", controller.GetAllMyAppointmentsDoctor).Methods("GET")
	router.HandleFunc("/myAvailableAppointmentsDoctor", controller.GetMyAvailableAppointmentsDoctor).Methods("GET")
	router.HandleFunc("/myTakenAppointmentsDoctor", controller.GetMyTakenAppointmentsDoctor).Methods("GET")
	router.HandleFunc("/allAvailableAppointments", controller.GetAllAvailableAppointments).Methods("GET")
	router.HandleFunc("/getAppointmentByID/{id}", controller.GetAppointmentByID).Methods("GET")
	router.HandleFunc("/newAppointment", controller.CreateNewAppointment).Methods("POST")
	router.HandleFunc("/setAppointment/{id}", controller.SetAppointment).Methods("PUT")
	router.HandleFunc("/deleteAppointmentByID/{id}", controller.DeleteAppointmentByID).Methods("DELETE")

	router.HandleFunc("/allVaccinations", controller.GetAllVaccinations).Methods("GET")
	router.HandleFunc("/myVaccinationsDoctor", controller.GetAllMyVaccinationsDoctor).Methods("GET")
	router.HandleFunc("/myAvailableVaccinationsDoctor", controller.GetMyAvailableVaccinationsDoctor).Methods("GET")
	router.HandleFunc("/myTakenVaccinationsDoctor", controller.GetMyTakenVaccinationsDoctor).Methods("GET")
	router.HandleFunc("/allAvailableVaccinations", controller.GetAllAvailableVaccinations).Methods("GET")
	router.HandleFunc("/myTakenVaccinationsRegular", controller.GetMyTakenVaccinationsRegular).Methods("GET")
	router.HandleFunc("/getVaccinationByID/{id}", controller.GetVaccinationByID).Methods("GET")
	router.HandleFunc("/newVaccination", controller.CreateNewVaccination).Methods("POST")
	router.HandleFunc("/setVaccination/{id}", controller.SetVaccination).Methods("PUT")
	router.HandleFunc("/deleteVaccinationByID/{id}", controller.DeleteVaccinationByID).Methods("DELETE")

	router.HandleFunc("/allZdravstvenaStanja", controller.GetAllZdravstvenaStanja).Methods("GET")
	router.HandleFunc("/getZdravstvenoStanjeByID/{id}", controller.GetZdravstvenoStanjeByID).Methods("GET")
	router.HandleFunc("/getZdravstvenoStanjeByJMBG/{jmbg}", controller.GetZdravstvenoStanjeByJMBG).Methods("GET")
	router.HandleFunc("/myZdravstvenoStanje", controller.GetMyZdravstvenoStanje).Methods("GET")
	router.HandleFunc("/newZdravstvenoStanje", controller.CreateNewZdravstvenoStanje).Methods("POST")
	router.HandleFunc("/deleteZdravstvenoStanjeByJMBG/{jmbg}", controller.DeleteZdravstvenoStanjeByJMBG).Methods("DELETE")

	router.HandleFunc("/addPersonToRegistry", controller.AddPersonToRegistry).Methods("POST")
	router.HandleFunc("/getMe", controller.GetMe).Methods("GET")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8005", authorization.Authorizer(authEnforcer)(router)))
}

func (controller *HealthcareController) GetAllAppointments(writer http.ResponseWriter, req *http.Request) {
	appointments, err := controller.service.GetAllAppointments()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	jsonResponse(appointments, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) GetAllMyAppointmentsDoctor(writer http.ResponseWriter, req *http.Request) {
	jmbg, err := extractJMBGFromClaims(writer, req)

	appointments, err := controller.service.GetMyAppointmentsDoctor(jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	jsonResponse(appointments, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) GetMyAvailableAppointmentsDoctor(writer http.ResponseWriter, req *http.Request) {
	jmbg, err := extractJMBGFromClaims(writer, req)

	appointments, err := controller.service.GetMyAvailableAppointmentsDoctor(jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	jsonResponse(appointments, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) GetMyTakenAppointmentsDoctor(writer http.ResponseWriter, req *http.Request) {
	jmbg, err := extractJMBGFromClaims(writer, req)

	appointments, err := controller.service.GetMyTakenAppointmentsDoctor(jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	jsonResponse(appointments, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) GetAllAvailableAppointments(writer http.ResponseWriter, req *http.Request) {
	appointments, err := controller.service.GetAllAvailableAppointments()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse(appointments, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) GetAppointmentByID(writer http.ResponseWriter, req *http.Request) {
	objectID, err := getIDFromReqAsPrimitive(writer, req)

	appointment, err := controller.service.GetAppointmentByID(objectID)
	if err != nil {
		log.Println("Error finding Appointment By ID")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonResponse(appointment, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) CreateNewAppointment(writer http.ResponseWriter, req *http.Request) {
	var appointment model.Appointment
	err := json.NewDecoder(req.Body).Decode(&appointment)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("There is a problem in decoding JSON"))
		return
	}

	jmbg, err := extractJMBGFromClaims(writer, req)

	value, err := controller.service.CreateNewAppointment(&appointment, jmbg)
	if value == 1 {
		writer.WriteHeader(http.StatusNotAcceptable)
		writer.Write([]byte("Appointment already exists in that time"))
		return
	}
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse(appointment, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) SetAppointment(writer http.ResponseWriter, req *http.Request) {
	objectID, err := getIDFromReqAsPrimitive(writer, req)
	jmbg, err := extractJMBGFromClaims(writer, req)

	appointment, err := controller.service.SetAppointment(objectID, jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	jsonResponse(appointment, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) DeleteAppointmentByID(writer http.ResponseWriter, req *http.Request) {
	objectID, err := getIDFromReqAsPrimitive(writer, req)

	err = controller.service.DeleteAppointmentByID(objectID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) GetAllVaccinations(writer http.ResponseWriter, req *http.Request) {
	vaccinations, err := controller.service.GetAllVaccinations()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse(vaccinations, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) GetAllMyVaccinationsDoctor(writer http.ResponseWriter, req *http.Request) {
	jmbg, err := extractJMBGFromClaims(writer, req)

	vaccinations, err := controller.service.GetMyVaccinationsDoctor(jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	jsonResponse(vaccinations, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) GetMyAvailableVaccinationsDoctor(writer http.ResponseWriter, req *http.Request) {
	jmbg, err := extractJMBGFromClaims(writer, req)

	vaccinations, err := controller.service.GetMyAvailableVaccinationsDoctor(jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	jsonResponse(vaccinations, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) GetMyTakenVaccinationsDoctor(writer http.ResponseWriter, req *http.Request) {
	jmbg, err := extractJMBGFromClaims(writer, req)

	vaccinations, err := controller.service.GetMyTakenVaccinationsDoctor(jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	jsonResponse(vaccinations, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) GetAllAvailableVaccinations(writer http.ResponseWriter, req *http.Request) {
	vaccinations, err := controller.service.GetAllAvailableVaccinations()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse(vaccinations, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) GetMyTakenVaccinationsRegular(writer http.ResponseWriter, req *http.Request) {
	jmbg, err := extractJMBGFromClaims(writer, req)

	vaccinations, err := controller.service.GetMyTakenVaccinationsRegular(jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	jsonResponse(vaccinations, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) GetVaccinationByID(writer http.ResponseWriter, req *http.Request) {
	objectID, err := getIDFromReqAsPrimitive(writer, req)

	vaccination, err := controller.service.GetVaccinationByID(objectID)
	if err != nil {
		log.Println("Error finding Appointment By ID")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonResponse(vaccination, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) CreateNewVaccination(writer http.ResponseWriter, req *http.Request) {
	var vaccination model.Vaccination
	err := json.NewDecoder(req.Body).Decode(&vaccination)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("There is a problem in decoding JSON"))
		return
	}

	jmbg, err := extractJMBGFromClaims(writer, req)

	value, err := controller.service.CreateNewVaccination(&vaccination, jmbg)
	if value == 1 {
		writer.WriteHeader(http.StatusNotAcceptable)
		writer.Write([]byte("Vaccination already exists in that time"))
		return
	}
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse(vaccination, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) SetVaccination(writer http.ResponseWriter, req *http.Request) {
	objectID, err := getIDFromReqAsPrimitive(writer, req)

	jmbg, err := extractJMBGFromClaims(writer, req)

	vaccination, err := controller.service.SetVaccination(objectID, jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	jsonResponse(vaccination, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) DeleteVaccinationByID(writer http.ResponseWriter, req *http.Request) {
	objectID, err := getIDFromReqAsPrimitive(writer, req)

	err = controller.service.DeleteVaccinationByID(objectID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) GetAllZdravstvenaStanja(writer http.ResponseWriter, req *http.Request) {
	zdravstvenaStanja, err := controller.service.GetAllZdravstvenoStanje()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse(zdravstvenaStanja, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) GetZdravstvenoStanjeByID(writer http.ResponseWriter, req *http.Request) {
	objectID, err := getIDFromReqAsPrimitive(writer, req)

	zdravstvenoStanje, err := controller.service.GetZdravstvenoStanjeByID(objectID)
	if err != nil {
		log.Println("Error finding Zdravstveno Stanje By ID")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonResponse(zdravstvenoStanje, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) GetZdravstvenoStanjeByJMBG(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jmbg := vars["jmbg"]

	zdravstvenoStanje, err := controller.service.GetZdravstvenoStanjeByJMBG(jmbg)
	if err != nil {
		log.Println("Error finding Zdravstveno Stanje By JMBG")
		log.Println("Found no Zdravstveno Stanje with that JMBG")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonResponse(zdravstvenoStanje, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) GetMyZdravstvenoStanje(writer http.ResponseWriter, req *http.Request) {
	jmbg, err := extractJMBGFromClaims(writer, req)

	zdravstvenoStanje, err := controller.service.GetMyZdravstvenoStanje(jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	jsonResponse(zdravstvenoStanje, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) CreateNewZdravstvenoStanje(writer http.ResponseWriter, req *http.Request) {
	var zdravstvenoStanje model.ZdravstvenoStanje
	err := json.NewDecoder(req.Body).Decode(&zdravstvenoStanje)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("There is a problem in decoding JSON"))
		return
	}

	//Validacije na frontu, ne moram bas i ovde
	//if zdravstvenoStanje.Jmbg == "" {
	//	writer.WriteHeader(http.StatusBadRequest)
	//	writer.Write([]byte("JMBG is required"))
	//	return
	//}

	existingZdravstveno, _ := controller.service.GetZdravstvenoStanjeByJMBG(zdravstvenoStanje.Jmbg)
	if existingZdravstveno != nil {
		writer.WriteHeader(http.StatusAccepted)
		writer.Write([]byte("Zdravstveno Stanje with that ID already exists"))
		return
	}

	err = controller.service.CreateNewZdravstvenoStanje(&zdravstvenoStanje)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse(zdravstvenoStanje, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) DeleteZdravstvenoStanjeByJMBG(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jmbg := vars["jmbg"]

	err := controller.service.DeleteZdravstvenoStanjeByJMBG(jmbg)
	if err != nil {
		log.Println("Error in deleting Zdravstveno Stanje by JMBG")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) GetMe(writer http.ResponseWriter, req *http.Request) {
	jmbg, err := extractJMBGFromClaims(writer, req)

	user, err := controller.service.GetMe(jmbg)
	if err != nil {
		log.Println("Error getting User")
	}

	jsonResponse(user, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) AddPersonToRegistry(writer http.ResponseWriter, req *http.Request) {
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("There is a problem in decoding JSON"))
		return
	}

	newUser, value := controller.service.AddPersonToRegistry(&user)
	if value == 1 {
		writer.WriteHeader(http.StatusAccepted)
		writer.Write([]byte("User already exists in database"))
		return
	}

	jsonResponse(newUser, writer)
	writer.WriteHeader(http.StatusOK)
}
