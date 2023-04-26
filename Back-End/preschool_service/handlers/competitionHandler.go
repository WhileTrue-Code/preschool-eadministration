package handlers

import (
	"auth_service/data"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type KeyCompetition struct{}

type CompetitionsHandler struct {
	logger *log.Logger
	repo   *data.CompetitionRepo
}

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

func NewCompetitionsHandler(l *log.Logger, r *data.CompetitionRepo) *CompetitionsHandler {
	return &CompetitionsHandler{l, r}
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

func (p *CompetitionsHandler) PostCompetition(rw http.ResponseWriter, h *http.Request) {
	var insertComp data.Competition
	eerr := json.NewDecoder(h.Body).Decode(&insertComp)

	if eerr != nil {
		fmt.Println(eerr)
		http.Error(rw, "Cannot unmarshal body", 500)
		return
	}

	err := p.repo.PostCompetition(&insertComp)
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