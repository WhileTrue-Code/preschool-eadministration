package repo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func GetMongoClient(host, port string, logger *zap.Logger) (client *mongo.Client) {
	dbUri := fmt.Sprintf("mongodb://%s:%s/", host, port)
	cliOpts := options.Client().ApplyURI(dbUri)
	client, err := mongo.Connect(context.Background(), cliOpts)
	if err != nil {
		logger.Error("error in connecting to MongoDB and getting client.",
			zap.Error(err),
		)
		return
	}

	return
}
