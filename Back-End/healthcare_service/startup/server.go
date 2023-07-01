package startup

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
	"healthcare_service/controller"
	"healthcare_service/repository"
	"healthcare_service/repository/repository_impl"
	"healthcare_service/service"
	"healthcare_service/startup/config"
	"log"
	messageBroker "nats"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	//connectiong to NATS container
	natsConnection := messageBroker.Conn()
	defer natsConnection.Close()

	mongoClient := server.initMongoClient()
	defer func(mongoClient *mongo.Client, ctx context.Context) {
		err := mongoClient.Disconnect(ctx)
		if err != nil {
			log.Println(err)
		}
	}(mongoClient, context.Background())

	healthcareRepository := server.initHealthcareRepository(mongoClient)
	healthcareService := server.initHealthcareService(healthcareRepository, natsConnection)
	healthcareController := server.initHealthcareController(healthcareService)

	healthcareService.SubscribeToNats(natsConnection)

	server.start(healthcareController)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := repository.GetMongoClient(server.config.HealthcareDBHost, server.config.HealthcareDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initHealthcareRepository(client *mongo.Client) repository.HealthcareRepository {
	store := repository_impl.NewAuthRepositoryImpl(client)
	return store
}

func (server *Server) initHealthcareService(store repository.HealthcareRepository, natsConnection *nats.Conn) *service.HealthcareService {
	return service.NewHealthcareService(store, natsConnection)
}

func (server *Server) initHealthcareController(service *service.HealthcareService) *controller.HealthcareController {
	return controller.NewHealthcareController(service)
}

func (server *Server) start(healthcareController *controller.HealthcareController) {
	router := mux.NewRouter()
	healthcareController.Init(router)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", server.config.HealthcareServicePort),
		Handler: router,
	}

	wait := time.Second * 15
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Error Shutting Down Server %s", err)
	}
	log.Println("Server Gracefully Stopped")
}
