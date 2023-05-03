package service

import (
	"apr_service/domain"

	"go.uber.org/zap"
)

type AprServiceImpl struct {
	AprRepository domain.AprRepository
	Logger        *zap.Logger
}

func NewAprService(aprRepo domain.AprRepository, logger *zap.Logger) domain.AprService {
	return &AprServiceImpl{
		AprRepository: aprRepo,
		Logger:        logger,
	}
}

func (service *AprServiceImpl) RegisterAprAccount(account *domain.AprAccount) error {

	return nil
}

func (service *AprServiceImpl) FindAprByFounderID(founderID string) ([]domain.AprAccount, error) {

	return nil, nil
}
