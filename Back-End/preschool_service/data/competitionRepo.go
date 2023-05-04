package data

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type CompetitionRepo struct {
	cli    *mongo.Client
	logger *log.Logger
}

const (
	DATABASE               = "competition"
	CREDENTIALS_COLLECTION = "credentials"
)

type CompetitionRepoMongoDB struct {
	credentials mongo.Collection
}

type Competitions []*Competition

//type SmtpServer struct {
//	host string
//	port string
//}
//
//func (s *SmtpServer) ServerName() string {
//	return s.host + ":" + s.port
//}

func New(ctx context.Context, logger *log.Logger) (*CompetitionRepo, error) {
	db := os.Getenv("PRESCHOOL_DB_HOST")
	dbport := os.Getenv("PRESCHOOL_DB_PORT")

	host := fmt.Sprintf("%s:%s", db, dbport)
	client, err := mongo.NewClient(options.Client().ApplyURI(`mongodb://` + host))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &CompetitionRepo{
		cli:    client,
		logger: logger,
	}, nil
}

// GetOneUser TODO

/*func (store *AuthRepoMongoDb) filterOne(filter interface{}) (user *User, err error) {
	result := store.credentials.FindOne(context.TODO(), filter)
	err = result.Decode(&user)
	return
}*/

func (pr *CompetitionRepo) GetAll() (Competitions, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	competitionCollection := pr.getCollection()

	var competitions Competitions
	usersCursor, err := competitionCollection.Find(ctx, bson.M{})
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	if err = usersCursor.All(ctx, &competitions); err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	return competitions, nil
}

func (pr *CompetitionRepo) GetById(id string) (*Competition, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	competitionCollection := pr.getCollection()

	var competition Competition
	objID, _ := primitive.ObjectIDFromHex(id)
	err := competitionCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&competition)
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	return &competition, nil
}

func (pr *CompetitionRepo) PostCompetition(competition *Competition) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	competitionsCollection := pr.getCollection()

	competition.ID = primitive.NewObjectID()

	result, err := competitionsCollection.InsertOne(ctx, &competition)
	if err != nil {
		pr.logger.Println(err)
		return err
	}
	pr.logger.Printf("Documents ID: %v\n", result.InsertedID)

	return nil
}

// Disconnect from database
func (pr *CompetitionRepo) Disconnect(ctx context.Context) error {
	err := pr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Check database connection
func (pr *CompetitionRepo) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := pr.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		pr.logger.Println(err)
	}

	// Print available databases
	databases, err := pr.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		pr.logger.Println(err)
	}
	fmt.Println(databases)
}

func (pr *CompetitionRepo) getCollection() *mongo.Collection {
	competitionDatabase := pr.cli.Database("mongodb")
	competitionsCollection := competitionDatabase.Collection("competitions")
	return competitionsCollection
}
