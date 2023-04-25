package controller

import (
	"encoding/json"
	"github.com/casbin/casbin"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"registrar_service/authorization"
	"registrar_service/model/entity"
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

	router.HandleFunc("/registry", controller.CreateNewBirthCertificate).Methods("POST")
	router.HandleFunc("/test", controller.Test).Methods("GET")
	router.HandleFunc("/marriage", controller.Marriage).Methods("POST")
	router.HandleFunc("/died", controller.UpdateCertificate).Methods("POST")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8001", authorization.Authorizer(authEnforcer)(router)))

}

func (controller *RegistrarController) CreateNewBirthCertificate(writer http.ResponseWriter, req *http.Request) {

	log.Println("Hello Birth")

	var user entity.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Problem to parsing JSON to entity!"))
		return
	}

	log.Println(user)
	value, err := controller.service.CreateNewBirthCertificate(user)
	if value == 1 {
		writer.WriteHeader(http.StatusAccepted)
		writer.Write([]byte("JMBG already exist in system!"))
		return
	}
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Okej"))
}

func (controller *RegistrarController) Marriage(writer http.ResponseWriter, req *http.Request) {

	var marriage entity.ExcerptFromTheMarriageRegister
	err := json.NewDecoder(req.Body).Decode(&marriage)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Problem to parsing JSON to entity!"))
		return
	}

	//find Svedok1 i Svedok2
	var svedok1 *entity.User
	var svedok2 *entity.User

	svedok1 = controller.service.FindOneUser(marriage.Svedok1.JMBG)
	svedok2 = controller.service.FindOneUser(marriage.Svedok2.JMBG)

	//kreiranje vencanja je moguce samo ukoliko postoje oba svedoka u bazi
	if svedok1 == nil || svedok2 == nil {
		writer.WriteHeader(http.StatusAccepted)
		writer.Write([]byte("Ne postoji jedan od svedoka u sistemu"))
		return
	}

	marriage.Svedok1 = *svedok1
	marriage.Svedok2 = *svedok2

	controller.service.CreateNewMarriage(marriage)
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Okej"))
	//jsonResponse(token, writer)
}

func (controller *RegistrarController) UpdateCertificate(writer http.ResponseWriter, req *http.Request) {

	var userDied entity.UserDied
	err := json.NewDecoder(req.Body).Decode(&userDied)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Problem to parsing JSON to entity!"))
		return
	}

	err = controller.service.UpdateCertificate(userDied)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Okej"))
	//jsonResponse(token, writer)
}

func (controller *RegistrarController) GetChildren(writer http.ResponseWriter, req *http.Request) {

	//vars := req.URL.

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Okej"))
	//jsonResponse(token, writer)
}

func (controller *RegistrarController) Test(writer http.ResponseWriter, req *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Okej"))
	//jsonResponse(token, writer)
}
