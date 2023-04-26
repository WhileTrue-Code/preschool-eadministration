package controller

import (
	"authorization"
	"github.com/casbin/casbin"
	"github.com/gorilla/mux"
	"healthcare_service/service"
	"log"
	"net/http"
)

type HealthcareController struct {
	service *service.HealthcareService
}

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

	////router.HandleFunc("/getAll", controller.GetAll).Methods("GET")
	//router.HandleFunc("/registration", controller.SignUp).Methods("POST")
	//router.HandleFunc("/login", controller.Login).Methods("POST")
	//router.HandleFunc("/test", controller.Test).Methods("GET")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8005", authorization.Authorizer(authEnforcer)(router)))
}