package repo

import (
	"context"
	"croso_service/domain"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var (
	DB                   = "croso"
	COLLECTION           = "accounts"
	COLLECTION_EMPLOYEES = "employees"
)

type CrosoMongoRepo struct {
	Collection          *mongo.Collection
	CollectionEmployees *mongo.Collection
	Logger              *zap.Logger
}

func NewMongoRepo(client *mongo.Client, logger *zap.Logger) domain.CrosoRepository {
	return &CrosoMongoRepo{
		Collection:          client.Database(DB).Collection(COLLECTION),
		CollectionEmployees: client.Database(DB).Collection(COLLECTION_EMPLOYEES),
		Logger:              logger,
	}
}

func (repo *CrosoMongoRepo) SaveCrosoAccount(crosoAccount *domain.CrosoAccount) (err error) {
	_, err = repo.Collection.InsertOne(context.Background(), crosoAccount)
	if err != nil {
		repo.Logger.Error("error in inserting AprAccount.",
			zap.Error(err),
		)
		return
	}
	return
}

func (repo *CrosoMongoRepo) FindCrosoAccountsByFounderID(founderID string) (results []domain.CrosoAccount) {
	cursor, err := repo.Collection.Find(context.Background(), bson.D{{Key: "founderID", Value: founderID}})
	if err != nil {
		repo.Logger.Info("Error in getting accounts by founderID",
			zap.Error(err),
		)
		return
	}

	results = make([]domain.CrosoAccount, 0)
	err = cursor.All(context.Background(), &results)
	if err != nil {
		repo.Logger.Info("Error in decoding by results with All()",
			zap.Error(err),
		)
		return
	}

	return
}

func (repo *CrosoMongoRepo) SaveEmployee(request *domain.Employee) (err error) {
	_, err = repo.CollectionEmployees.InsertOne(context.Background(), *request)
	if err != nil {
		repo.Logger.Error("error in inserting RegisterEmployeeRequest",
			zap.Error(err),
		)
	}

	return
}

func (repo *CrosoMongoRepo) GetEmployee(filter bson.M) (employee *domain.Employee) {
	result := repo.CollectionEmployees.FindOne(context.Background(), filter)
	if result.Err() != nil {
		repo.Logger.Error("error in getting employee from db",
			zap.Any("filter", filter),
			zap.Error(result.Err()),
		)
		return
	}
	employeeDecode := domain.Employee{}
	err := result.Decode(&employeeDecode)
	if err != nil {
		repo.Logger.Error("error in decoding result into struct Employee",
			zap.Error(err),
		)
		return
	}

	employee = &employeeDecode

	return
}

func (repo *CrosoMongoRepo) GetEmployees(filter bson.D) (employees []domain.Employee) {
	results, err := repo.CollectionEmployees.Find(context.Background(), filter)

	employees = make([]domain.Employee, 0)

	err = results.Decode(&employees)
	if err != nil {
		repo.Logger.Error("error in decoding cursor results to []domain.Employee",
			zap.Error(err),
		)
		return
	}

	return
}

func (repo *CrosoMongoRepo) UpdateEmployee(employee *domain.Employee) (err error) {
	_, err = repo.CollectionEmployees.UpdateByID(context.Background(), employee.ID, bson.M{"$set": employee})
	return
}

func (repo *CrosoMongoRepo) FindEmployeesWithCompanyID(companyID string) (employees []domain.Employee, err error) {
	intComp, _ := strconv.Atoi(companyID)
	cursor, err := repo.CollectionEmployees.Find(context.Background(),
		bson.D{
			{Key: "companyID", Value: intComp},
			{Key: "registrationStatus", Value: domain.ACCEPTED},
		},
	)
	if err != nil {
		repo.Logger.Error("error in finding employees with companyID",
			zap.Error(err),
		)
		return
	}
	employees = []domain.Employee{}
	err = cursor.All(context.Background(), &employees)
	if err != nil {
		repo.Logger.Error("error in decoding cursor into slice of struct",
			zap.Error(err),
		)
	}

	repo.Logger.Info("logging slice", zap.Any("slice", employees))

	return
}
