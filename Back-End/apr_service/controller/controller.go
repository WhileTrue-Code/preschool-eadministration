package controller

import (
	"apr_service/domain"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type AprController struct {
	Logger     *zap.Logger
	AprService domain.AprService
}

func NewController(aprService domain.AprService, logger *zap.Logger) *AprController {
	return &AprController{
		Logger:     logger,
		AprService: aprService,
	}
}

func (controller *AprController) Init(router *mux.Router) {
	router.HandleFunc("/register", controller.RegisterAprCompany).Methods("POST")
	router.HandleFunc("/{founderID}", controller.FindAprByFounderID).Methods("GET")
	http.Handle("/", router)
	controller.Logger.Info("Controller router endpoints initialized and handle run.")
}

func (controller *AprController) RegisterAprCompany(writer http.ResponseWriter, req *http.Request) {
	controller.Logger.Info("Started registering new apr company account.")
	writer.Write([]byte("All is ok!"))

}

func (controller *AprController) FindAprByFounderID(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	founderID, ok := vars["founderID"]
	if !ok {
		http.Error(writer, "founder id wasn't provided.", http.StatusBadRequest)
		return
	}

	log := controller.Logger.WithOptions(zap.Fields(
		zap.String("founderID", founderID),
	))
	log.Info("Started getting owned apr accounts by founderID")

	writer.Write([]byte("All is ok!"))

}
