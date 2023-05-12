package controller

import (
	"authorization"
	"encoding/json"
	"github.com/casbin/casbin"
	"github.com/cristalhq/jwt/v4"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	router.HandleFunc("/getVaccinationByID/{id}", controller.GetVaccinationByID).Methods("GET")
	router.HandleFunc("/newVaccination", controller.CreateNewVaccination).Methods("POST")
	router.HandleFunc("/setVaccination/{id}", controller.SetVaccination).Methods("PUT")
	router.HandleFunc("/deleteVaccinationByID/{id}", controller.DeleteVaccinationByID).Methods("DELETE")

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
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		log.Println("Get ID from req error")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Convert to Primitive error")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	appointment, err := controller.service.GetAppointmentByID(objectID)
	if err != nil {
		log.Println("Error finding Appointment By ID")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonResponse(appointment, writer)
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

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Added"))
}

func (controller *HealthcareController) SetAppointment(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		log.Println("Get ID from req error")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	jmbg, err := extractJMBGFromClaims(writer, req)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Convert to Primitive error")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = controller.service.SetAppointment(objectID, jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Updated"))
}

func (controller *HealthcareController) DeleteAppointmentByID(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		log.Println("Get ID from req error")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Convert to Primitive error")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.service.DeleteAppointmentByID(objectID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Deleted"))
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
	appointments, err := controller.service.GetAllAvailableAppointments()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse(appointments, writer)
	writer.WriteHeader(http.StatusOK)
}

func (controller *HealthcareController) GetVaccinationByID(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		log.Println("Get ID from req error")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Convert to Primitive error")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

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

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Added"))
}

func (controller *HealthcareController) SetVaccination(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		log.Println("Get ID from req error")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	jmbg, err := extractJMBGFromClaims(writer, req)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Convert to Primitive error")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = controller.service.SetVaccination(objectID, jmbg)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Updated"))
}

func (controller *HealthcareController) DeleteVaccinationByID(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		log.Println("Get ID from req error")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Convert to Primitive error")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = controller.service.DeleteVaccinationByID(objectID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Deleted"))
}
