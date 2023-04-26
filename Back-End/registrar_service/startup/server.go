package startup

import (
	"context"
	"fmt"
	"log"
	messageBroker "nats"
	"net/http"
	"os"
	"os/signal"
	"registrar_service/controller"
	"registrar_service/repository"
	"registrar_service/repository/repository_impl"
	"registrar_service/service"
	"registrar_service/startup/config"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/mongo"
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

	registrarRepository := server.initRegistrarRepository(mongoClient)
	registrarService := server.initRegistrarService(registrarRepository, natsConnection)
	registrarController := server.initRegistrarController(registrarService)

	//pozivanje metode za subscribe na nats
	registrarService.SubscribeToNats(natsConnection)

	server.start(registrarController)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := repository.GetMongoClient(server.config.RegistrarDBHost, server.config.RegistrarDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initRegistrarRepository(client *mongo.Client) repository.RegistrarRepository {
	store := repository_impl.NewRegistrarRepositoryImpl(client)
	return store
}

func (server *Server) initRegistrarService(store repository.RegistrarRepository, natsConnection *nats.Conn) *service.RegistrarService {
	return service.NewRegistrarService(store, natsConnection)
}

func (server *Server) initRegistrarController(service *service.RegistrarService) *controller.RegistrarController {
	return controller.NewRegistrarController(service)
}

// start
func (server *Server) start(registrarController *controller.RegistrarController) {
	router := mux.NewRouter()
	registrarController.Init(router)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", server.config.RegistrarServicePort),
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
