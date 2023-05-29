package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"preschool_service/client/registar_service"
	"preschool_service/data"
)

type KeyCompetition struct{}

type CompetitionsHandler struct {
	logger *log.Logger
	repo   *data.CompetitionRepo
}

type ApplyCompetitionHandler struct {
	logger          *log.Logger
	repo            *data.ApplyCompetitionRepo
	registarService registar_service.Client
}

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

func NewCompetitionsHandler(l *log.Logger, r *data.CompetitionRepo) *CompetitionsHandler {
	return &CompetitionsHandler{l, r}
}

func NewApplyCompetitionsHandler(l *log.Logger, r *data.ApplyCompetitionRepo, registarService registar_service.Client) *ApplyCompetitionHandler {
	return &ApplyCompetitionHandler{l, r, registarService}
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

	isParent, errr := p.registarService.GetIsParent(insertComp.Dete.JMBG, authToken)
	if errr != nil {
		http.Error(rw, errr.Error(), http.StatusInternalServerError)
		return
	}

	if !isParent {
		http.Error(rw, "You are not a parent "+insertComp.Dete.JMBG, 403)
		return
	}

	err := p.repo.ApplyForCompetition(competitionID, &insertComp)
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

func (p *CompetitionsHandler) GetAllCompetitions(rw http.ResponseWriter, h *http.Request) {
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

func (p *CompetitionsHandler) GetAllVrtici(rw http.ResponseWriter, h *http.Request) {
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

func (p *CompetitionsHandler) GetCompetitionById(rw http.ResponseWriter, h *http.Request) {
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

func (p *CompetitionsHandler) GetVrticById(rw http.ResponseWriter, h *http.Request) {
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

func (p *CompetitionsHandler) PostCompetition(rw http.ResponseWriter, h *http.Request) {
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

func (p *CompetitionsHandler) PostVrtic(rw http.ResponseWriter, h *http.Request) {
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

//func (p *CompetitionsHandler) PostCompetition(rw http.ResponseWriter, h *http.Request) {
//	usr := h.Context().Value(KeyCompetition{}).(*data.Competition)
//	err := p.repo.PostCompetition(usr)
//	if err != nil {
//		rw.WriteHeader(http.StatusInternalServerError)
//	}
//
//	rw.WriteHeader(http.StatusCreated)
//}

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

func (p *CompetitionsHandler) MiddlewareCompetitionDeserialization(next http.Handler) http.Handler {
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
