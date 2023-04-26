package startup

import (
	"apr_service/controller"
	"apr_service/domain"
	"apr_service/service"
	"apr_service/startup/config"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
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
	service := server.initAprService()
	controller := server.initController(service)
	server.start(controller)
}

func (server *Server) initAprService() domain.AprService {
	return service.NewAprService(nil, server.Logger)
}

func (server *Server) initController(service domain.AprService) *controller.AprController {
	return controller.NewController(service, server.Logger)
}

func (server *Server) start(controller *controller.AprController) {
	router := mux.NewRouter()
	controller.Init(router)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", server.Config.SERVICE_PORT),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			server.Logger.Info("Server served and started listening")
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
		server.Logger.Error("Error on shutting down server.")
	}

	server.Logger.Info("Gracefully shutdown executed.")
}
