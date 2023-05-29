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

type ApplyCompetitionRepo struct {
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

type Prijave []*Prijava

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

func (pr *ApplyCompetitionRepo) GetAllApplyes() (Prijave, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	competitionCollection := pr.getCollectionCompetitionApply()

	var prijave Prijave
	usersCursor, err := competitionCollection.Find(ctx, bson.M{})
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	if err = usersCursor.All(ctx, &prijave); err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	return prijave, nil
}

func (pr *ApplyCompetitionRepo) GetPrijavaById(id string) (*Prijava, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	competitionCollection := pr.getCollectionCompetitionApply()

	var prijava Prijava
	objID, _ := primitive.ObjectIDFromHex(id)
	err := competitionCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&prijava)
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	return &prijava, nil
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

func (pr *ApplyCompetitionRepo) ApplyForCompetition(competitionID string, prijava *Prijava) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	competitionsCollection := pr.getCollectionCompetitionApply()

	deteCollection := pr.getCollectionDete()

	prijava.ID = primitive.NewObjectID()
	prijava.CompetitionID, _ = primitive.ObjectIDFromHex(competitionID) //proveriti
	//zdravstvenoStanje := client.getZS(jmbg)
	prijava.Dete.ID = primitive.NewObjectID()

	dete := Dete{
		ID:            prijava.Dete.ID,
		JMBG:          prijava.Dete.JMBG,
		DatumRodjenja: prijava.Dete.DatumRodjenja,
		Ime:           prijava.Dete.Ime,
		Prezime:       prijava.Dete.Prezime,
		Opstina:       prijava.Dete.Opstina,
		Adresa:        prijava.Dete.Adresa,
	}

	result, eerr := deteCollection.InsertOne(ctx, &dete)
	if eerr != nil {
		pr.logger.Println(eerr)
		return eerr
	}

	result, err := competitionsCollection.InsertOne(ctx, &prijava)
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

func (pr *ApplyCompetitionRepo) getCollectionCompetitionApply() *mongo.Collection {
	competitionDatabase := pr.cli.Database("mongodb")
	competitionsCollection := competitionDatabase.Collection("prijava")
	return competitionsCollection
}

func (pr *ApplyCompetitionRepo) getCollectionDete() *mongo.Collection {
	competitionDatabase := pr.cli.Database("mongodb")
	competitionsCollection := competitionDatabase.Collection("dete")
	return competitionsCollection
}
