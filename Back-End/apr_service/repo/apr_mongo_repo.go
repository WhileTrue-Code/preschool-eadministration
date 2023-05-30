package repo

import (
	"apr_service/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
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

func (repo *AprMongoRepo) FindAprAccountsByFounderID(founderID string) (results []domain.AprAccount, err error) {
	cursor, err := repo.Collection.Find(context.Background(), bson.D{{Key: "founderID", Value: founderID}})
	if err != nil {
		repo.Logger.Info("Error in getting accounts by founderID",
			zap.Error(err),
		)
		return
	}

	results = []domain.AprAccount{}
	err = cursor.All(context.Background(), &results)
	if err != nil {
		repo.Logger.Info("Error in decoding by results with All()",
			zap.Error(err),
		)
		return
	}

	return
}

func (repo *AprMongoRepo) FindCompanyByFounderIDAndCompanyID(founderID string,
	companyID int) (company domain.AprAccount, err error) {
	result := repo.Collection.FindOne(context.Background(),
		bson.D{
			{Key: "founderID", Value: founderID},
			{Key: "companyID", Value: companyID},
		},
	)
	if err != nil {
		repo.Logger.Info("Error in getting accounts by founderID",
			zap.Error(err),
		)
		return
	}

	err = result.Decode(&company)
	if err != nil {
		repo.Logger.Info("Error in decoding result",
			zap.Error(err),
		)
		return
	}

	return
}

func (repo *AprMongoRepo) DoesExistAprWithID(ID int) (exists bool) {

	result := repo.Collection.FindOne(context.Background(), bson.M{"companyID": ID})
	if err := result.Err(); err != nil {
		return false
	}

	return true
}
