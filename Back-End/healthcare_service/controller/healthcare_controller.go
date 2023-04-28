package controller

import (
	"authorization"
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

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8005", authorization.Authorizer(authEnforcer)(router)))
}

func (controller *HealthcareController) CreateNewAppointment(writer http.ResponseWriter, req *http.Request) {
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

	var appointment model.Appointment
	err = controller.service.CreateNewAppointment(appointment, jmbg)
	if err != nil {
		return
	}
}
