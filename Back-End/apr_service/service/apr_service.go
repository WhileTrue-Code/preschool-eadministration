package service

import (
	"apr_service/domain"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type AprServiceImpl struct {
	Repo   domain.AprRepository
	Logger *zap.Logger
}

var (
	GET_COMPANY_QUEUE = "COMPANY_GET_BY_FOUNDER_COMPANY_ID"
)

func NewAprService(aprRepo domain.AprRepository, logger *zap.Logger) domain.AprService {
	return &AprServiceImpl{
		Repo:   aprRepo,
		Logger: logger,
	}
}

func (service *AprServiceImpl) RegisterAprAccount(account *domain.AprAccount) error {
	account.CompanyID = service.generatePIB()
	return service.Repo.SaveAprAccount(account)
}

func (service *AprServiceImpl) FindAprByFounderID(founderID string) ([]domain.AprAccount, error) {
	return service.Repo.FindAprAccountsByFounderID(founderID)
}

func (service *AprServiceImpl) FindByFounderIDAndCompanyID(founderID string,
	companyID int) (company domain.AprAccount, err error) {

	return service.Repo.FindCompanyByFounderIDAndCompanyID(founderID, companyID)
}

func (service *AprServiceImpl) generatePIB() (pibI int) {
	rand.Seed(time.Now().UnixNano())
	pib := fmt.Sprintf("%d", rand.Intn(9)+1)

	for i := 0; i < 7; i++ {
		pib += fmt.Sprintf("%d", rand.Intn(10))
	}
	pibInt, err := strconv.Atoi(pib)
	if err != nil {
		service.Logger.Error("error in ATOI func", zap.Error(err))
		return
	}

	pibI = pibInt
	for service.Repo.DoesExistAprWithID(pibI) {
		pibI = service.generatePIB()
		service.Logger.Info("New generating PIB!",
			zap.Int("newPib", pibI),
		)
	}
	//uslov za postojanje/nepostojanje!
	// for !isValidPIB(pib) {
	// 	pibI = service.generatePIB()
	// }

	return
}

func (service *AprServiceImpl) SubscribeToNats(connection *nats.Conn) {
	_, err := connection.QueueSubscribe(os.Getenv(GET_COMPANY_QUEUE), GET_COMPANY_QUEUE, func(message *nats.Msg) {
		var request map[string]string
		err := json.Unmarshal(message.Data, &request)
		if err != nil {
			service.Logger.Error("error in unmarshalling GetCompany struct")
			return
		}

		founderID := request["founderID"]
		companyID, _ := strconv.Atoi(request["companyID"])

		company, err := service.FindByFounderIDAndCompanyID(founderID, companyID)
		if err != nil {
			service.Logger.Error("error in finding company",
				zap.Error(err),
				zap.String("founderID", founderID),
				zap.Int("companyID", companyID),
			)
			return
		}

		response, err := json.Marshal(company)
		if err != nil {
			service.Logger.Error("error in marshalling Company struct",
				zap.Error(err),
			)
			return
		}

		err = connection.Publish(message.Reply, response)
		if err != nil {
			service.Logger.Error("error in publishing response",
				zap.Error(err),
			)
			return
		}

	})
	if err != nil {
		service.Logger.Error("error in subscribing to NATS queue",
			zap.Error(err),
		)
		return
	}

	log.Printf("Subscribed to channel: %s", os.Getenv(GET_COMPANY_QUEUE))
}
