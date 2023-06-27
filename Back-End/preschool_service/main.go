package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	natsBroker "nats"
	"net/http"
	"os"
	"os/signal"
	"preschool_service/client/registar_service"
	"preschool_service/data"
	"preschool_service/handlers"
	"time"
)

func main() {

	natsConnection := natsBroker.Conn()
	defer natsConnection.Close()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8003"
	}

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger := log.New(os.Stdout, "[preschool-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[preschool-store] ", log.LstdFlags)

	store, err := data.New(timeoutContext, storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer store.Disconnect(timeoutContext)

	store.Ping()

	//
	//applyCompetitionRepo, _ := data.NewApplyCompetitionRepo(logger)
	//
	//applyCompetitionsHandler := handlers.NewApplyCompetitionsHandler(logger, applyCompetitionRepo, registarServiceClient)

	registarServiceClient := registar_service.NewClient("registrar_service", "8001")

	applyCompetitionsHandler := handlers.NewApplyCompetitionsHandler(logger, store, registarServiceClient, natsConnection)

	router := mux.NewRouter()

	getCompetitions := router.Methods(http.MethodGet).Subrouter()
	getCompetitions.HandleFunc("/competitions/all", applyCompetitionsHandler.GetAllCompetitions)

	postCompetition := router.Methods(http.MethodPost).Subrouter()
	postCompetition.HandleFunc("/vrtic/{id}/competitions/add", applyCompetitionsHandler.PostCompetition)
	//postCompetition.Use(competitionsHandler.MiddlewareCompetitionDeserialization)

	//patchRouter := router.Methods(http.MethodPut).Subrouter()
	//patchRouter.HandleFunc("/competitions/{id}/changeStatus", applyCompetitionsHandler.PatchStatusCompetition)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/competitions/{id}", applyCompetitionsHandler.DeleteCompetition)

	changeStanjeKonkursa := router.Methods(http.MethodPut).Subrouter()
	changeStanjeKonkursa.HandleFunc("/competitions/{id}", applyCompetitionsHandler.ChangeStatus)

	postVrtic := router.Methods(http.MethodPost).Subrouter()
	postVrtic.HandleFunc("/vrtic/add", applyCompetitionsHandler.PostVrtic)

	getVrticById := router.Methods(http.MethodGet).Subrouter()
	getVrticById.HandleFunc("/vrtic/{id}", applyCompetitionsHandler.GetVrticById)

	getAllVrtici := router.Methods(http.MethodGet).Subrouter()
	getAllVrtici.HandleFunc("/vrtici/all", applyCompetitionsHandler.GetAllVrtici)

	getCompetitionById := router.Methods(http.MethodGet).Subrouter()
	getCompetitionById.HandleFunc("/competitions/getById/{id}", applyCompetitionsHandler.GetCompetitionById)

	postApplyForCompetition := router.Methods(http.MethodPost).Subrouter()
	postApplyForCompetition.HandleFunc("/competitions/{id}/apply", applyCompetitionsHandler.ApplyForCompetition)

	getAllCompetitionsApplyes := router.Methods(http.MethodGet).Subrouter()
	getAllCompetitionsApplyes.HandleFunc("/competitions/applyes", applyCompetitionsHandler.GetAllCompetitionApplyes)

	getAllApplyesForOneCompetition := router.Methods(http.MethodGet).Subrouter()
	getAllApplyesForOneCompetition.HandleFunc("/competitions/{id}/applyes", applyCompetitionsHandler.GetAllApplyesForOneCompetition)

	getApplyByID := router.Methods(http.MethodGet).Subrouter()
	getApplyByID.HandleFunc("/competitions/applyes/{id}", applyCompetitionsHandler.GetApplyById)

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
		//IdleTimeout:  120 * time.Second,
		//ReadTimeout:  1 * time.Second,
		//WriteTimeout: 1 * time.Second,
	}

	logger.Println("Server listening on port", port)
	go func() {

		if err := server.ListenAndServe(); err != nil {
			log.Println("Server served and started listening")
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	//Try to shutdown gracefully
	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")
}
