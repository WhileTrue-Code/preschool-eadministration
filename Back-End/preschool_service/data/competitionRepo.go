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

type Vrtici []*Vrtic

type Prijave []Prijava

func New(ctx context.Context, logger *log.Logger) (*ApplyCompetitionRepo, error) {
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

	//prijaveCollection := client.Database("mongodb").Collection("prijava")

	return &ApplyCompetitionRepo{
		cli:    client,
		logger: logger,
	}, nil
}

//func NewApplyCompetitionRepo(logger *log.Logger) (*ApplyCompetitionRepo, error) {
//	db := os.Getenv("PRESCHOOL_DB_HOST")
//	dbport := os.Getenv("PRESCHOOL_DB_PORT")
//
//	host := fmt.Sprintf("%s:%s", db, dbport)
//	client, err := mongo.NewClient(options.Client().ApplyURI(`mongodb://` + host))
//	if err != nil {
//		return nil, err
//	}
//
//	err = client.Connect(context.Background())
//	if err != nil {
//		return nil, err
//	}
//
//	//prijaveCollection := client.Database("mongodb").Collection("prijava")
//
//	return &ApplyCompetitionRepo{
//		cli:    client,
//		logger: logger,
//	}, nil
//
//	//return &ApplyCompetitionRepo{
//	//	cli:     client,
//	//	logger:  logger,
//	//	prijave: prijaveCollection,
//	//}, nil
//}

func (pr *ApplyCompetitionRepo) GetAll() (Competitions, error) {
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

func (pr *ApplyCompetitionRepo) GetAllVrtici() (Vrtici, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	vrticiCollection := pr.getCollectionVrtic()

	var vrtici Vrtici
	usersCursor, err := vrticiCollection.Find(ctx, bson.M{})
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	if err = usersCursor.All(ctx, &vrtici); err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	return vrtici, nil
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

func (pr *ApplyCompetitionRepo) GetById(id string) (*Competition, error) {
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

type PrijaveRepositoryImpl struct {
	prijave *mongo.Collection //izvodi
}

//func (store *ApplyCompetitionRepo) Get(compId string) (prijave Prijave) {
//	prijave = Prijave{}
//	var filter interface{}
//	competitionID, _ := primitive.ObjectIDFromHex(compId)
//
//	filter = bson.M{"_idCompetition": competitionID}
//
//	pronadjenePrijave, err := store.prijave.Find(context.Background(), filter)
//	if err != nil {
//		return nil
//	}
//
//	//pronadjenePrijave.Decode(&prijave)
//	var found []Prijava
//	err = pronadjenePrijave.Decode(&found)
//	if err != nil {
//		log.Println(err)
//	}
//	return found
//}

func (pr *ApplyCompetitionRepo) GetVrticById(id string) (*Vrtic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	vrticCollection := pr.getCollectionVrtic()

	var vrtic Vrtic
	objID, _ := primitive.ObjectIDFromHex(id)
	err := vrticCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&vrtic)
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	return &vrtic, nil
}

func (pr *ApplyCompetitionRepo) ChangeStatus(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	compCollection := pr.getCollection()
	prijavaCollection := pr.getCollectionCompetitionApply()

	// What happens if set value for index=10, but we only have 3 phone numbers?
	// -> Every value in between will be set to an empty string
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.M{"$set": bson.M{
		"status": "Zatvoren",
	}}

	var competition Competition
	err := compCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&competition)
	if err != nil {
		pr.logger.Println(err)
		return nil
	}

	prijave, _ := pr.GetAllApplyesForOneCompetition(id)

	for i, prijava := range prijave {
		if i < competition.BrojDece {
			filter1 := bson.D{{Key: "_id", Value: prijava.ID}}
			update1 := bson.M{"$set": bson.M{
				"status": "Upisan",
			}}
			prijavaCollection.UpdateOne(ctx, filter1, update1)
		} else {
			filter2 := bson.D{{Key: "_id", Value: prijava.ID}}
			update2 := bson.M{"$set": bson.M{
				"status": "Odbijen",
			}}
			prijavaCollection.UpdateOne(ctx, filter2, update2)
		}

	}

	result, err := compCollection.UpdateOne(ctx, filter, update)
	pr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	pr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		pr.logger.Println(err)
		return err
	}
	return nil
}

