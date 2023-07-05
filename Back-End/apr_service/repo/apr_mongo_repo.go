package repo

import (
	"apr_service/domain"
	"apr_service/errors"
	"context"
	"fmt"
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

	cursor, err := repo.Collection.Find(context.Background(), bson.D{{Key: "founderID", Value: founderID},
		{Key: "isLiquidated", Value: false},
	})
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

func (repo *AprMongoRepo) FindAprAccountsByCompanyID(companyID int) (found domain.AprAccount, err error) {
	result := repo.Collection.FindOne(context.Background(), bson.D{{Key: "companyID", Value: companyID}})
	if result.Err() != nil {
		repo.Logger.Info("Error in getting accounts by founderID",
			zap.Error(err),
		)
		return
	}
	found = domain.AprAccount{}
	err = result.Decode(&found)
	if err != nil {
		repo.Logger.Info("Error in decoding found result.",
			zap.Error(err),
		)
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

func (repo *AprMongoRepo) PatchCompany(newCompany domain.AprAccount) (err error) {
	result, err := repo.Collection.UpdateByID(context.Background(), newCompany.ID, bson.M{"$set": newCompany})
	if result.UpsertedCount != 1 && err != nil {
		repo.Logger.Error(errors.ERR_PATCHING_BY_ID,
			zap.Error(err),
			zap.Any("newCompany", newCompany),
		)
		return fmt.Errorf(errors.ERR_PATCHING_BY_ID)
	}

	return nil

}

//func (repo *AprMongoRepo) PatchCompany(newCompany domain.AprAccount) (err error) {
//	repo.Logger.Info("company\n\n\n\n\n\n", zap.Any("company", newCompany))
//	result, err := repo.Collection.UpdateByID(context.Background(), newCompany.ID, newCompany)
//	if result.UpsertedCount != 1 && err != nil {
//		repo.Logger.Error(errors.ERR_PATCHING_BY_ID,
//			zap.Error(err),
//			zap.Any("newCompany", newCompany),
//		)
//		return fmt.Errorf(errors.ERR_PATCHING_BY_ID)
//	}
//
//	return nil
//
//}
