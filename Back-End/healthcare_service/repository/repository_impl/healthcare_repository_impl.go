package repository_impl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"healthcare_service/model"
	"healthcare_service/repository"
	"log"
)

type HealthcareRepositoryImpl struct {
	appointment       *mongo.Collection
	vaccination       *mongo.Collection
	zdravstvenoStanje *mongo.Collection
}

const (
	DATABASE                      = "healthcare"
	COLLECTION_APPOINTMENT        = "appointment"
	COLLECTION_VACCINATION        = "vaccination"
	COLLECTION_ZDRAVSTVENO_STANJE = "zdravstvenoStanje"
)

func NewAuthRepositoryImpl(client *mongo.Client) repository.HealthcareRepository {
	appointment := client.Database(DATABASE).Collection(COLLECTION_APPOINTMENT)
	vaccination := client.Database(DATABASE).Collection(COLLECTION_VACCINATION)
	zdravstvenoStanje := client.Database(DATABASE).Collection(COLLECTION_ZDRAVSTVENO_STANJE)

	return &HealthcareRepositoryImpl{
		appointment:       appointment,
		vaccination:       vaccination,
		zdravstvenoStanje: zdravstvenoStanje,
	}
}

//Appointments
func (repository *HealthcareRepositoryImpl) GetAllAppointments() ([]*model.Appointment, error) {
	filter := bson.M{}
	return repository.filterAppointments(filter)
}

func (repository *HealthcareRepositoryImpl) GetMyAppointmentsDoctor(id primitive.ObjectID) ([]*model.Appointment, error) {
	filter := bson.M{"doctor._id": id}
	return repository.filterAppointments(filter)
}

func (repository *HealthcareRepositoryImpl) GetMyAvailableAppointmentsDoctor(id primitive.ObjectID) ([]*model.Appointment, error) {
	filter := bson.M{"doctor._id": id, "user": nil}
	return repository.filterAppointments(filter)
}

func (repository *HealthcareRepositoryImpl) GetMyTakenAppointmentsDoctor(id primitive.ObjectID) ([]*model.Appointment, error) {
	filter := bson.M{"doctor._id": id, "user": bson.M{"$ne": nil}}
	return repository.filterAppointments(filter)
}

func (repository *HealthcareRepositoryImpl) GetAllAvailableAppointments() ([]*model.Appointment, error) {
	filter := bson.M{"user": nil}
	return repository.filterAppointments(filter)
}

func (repository *HealthcareRepositoryImpl) GetAppointmentByID(id primitive.ObjectID) (*model.Appointment, error) {
	filter := bson.M{"_id": id}
	return repository.filterOneAppointment(filter)
}

func (repository *HealthcareRepositoryImpl) CreateNewAppointment(appointment *model.Appointment) error {
	_, err := repository.appointment.InsertOne(context.Background(), appointment)
	if err != nil {
		return err
	}
	return nil
}

func (repository *HealthcareRepositoryImpl) SetAppointment(appointment *model.Appointment) error {
	filter := bson.M{"_id": appointment.ID}
	update := bson.D{{"$set", appointment}}
	_, err := repository.appointment.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Updating Appointment Error MongoDB", err.Error())
		return err
	}

	return nil
}

