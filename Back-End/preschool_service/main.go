package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"preschool_service/client/registar_service"
	"preschool_service/data"
	"preschool_service/handlers"
	"time"
)

func main() {
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

	registarServiceClient := registar_service.NewClient("registrar_service", "8001")

	competitionsHandler := handlers.NewCompetitionsHandler(logger, store)
	applyCompetitionsHandler := handlers.NewApplyCompetitionsHandler(logger, (*data.ApplyCompetitionRepo)(store), registarServiceClient)

	router := mux.NewRouter()

	getCompetitions := router.Methods(http.MethodGet).Subrouter()
	getCompetitions.HandleFunc("/competitions/all", competitionsHandler.GetAllCompetitions)

	postCompetition := router.Methods(http.MethodPost).Subrouter()
	postCompetition.HandleFunc("/vrtic/{id}/competitions/add", competitionsHandler.PostCompetition)
	//postCompetition.Use(competitionsHandler.MiddlewareCompetitionDeserialization)

	postVrtic := router.Methods(http.MethodPost).Subrouter()
	postVrtic.HandleFunc("/vrtic/add", competitionsHandler.PostVrtic)

	getVrticById := router.Methods(http.MethodGet).Subrouter()
	getVrticById.HandleFunc("/vrtic/{id}", competitionsHandler.GetVrticById)

	getAllVrtici := router.Methods(http.MethodGet).Subrouter()
	getAllVrtici.HandleFunc("/vrtici/all", competitionsHandler.GetAllVrtici)

	getCompetitionById := router.Methods(http.MethodGet).Subrouter()
	getCompetitionById.HandleFunc("/competitions/getById/{id}", competitionsHandler.GetCompetitionById)

	postApplyForCompetition := router.Methods(http.MethodPost).Subrouter()
	postApplyForCompetition.HandleFunc("/competitions/{id}/apply", applyCompetitionsHandler.ApplyForCompetition)

	getAllCompetitionsApplyes := router.Methods(http.MethodGet).Subrouter()
	getAllCompetitionsApplyes.HandleFunc("/competitions/applyes", applyCompetitionsHandler.GetAllCompetitionApplyes)

	//getAllApplyesForOneCompetition := router.Methods(http.MethodGet).Subrouter()
	//getAllApplyesForOneCompetition.HandleFunc("/competitions/{id}/applyes", applyCompetitionsHandler.GetAllApplyesForOneCompetition)

	getApplyByID := router.Methods(http.MethodGet).Subrouter()
	getApplyByID.HandleFunc("/competitions/applyes/{id}", applyCompetitionsHandler.GetApplyById)

	//cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"http://localhost:4200"}),
	//	gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
	//	gorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"}),
	//	gorillaHandlers.AllowCredentials())

	server := http.Server{
		Addr:         ":" + port,
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	//certFile := "twitter.crt"
	//keyFile := "twitter.key"

	logger.Println("Server listening on port", port)
	//Distribute all the connections to goroutines
	go func() {
		//err := server.ListenAndServeTLS(certFile, keyFile)

		err := server.ListenAndServe()

		if err != nil {
			logger.Fatal(err)
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
