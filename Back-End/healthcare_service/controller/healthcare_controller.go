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

	router.HandleFunc("/newAppointment", controller.CreateNewAppointment).Methods("POST")
	router.HandleFunc("/allAppointments", controller.GetAllAppointments).Methods("GET")
	router.HandleFunc("/allAvailableAppointments", controller.GetAllAvailableAppointments).Methods("GET")

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

func (controller *HealthcareController) GetAllVaccinations(writer http.ResponseWriter, req *http.Request) {
	vaccinations, err := controller.service.GetAllVaccinations()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse(vaccinations, writer)
	writer.WriteHeader(http.StatusOK)
}
