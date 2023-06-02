package startup

import (
	"context"
	"croso_service/controller"
	"croso_service/domain"
	"croso_service/repo"
	"croso_service/service"
	"croso_service/startup/config"
	"fmt"
	"log"
	natsBroker "nats"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type Server struct {
	Config *config.Config
	Logger *zap.Logger
}

func NewServer() *Server {
	logger, err := InitAndConfigureLogger()
	if err != nil {
		log.Printf("Error in initialization logger cause of: %s", err)
		panic(err)
	}
	server := &Server{
		Config: config.NewConfig(),
		Logger: logger,
	}

	server.Logger.Info("Server object successfully initialized", zap.Any("server", *server))
	return server
}

func (server *Server) Start() {
	natsConnection := natsBroker.Conn()
	defer natsConnection.Close()

	server.Logger.Info("LOG AFTER CONNECTION FROM NATS GOT")

	repository := server.initCrosoRepository()
	serviceInstance := server.initCrosoService(repository, natsConnection)
	serviceInstance.SubscribeToNats(natsConnection)

	controllerInstance := server.initController(serviceInstance)

	server.start(controllerInstance)
}

func (server *Server) initCrosoRepository() domain.CrosoRepository {
	cli := repo.GetMongoClient(server.Config.DB_HOST, server.Config.DB_PORT, server.Logger)

	repoLogger := server.Logger.Named("[CROSO / REPOSITORY]")
	if cli == nil {
		repoLogger.Error("MongoDB cli is null, shutting down...")
		os.Exit(1)
	}

	return repo.NewMongoRepo(cli, repoLogger)
}

func (server *Server) initCrosoService(repo domain.CrosoRepository, nats *nats.Conn) domain.CrosoService {
	return service.NewAprService(repo, nats, server.Logger.Named("[CROSO / SERVICE]"))
}

func (server *Server) initController(service domain.CrosoService) *controller.CrosoController {
	return controller.NewController(service, server.Logger.Named("[CROSO / CONTROLLER]"))
}

func (server *Server) start(controller *controller.CrosoController) {
	router := mux.NewRouter()
	controller.Init(router)
	log := server.Logger.Named("[CROSO / SERVER]")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", server.Config.SERVICE_PORT),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Info("Server served and started listening")
		}
	}()

	gShutdownChannel := make(chan os.Signal, 1)
	signal.Notify(gShutdownChannel,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	<-gShutdownChannel

	timeout := time.Second * 15
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Error on shutting down server.")
	}

	log.Info("Gracefully shutdown executed.")
}
