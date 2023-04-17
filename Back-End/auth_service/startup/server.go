package startup

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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

	mongoClient := server.initMongoClient()
	defer func(mongoClient *mongo.Client, ctx context.Context) {
		err := mongoClient.Disconnect(ctx)
		if err != nil {
			log.Println(err)
		}
	}(mongoClient, context.Background())

	authRepository := server.initAuthRepository(mongoClient)
	authService := server.initAuthService(authRepository)
	authController := server.initAuthController(authService)

	server.start(authController)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := repository.GetMongoClient(server.config.AuthDBHost, server.config.AuthDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initAuthRepository(client *mongo.Client) repository.AuthRepository {
	store := repository_impl.NewAuthRepositoryImpl(client)
	return store
}

func (server *Server) initAuthService(store repository.AuthRepository) *service.AuthService {
	return service.NewAuthService(store)
}

func (server *Server) initAuthController(service *service.AuthService) *controller.AuthController {
	return controller.NewAuthController(service)
}

// start
func (server *Server) start(registrarController *controller.AuthController) {
	router := mux.NewRouter()
	registrarController.Init(router)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", server.config.AuthServicePort),
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