func (pr *ApplyCompetitionRepo) GetAllApplyesForOneCompetition(compId string) (prijave Prijave, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	competitionCollection := pr.getCollectionCompetitionApply()

	//var konkurs Competition
	var prijave2 Prijave
	competitionID, _ := primitive.ObjectIDFromHex(compId)

	for _, prijava := range prijave {
		if prijava.CompetitionID == competitionID {
			prijave2 = append(prijave2, prijava)
		}
	}
	fmt.Println(prijave2)
	opts := options.Find().SetSort(bson.D{{"bodovi", -1}})
	usersCursor, err := competitionCollection.Find(ctx, bson.M{"_idCompetition": competitionID}, opts)
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	if err = usersCursor.All(ctx, &prijave2); err != nil {
		pr.logger.Println(err)
		return nil, err
	}

	return prijave2, nil
}

func (pr *ApplyCompetitionRepo) PostCompetition(vrticID string, competition *Competition) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	competitionsCollection := pr.getCollection()

	competition.ID = primitive.NewObjectID()
	competition.Vrtic, _ = pr.GetVrticById(vrticID)
	competition.Status = "Otvoren"

	result, err := competitionsCollection.InsertOne(ctx, &competition)
	if err != nil {
		pr.logger.Println(err)
		return err
	}
	pr.logger.Printf("Documents ID: %v\n", result.InsertedID)

	return nil
}

func (pr *ApplyCompetitionRepo) PostVrtic(vrtic *Vrtic) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	vrticCollection := pr.getCollectionVrtic()

	vrtic.ID = primitive.NewObjectID()

	result, err := vrticCollection.InsertOne(ctx, &vrtic)
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
	prijava.Status = "Prijavljen"

	dete := Dete{
		ID:                prijava.Dete.ID,
		JMBG:              prijava.Dete.JMBG,
		DatumRodjenja:     prijava.Dete.DatumRodjenja,
		Ime:               prijava.Dete.Ime,
		Prezime:           prijava.Dete.Prezime,
		Opstina:           prijava.Dete.Opstina,
		Adresa:            prijava.Dete.Adresa,
		ZdravstvenoStanje: prijava.Dete.ZdravstvenoStanje,
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

//func (pr *ApplyCompetitionRepo) UpdateCompetitionStatus(id string, competition Competition) error {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	patientsCollection := pr.getCollection()
//
//	objID, _ := primitive.ObjectIDFromHex(id)
//	filter := bson.M{"_id": objID}
//	update := bson.M{"$set": bson.M{
//		"status": competition.Status,
//	}}
//	result, err := patientsCollection.UpdateOne(ctx, filter, update)
//	pr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
//	pr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)
//
//	if err != nil {
//		pr.logger.Println(err)
//		return err
//	}
//	return nil
//}

func (pr *ApplyCompetitionRepo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	compCollection := pr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := compCollection.DeleteOne(ctx, filter)
	if err != nil {
		pr.logger.Println(err)
		return err
	}
	pr.logger.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}

// Disconnect from database
func (pr *ApplyCompetitionRepo) Disconnect(ctx context.Context) error {
	err := pr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Check database connection
func (pr *ApplyCompetitionRepo) Ping() {
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

func (pr *ApplyCompetitionRepo) getCollection() *mongo.Collection {
	competitionDatabase := pr.cli.Database("mongodb")
	competitionsCollection := competitionDatabase.Collection("competitions")
	return competitionsCollection
}

func (pr *ApplyCompetitionRepo) getCollectionVrtic() *mongo.Collection {
	competitionDatabase := pr.cli.Database("mongodb")
	competitionsCollection := competitionDatabase.Collection("vrtic")
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
