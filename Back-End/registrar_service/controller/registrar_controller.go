package controller

import (
	"github.com/casbin/casbin"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"registrar_service/authorization"
	"registrar_service/service"
)

type RegistrarController struct {
	service *service.RegistrarService
}

func NewRegistrarController(service *service.RegistrarService) *RegistrarController {
	return &RegistrarController{
		service: service,
	}
}

func (controller *RegistrarController) Init(router *mux.Router) {

	authEnforcer, err := casbin.NewEnforcerSafe("./auth_model.conf", "./policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/vencani", controller.Test).Methods("POST")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8001", authorization.Authorizer(authEnforcer)(router)))

}

func (controller *RegistrarController) Test(writer http.ResponseWriter, req *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Okej"))
	//jsonResponse(token, writer)
}
