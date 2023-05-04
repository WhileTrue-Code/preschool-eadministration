package repo

import (
	"apr_service/domain"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var (
	DB         = "apr"
	COLLECTION = "accounts"
)

type AprMongoRepo struct {
	Collection *mongo.Collection
	Logger     *zap.Logger
}

func NewMongoRepo(client *mongo.Client, logger *zap.Logger) domain.AprRepository {
	return &AprMongoRepo{
		Collection: client.Database(DB).Collection(COLLECTION),
		Logger:     logger,
	}
}

func (repo *AprMongoRepo) SaveAprAccount(aprAccount *domain.AprAccount) (err error) {
	_, err = repo.Collection.InsertOne(context.Background(), aprAccount)
	if err != nil {
		repo.Logger.Error("error in inserting AprAccount.",
			zap.Error(err),
		)
		return
	}
	return
}

func (repo *AprMongoRepo) FindAprAccountsByFounderID(founderID string) ([]domain.AprAccount, error) {
	return nil, nil
}
