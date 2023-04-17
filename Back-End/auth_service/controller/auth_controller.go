package controller

import (
	"github.com/gorilla/mux"
	"net/http"
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
	router.HandleFunc("/test", controller.Test).Methods("GET")
	http.Handle("/", router)
}

func (controller *RegistrarController) Test(writer http.ResponseWriter, req *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Okej"))
	//jsonResponse(token, writer)
}
