package controller

import (
	"apr_service/domain"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type AprController struct {
	Logger  *zap.Logger
	Service domain.AprService
}

func NewController(aprService domain.AprService, logger *zap.Logger) *AprController {
	return &AprController{
		Logger:  logger,
		Service: aprService,
	}
}

func (controller *AprController) Init(router *mux.Router) {
	router.HandleFunc("/register", controller.RegisterAprCompany).Methods("POST")
	router.HandleFunc("/{founderID}", controller.FindAprByFounderID).Methods("GET")
	http.Handle("/", router)
	controller.Logger.Info("Controller router endpoints initialized and handle run.")
}

func (controller *AprController) RegisterAprCompany(writer http.ResponseWriter, req *http.Request) {
	controller.Logger.Info("Started registering new APR account.")
	var account domain.AprAccount
	bytes, err := io.ReadAll(req.Body)
	if err != nil {
		controller.Logger.Error("error in reading request body bytes.",
			zap.Error(err),
		)
		http.Error(writer, "error in reading request body bytes.", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(bytes, &account)
	if err != nil {
		controller.Logger.Error("error in unmarshalling bytes to model.",
			zap.Error(err),
		)
		http.Error(writer, "error in unmarshalling bytes to model.", http.StatusBadRequest)
		return
	}

	controller.Logger.Info("ACCOUNT MODEL STRUCT",
		zap.Any("account", account),
	)

	err = controller.Service.RegisterAprAccount(&account)
	if err != nil {
		http.Error(writer, "error in saving AprAccount to repository.", http.StatusInternalServerError)
		return
	}

	controller.Logger.Info("APR account registered successfully.")
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("Registration was successfully."))

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
