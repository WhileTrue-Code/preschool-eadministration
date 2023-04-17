package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"registrar_service/service"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (controller *AuthController) Init(router *mux.Router) {
	router.HandleFunc("/test", controller.Test).Methods("GET")
	http.Handle("/", router)
}

func (controller *AuthController) Test(writer http.ResponseWriter, req *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Okej"))
	//jsonResponse(token, writer)
}
