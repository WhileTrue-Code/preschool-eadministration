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
	"strings"
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
	router.HandleFunc("/allAvailableAppointments", controller.GetAllAvailableAppointments).Methods("GET")
	router.HandleFunc("/getAppointmentByID/{id}", controller.GetAppointmentByID).Methods("GET")
	router.HandleFunc("/newAppointment", controller.CreateNewAppointment).Methods("POST")
	router.HandleFunc("/setAppointment/{id}", controller.SetAppointment).Methods("PUT")
	router.HandleFunc("/deleteAppointmentByID/{id}", controller.DeleteAppointmentByID).Methods("DELETE")

	router.HandleFunc("/allVaccinations", controller.GetAllVaccinations).Methods("GET")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8005", authorization.Authorizer(authEnforcer)(router)))
}

func (controller *HealthcareController) CreateNewAppointment(writer http.ResponseWriter, req *http.Request) {
	var appointment model.Appointment
	err := json.NewDecoder(req.Body).Decode(&appointment)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("There is a problem in decoding JSON"))
		return
	}

	bearer := req.Header.Get("Authorization")
	bearerToken := strings.Split(bearer, "Bearer ")
	tokenString := bearerToken[1]

	token, err := jwt.Parse([]byte(tokenString), verifier)

	if err != nil {
		log.Println(err)
		http.Error(writer, "unauthorized", http.StatusUnauthorized)
		return
	}

	claims := authorization.GetMapClaims(token.Bytes())
	jmbg := claims["jmbg"]

	value, err := controller.service.CreateNewAppointment(appointment, jmbg)
	if value == 1 {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error in Service"))
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

	bearer := req.Header.Get("Authorization")
	bearerToken := strings.Split(bearer, "Bearer ")
	tokenString := bearerToken[1]

	token, err := jwt.Parse([]byte(tokenString), verifier)

	if err != nil {
		log.Println(err)
		http.Error(writer, "unauthorized", http.StatusUnauthorized)
		return
	}

	claims := authorization.GetMapClaims(token.Bytes())
	jmbg := claims["jmbg"]

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

func (controller *HealthcareController) GetAllAppointments(writer http.ResponseWriter, req *http.Request) {
	appointments, err := controller.service.GetAllAppointments()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
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
