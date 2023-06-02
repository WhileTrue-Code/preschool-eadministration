package controller

import (
	"apr_service/domain"
	"authorization"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	router.HandleFunc("/", controller.FindAprByFounderID).Methods("GET")
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
	account.ID = primitive.NewObjectID()

	authToken := req.Header.Get("Authorization")
	splitted := strings.Split(authToken, " ")
	claims := authorization.GetMapClaims([]byte(splitted[1]))

	founderID := claims["jmbg"]
	account.FounderID = founderID

	err = controller.Service.RegisterAprAccount(&account)
	if err != nil {
		http.Error(writer, "error in saving AprAccount to repository.", http.StatusInternalServerError)
		return
	}

	controller.Logger.Info("ACCOUNT MODEL STRUCT",
		zap.Any("account", account),
	)
	controller.Logger.Info("APR account registered successfully.")
	var response any = "Registration firme uspešna."
	stringResp, _ := json.Marshal(response)
	writer.Write([]byte(stringResp))

}

func (controller *AprController) FindAprByFounderID(writer http.ResponseWriter, req *http.Request) {
	authToken := req.Header.Get("Authorization")
	splitted := strings.Split(authToken, " ")
	claims := authorization.GetMapClaims([]byte(splitted[1]))

	founderID := claims["jmbg"]
	res, err := controller.Service.FindAprByFounderID(founderID)
	if err != nil {
		http.Error(writer, "Error on server-side, please try again later.", http.StatusInternalServerError)
		return
	}

	log := controller.Logger.WithOptions(zap.Fields(
		zap.String("founderID", founderID),
	))
	log.Info("Started getting owned apr accounts by founderID")

	resBytes, _ := json.Marshal(res)
	writer.Write(resBytes)

}
