package handlers

import (
	"authorization"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nats-io/nats.go"
	"log"
	"net/http"
	"os"
	"preschool_service/client/registar_service"
	"preschool_service/data"
	"strings"
	"time"
)

type KeyCompetition struct{}

type ApplyCompetitionHandler struct {
	logger          *log.Logger
	repo            *data.ApplyCompetitionRepo
	registarService registar_service.Client
	nats            *nats.Conn
}

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

func NewApplyCompetitionsHandler(l *log.Logger, r *data.ApplyCompetitionRepo, registarService registar_service.Client, nats *nats.Conn) *ApplyCompetitionHandler {
	return &ApplyCompetitionHandler{l, r, registarService, nats}
}

func (p *ApplyCompetitionHandler) ApplyForCompetition(rw http.ResponseWriter, h *http.Request) {
	authToken := h.Header.Get("Authorization")

	var insertComp data.Prijava
	eerr := json.NewDecoder(h.Body).Decode(&insertComp)

	if eerr != nil {
		fmt.Println(eerr)
		http.Error(rw, "Cannot unmarshal body", 500)
		return
	}

	vars := mux.Vars(h)
	competitionID := vars["id"]

	//cole **********
	isParent, errr := p.registarService.GetIsParent(insertComp.Dete.JMBG, authToken)
	if errr != nil {
		http.Error(rw, errr.Error(), http.StatusInternalServerError)
		return
	}

	if !isParent {
		http.Error(rw, "You are not a parent "+insertComp.Dete.JMBG, 403)
		return
	}
	//cole **********

	// silja***********
	splitted := strings.Split(authToken, " ")
	claims := authorization.GetMapClaims([]byte(splitted[1]))

	request := map[string]string{
		"employeeID": claims["jmbg"],
	}

	requestBytes, err := json.Marshal(request)

	msg, err := p.nats.Request(os.Getenv("GET_EMPLOYEE_STATUS_BY_ID"), requestBytes, 5*time.Second)
	if err != nil {
		log.Println(err)
		println("eror pri preuzimanju req")
	}
	var response map[string]bool
	err = json.Unmarshal(msg.Data, &response)

	if response["employed"] {
		insertComp.Bodovi = 1
	}
	println(request)
	println(response)
	// silja***********

	// miljus *****
	dataToSend, err := json.Marshal(insertComp.Dete.JMBG)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	response1, err := p.nats.Request(os.Getenv("GET_STANJE_BY_JMBG"), dataToSend, 5*time.Second)

	var deteZS data.ZdravstvenoStanje

	if response1.Data == nil {
		http.Error(rw, "Dete nema ZS", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(response1.Data, &deteZS)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return
	}

	if deteZS.ZdravstveniProblemi != "" {
		insertComp.Bodovi = insertComp.Bodovi + 2
	}
	if deteZS.SmetnjeURazvoju != "" {
		insertComp.Bodovi = insertComp.Bodovi + 3
	}

	// miljus *****

	err = p.repo.ApplyForCompetition(competitionID, &insertComp)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	rw.WriteHeader(http.StatusCreated)
}

func (p *ApplyCompetitionHandler) GetAllCompetitionApplyes(rw http.ResponseWriter, h *http.Request) {
	allCompetitions, err := p.repo.GetAllApplyes()
	if err != nil {
		http.Error(rw, "Database exception", http.StatusInternalServerError)
		p.logger.Fatal("Database exception: ", err)
	}

	err = allCompetitions.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (p *ApplyCompetitionHandler) GetApplyById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	allCompetitions, err := p.repo.GetPrijavaById(id)
	if err != nil {
		http.Error(rw, "Database exception", http.StatusInternalServerError)
		p.logger.Fatal("Database exception: ", err)
	}

	err = allCompetitions.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (p *ApplyCompetitionHandler) GetAllApplyesForOneCompetition(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	allCompetitions, err := p.repo.GetAllApplyesForOneCompetition(id)
	if err != nil {
		http.Error(rw, "Database exception", http.StatusInternalServerError)
		p.logger.Fatal("Database exception: ", err)
	}

	err = allCompetitions.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (p *ApplyCompetitionHandler) GetAllCompetitions(rw http.ResponseWriter, h *http.Request) {
	allCompetitions, err := p.repo.GetAll()
	if err != nil {
		http.Error(rw, "Database exception", http.StatusInternalServerError)
		p.logger.Fatal("Database exception: ", err)
	}

	err = allCompetitions.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (p *ApplyCompetitionHandler) GetAllVrtici(rw http.ResponseWriter, h *http.Request) {
	vrtici, err := p.repo.GetAllVrtici()
	if err != nil {
		http.Error(rw, "Database exception", http.StatusInternalServerError)
		p.logger.Fatal("Database exception: ", err)
	}

	err = vrtici.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (p *ApplyCompetitionHandler) GetCompetitionById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	competition, err := p.repo.GetById(id)
	if err != nil {
		p.logger.Print("Database exception: ", err)
	}

	if competition == nil {
		http.Error(rw, "Patient with given id not found", http.StatusNotFound)
		p.logger.Printf("Patient with id: '%s' not found", id)
		return
	}

	err = competition.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (p *ApplyCompetitionHandler) GetVrticById(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	vrtic, err := p.repo.GetVrticById(id)
	if err != nil {
		p.logger.Print("Database exception: ", err)
	}

	if vrtic == nil {
		http.Error(rw, "Vrtic with given id not found", http.StatusNotFound)
		p.logger.Printf("Vrtic with id: '%s' not found", id)
		return
	}

	err = vrtic.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		p.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (p *ApplyCompetitionHandler) PostCompetition(rw http.ResponseWriter, h *http.Request) {
	var insertComp data.Competition
	eerr := json.NewDecoder(h.Body).Decode(&insertComp)

	if eerr != nil {
		fmt.Println(eerr)
		http.Error(rw, "Cannot unmarshal body", 500)
		return
	}

	vars := mux.Vars(h)
	vrticID := vars["id"]

	err := p.repo.PostCompetition(vrticID, &insertComp)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	rw.WriteHeader(http.StatusCreated)
}

func (p *ApplyCompetitionHandler) PostVrtic(rw http.ResponseWriter, h *http.Request) {
	var insertComp data.Vrtic
	eerr := json.NewDecoder(h.Body).Decode(&insertComp)

	if eerr != nil {
		fmt.Println(eerr)
		http.Error(rw, "Cannot unmarshal body", 500)
		return
	}

	err := p.repo.PostVrtic(&insertComp)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	rw.WriteHeader(http.StatusCreated)
}

func (p *ApplyCompetitionHandler) ChangeStatus(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	var status string
	d := json.NewDecoder(h.Body)
	d.Decode(&status)

	p.repo.ChangeStatus(id)
	rw.WriteHeader(http.StatusOK)
}

func (p *ApplyCompetitionHandler) DeleteCompetition(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	p.repo.Delete(id)
	rw.WriteHeader(http.StatusNoContent)
}

func jsonResponse(object interface{}, w http.ResponseWriter) {
	resp, err := json.Marshal(object)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (p *ApplyCompetitionHandler) MiddlewareCompetitionDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		competition := &data.Competition{}
		err := competition.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			p.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyCompetition{}, competition)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
