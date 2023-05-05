package service

import (
	"apr_service/domain"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"go.uber.org/zap"
)

type AprServiceImpl struct {
	Repo   domain.AprRepository
	Logger *zap.Logger
}

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
