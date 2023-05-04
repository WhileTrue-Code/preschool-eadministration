package service

import (
	"apr_service/domain"

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
	return service.Repo.SaveAprAccount(account)
}

func (service *AprServiceImpl) FindAprByFounderID(founderID string) ([]domain.AprAccount, error) {

	return nil, nil
}
