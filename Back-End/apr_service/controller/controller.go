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
<<<<<<< Updated upstream
	router.HandleFunc("/{id}", controller.UpdateCompanyData).Methods("PUT")
	router.HandleFunc("/liquidate/{id}", controller.LiquidateCompany).Methods("PUT")
=======
>>>>>>> Stashed changes
	http.Handle("/", router)
	controller.Logger.Info("Controller router endpoints initialized and handle run.")
}

func (controller *AprController) RegisterAprCompany(writer http.ResponseWriter, req *http.Request) {
	controller.Logger.Info("Started registering new APR account.")
	var account domain.AprAccount
	bytes, err := io.ReadAll(req.Body)
	controller.Logger.Info("Bytes from fend", zap.String("Json stringified", string(bytes)))
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
<<<<<<< Updated upstream
	var response any = "Registration firme uspešna."
=======
	// writer.WriteHeader(http.StatusCreated)
	var response any = "Registration was successfully."
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
		http.Error(writer, "Error on server-side, please try again later.", http.StatusInternalServerError)
=======
		http.Error(writer, "Error on server-side, please try again.", http.StatusInternalServerError)
>>>>>>> Stashed changes
		return
	}

	log := controller.Logger.WithOptions(zap.Fields(
		zap.String("founderID", founderID),
	))
	log.Info("Started getting owned apr accounts by founderID")

	resBytes, _ := json.Marshal(res)
	writer.Write(resBytes)

}

func (controller *AprController) UpdateCompanyData(writer http.ResponseWriter, req *http.Request) {
	controller.Logger.Info("started UpdateCompanyData")

	bytes, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(writer, "serverska greška, molimo pokušajte kasnije.", http.StatusInternalServerError)
		return
	}
	var newCompany domain.AprAccount
	err = json.Unmarshal(bytes, &newCompany)
	if err != nil {
		http.Error(writer, "bad request.", http.StatusBadRequest)
		return
	}

	err = controller.Service.UpdateCompanyData(newCompany)
	if err != nil {
		controller.Logger.Info("error in updating company data.",
			zap.Any("company", newCompany),
		)
	}

	controller.Logger.Info("finished UpdateCompanyData")
	msg := "Podaci preduzeća su uspešno ažurirani."
	writer.Write([]byte(msg))
}

func (controller *AprController) LiquidateCompany(writer http.ResponseWriter, req *http.Request) {
	controller.Logger.Info("started LiquidateCompany")
	vars := mux.Vars(req)

	companyID, ok := vars["id"]

	if !ok {
		http.Error(writer, "company id is not provided", http.StatusBadRequest)
		return
	}

	err := controller.Service.LiquidateCompany(companyID)
	if err != nil {
		controller.Logger.Info("error in liquidating company.",
			zap.String("companyID", companyID),
		)
	}

	controller.Logger.Info("finished LiquidateCompany")
	msg := "company successfully liquidated."

	writer.Write([]byte(msg))
}