func (repository *HealthcareRepositoryImpl) DeleteAppointmentByID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := repository.appointment.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (repository *HealthcareRepositoryImpl) filterAppointments(filter interface{}) ([]*model.Appointment, error) {
	cursor, err := repository.appointment.Find(context.Background(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeAppointment(cursor)
}

func (repository *HealthcareRepositoryImpl) filterOneAppointment(filter interface{}) (appointment *model.Appointment, err error) {
	result := repository.appointment.FindOne(context.Background(), filter)
	err = result.Decode(&appointment)
	return
}

//Vacinations
func (repository *HealthcareRepositoryImpl) GetAllVaccinations() ([]*model.Vaccination, error) {
	filter := bson.D{{}}
	return repository.filterVaccinations(filter)
}

func (repository *HealthcareRepositoryImpl) GetMyVaccinationsDoctor(id primitive.ObjectID) ([]*model.Vaccination, error) {
	filter := bson.M{"doctor._id": id}
	return repository.filterVaccinations(filter)
}

func (repository *HealthcareRepositoryImpl) GetMyAvailableVaccinationsDoctor(id primitive.ObjectID) ([]*model.Vaccination, error) {
	filter := bson.M{"doctor._id": id, "user": nil}
	return repository.filterVaccinations(filter)
}

func (repository *HealthcareRepositoryImpl) GetMyTakenVaccinationsDoctor(id primitive.ObjectID) ([]*model.Vaccination, error) {
	filter := bson.M{"doctor._id": id, "user": bson.M{"$ne": nil}}
	return repository.filterVaccinations(filter)
}

func (repository *HealthcareRepositoryImpl) GetAllAvailableVaccinations() ([]*model.Vaccination, error) {
	filter := bson.M{"user": nil}
	return repository.filterVaccinations(filter)
}

func (repository *HealthcareRepositoryImpl) GetMyTakenVaccinationsRegular(id primitive.ObjectID) ([]*model.Vaccination, error) {
	filter := bson.M{"user._id": id}
	return repository.filterVaccinations(filter)
}

func (repository *HealthcareRepositoryImpl) GetVaccinationByID(id primitive.ObjectID) (*model.Vaccination, error) {
	filter := bson.M{"_id": id}
	return repository.filterOneVaccination(filter)
}

func (repository *HealthcareRepositoryImpl) CreateNewVaccination(vaccination *model.Vaccination) error {
	_, err := repository.vaccination.InsertOne(context.Background(), vaccination)
	if err != nil {
		return err
	}
	return nil
}

func (repository *HealthcareRepositoryImpl) SetVaccination(vaccination *model.Vaccination) error {
	filter := bson.M{"_id": vaccination.ID}
	update := bson.D{{"$set", vaccination}}
	_, err := repository.vaccination.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Updating Vaccination Error MongoDB", err.Error())
		return err
	}

	return nil
}

func (repository *HealthcareRepositoryImpl) DeleteVaccinationByID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := repository.vaccination.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (repository *HealthcareRepositoryImpl) filterVaccinations(filter interface{}) ([]*model.Vaccination, error) {
	cursor, err := repository.vaccination.Find(context.Background(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeVaccination(cursor)
}

func (repository *HealthcareRepositoryImpl) filterOneVaccination(filter interface{}) (appointment *model.Vaccination, err error) {
	result := repository.vaccination.FindOne(context.Background(), filter)
	err = result.Decode(&appointment)
	return
}

//ZdravstvenoStanje
func (repository *HealthcareRepositoryImpl) GetAllZdravstvenoStanje() ([]*model.ZdravstvenoStanje, error) {
	filter := bson.D{{}}
	return repository.filterZdravstvenaStanja(filter)
}

func (repository *HealthcareRepositoryImpl) GetZdravstvenoStanjeByID(id primitive.ObjectID) (*model.ZdravstvenoStanje, error) {
	filter := bson.M{"_id": id}
	return repository.filterOneZdravstvenoStanje(filter)
}

func (repository *HealthcareRepositoryImpl) GetZdravstvenoStanjeByJMBG(jmbg string) (*model.ZdravstvenoStanje, error) {
	filter := bson.M{"jmbg": jmbg}
	return repository.filterOneZdravstvenoStanje(filter)
}

func (repository *HealthcareRepositoryImpl) CreateNewZdravstvenoStanje(zdravstvenoStanje *model.ZdravstvenoStanje) error {
	_, err := repository.zdravstvenoStanje.InsertOne(context.Background(), zdravstvenoStanje)
	if err != nil {
		return err
	}
	return nil
}

func (repository *HealthcareRepositoryImpl) DeleteZdravstvenoStanjeByJMBG(jmbg string) error {
	filter := bson.M{"jmbg": jmbg}
	_, err := repository.zdravstvenoStanje.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
func (repository *HealthcareRepositoryImpl) filterZdravstvenaStanja(filter interface{}) ([]*model.ZdravstvenoStanje, error) {
	cursor, err := repository.zdravstvenoStanje.Find(context.Background(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeZdravstvenoStanje(cursor)
}

func (repository *HealthcareRepositoryImpl) filterOneZdravstvenoStanje(filter interface{}) (zdravstvenoStanje *model.ZdravstvenoStanje, err error) {
	result := repository.zdravstvenoStanje.FindOne(context.Background(), filter)
	err = result.Decode(&zdravstvenoStanje)
	return
}

func decodeAppointment(cursor *mongo.Cursor) (appointments []*model.Appointment, err error) {
	for cursor.Next(context.Background()) {
		var appointment model.Appointment
		err = cursor.Decode(&appointment)
		if err != nil {
			return
		}
		appointments = append(appointments, &appointment)
	}
	err = cursor.Err()
	return
}

func decodeVaccination(cursor *mongo.Cursor) (vaccinations []*model.Vaccination, err error) {
	for cursor.Next(context.Background()) {
		var vaccination model.Vaccination
		err = cursor.Decode(&vaccination)
		if err != nil {
			return
		}
		vaccinations = append(vaccinations, &vaccination)
	}
	err = cursor.Err()
	return
}

func decodeZdravstvenoStanje(cursor *mongo.Cursor) (zdravstvenaStanja []*model.ZdravstvenoStanje, err error) {
	for cursor.Next(context.Background()) {
		var zdravstvenoStanje model.ZdravstvenoStanje
		err = cursor.Decode(&zdravstvenoStanje)
		if err != nil {
			return
		}
		zdravstvenaStanja = append(zdravstvenaStanja, &zdravstvenoStanje)
	}
	err = cursor.Err()
	return
}
