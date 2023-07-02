package service

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"healthcare_service/model"
	"healthcare_service/repository"
	"log"
	"os"
	"time"
)

type HealthcareService struct {
	repository     repository.HealthcareRepository
	natsConnection *nats.Conn
}

func NewHealthcareService(repository repository.HealthcareRepository, natsConnection *nats.Conn) *HealthcareService {
	return &HealthcareService{
		repository:     repository,
		natsConnection: natsConnection,
	}
}

func (service *HealthcareService) GetAllAppointments() ([]*model.Appointment, error) {
	return service.repository.GetAllAppointments()
}

func (service *HealthcareService) GetMyAppointmentsDoctor(jmbg string) ([]*model.Appointment, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var doctor model.User
	err = json.Unmarshal(response.Data, &doctor)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return nil, err
	}

	doctorID := doctor.ID

	return service.repository.GetMyAppointmentsDoctor(doctorID)
}

func (service *HealthcareService) GetMyAvailableAppointmentsDoctor(jmbg string) ([]*model.Appointment, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var doctor model.User
	err = json.Unmarshal(response.Data, &doctor)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return nil, err
	}

	doctorID := doctor.ID

	return service.repository.GetMyAvailableAppointmentsDoctor(doctorID)
}

func (service *HealthcareService) GetMyTakenAppointmentsDoctor(jmbg string) ([]*model.Appointment, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var doctor model.User
	err = json.Unmarshal(response.Data, &doctor)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return nil, err
	}

	doctorID := doctor.ID

	return service.repository.GetMyTakenAppointmentsDoctor(doctorID)
}

func (service *HealthcareService) GetAllAvailableAppointments() ([]*model.Appointment, error) {
	return service.repository.GetAllAvailableAppointments()
}

func (service *HealthcareService) GetAppointmentByID(id primitive.ObjectID) (*model.Appointment, error) {
	return service.repository.GetAppointmentByID(id)
}

func (service *HealthcareService) GetMe(jmbg string) (*model.User, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var user model.User
	err = json.Unmarshal(response.Data, &user)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return nil, err
	}

	return &user, nil
}

func (service *HealthcareService) CreateNewAppointment(appointment *model.Appointment, jmbg string) (int, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	existingAppointments, err := service.repository.GetAllAppointments()
	for _, existingAppointment := range existingAppointments {
		if (existingAppointment.StartOfAppointment >= appointment.StartOfAppointment && existingAppointment.StartOfAppointment <= appointment.EndOfAppointment) ||
			(existingAppointment.EndOfAppointment >= appointment.StartOfAppointment && existingAppointment.EndOfAppointment <= appointment.EndOfAppointment) ||
			(existingAppointment.StartOfAppointment >= appointment.StartOfAppointment && existingAppointment.EndOfAppointment <= appointment.EndOfAppointment) {
			return 1, nil
		}
	}
	if err != nil {
		log.Println("Error getting all Appointments", err)
		return 0, err
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var doctor model.User
	err = json.Unmarshal(response.Data, &doctor)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return 0, err
	}

	appointment.ID = primitive.NewObjectID()
	appointment.Doctor = &doctor
	appointment.User = nil

	err = service.repository.CreateNewAppointment(appointment)
	if err != nil {
		log.Println("Error in trying to save Appointment")
		return 0, err
	}

	return 0, nil
}

func (service *HealthcareService) SetAppointment(id primitive.ObjectID, jmbg string) (*model.Appointment, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var user model.User
	err = json.Unmarshal(response.Data, &user)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return nil, err
	}

	appointment, err := service.repository.GetAppointmentByID(id)
	if err != nil {
		log.Println("Error in finding Appointment By ID")
		return nil, err
	}

	appointment.User = &user

	err = service.repository.SetAppointment(appointment)
	if err != nil {
		log.Println("Error in Updating Appointment")
		return nil, err
	}

	return appointment, nil
}

func (service *HealthcareService) DeleteAppointmentByID(id primitive.ObjectID) error {
	return service.repository.DeleteAppointmentByID(id)
}

func (service *HealthcareService) GetAllVaccinations() ([]*model.Vaccination, error) {
	return service.repository.GetAllVaccinations()
}

func (service *HealthcareService) GetMyVaccinationsDoctor(jmbg string) ([]*model.Vaccination, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var doctor model.User
	err = json.Unmarshal(response.Data, &doctor)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return nil, err
	}

	doctorID := doctor.ID

	return service.repository.GetMyVaccinationsDoctor(doctorID)
}

func (service *HealthcareService) GetMyAvailableVaccinationsDoctor(jmbg string) ([]*model.Vaccination, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var doctor model.User
	err = json.Unmarshal(response.Data, &doctor)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return nil, err
	}

	doctorID := doctor.ID

	return service.repository.GetMyAvailableVaccinationsDoctor(doctorID)
}

func (service *HealthcareService) GetMyTakenVaccinationsDoctor(jmbg string) ([]*model.Vaccination, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var doctor model.User
	err = json.Unmarshal(response.Data, &doctor)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return nil, err
	}

	doctorID := doctor.ID

	return service.repository.GetMyTakenVaccinationsDoctor(doctorID)
}

func (service *HealthcareService) GetMyTakenVaccinationsRegular(jmbg string) ([]*model.Vaccination, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var user model.User
	err = json.Unmarshal(response.Data, &user)
	if err != nil {
		log.Print("Error in Unmarshalling JSON")
		return nil, err
	}

	userID := user.ID

	return service.repository.GetMyTakenVaccinationsRegular(userID)
}

func (service *HealthcareService) GetAllAvailableVaccinations() ([]*model.Vaccination, error) {
	return service.repository.GetAllAvailableVaccinations()
}

func (service *HealthcareService) GetVaccinationByID(id primitive.ObjectID) (*model.Vaccination, error) {
	return service.repository.GetVaccinationByID(id)
}

func (service *HealthcareService) CreateNewVaccination(vaccination *model.Vaccination, jmbg string) (int, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}
	existingVaccinations, err := service.repository.GetAllVaccinations()
	for _, existingVaccination := range existingVaccinations {
		if (existingVaccination.StartOfVaccination >= vaccination.StartOfVaccination && existingVaccination.StartOfVaccination <= vaccination.EndOfVaccination) ||
			(existingVaccination.EndOfVaccination >= vaccination.StartOfVaccination && existingVaccination.EndOfVaccination <= vaccination.EndOfVaccination) ||
			(existingVaccination.StartOfVaccination >= vaccination.StartOfVaccination && existingVaccination.EndOfVaccination <= vaccination.EndOfVaccination) {
			return 1, nil
		}
	}
	if err != nil {
		log.Println("Error getting All Vaccinations", err)
		return 0, err
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var doctor model.User
	err = json.Unmarshal(response.Data, &doctor)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return 0, err
	}

	vaccination.ID = primitive.NewObjectID()
	vaccination.Doctor = &doctor
	vaccination.User = nil

	err = service.repository.CreateNewVaccination(vaccination)
	if err != nil {
		log.Println("Error in trying to save Vaccination")
		return 0, err
	}
	return 0, nil
}

func (service *HealthcareService) SetVaccination(id primitive.ObjectID, jmbg string) (*model.Vaccination, error) {
	dataToSend, err := json.Marshal(jmbg)
	if err != nil {
		log.Println("Error Marshaling JMBG")
	}

	response, err := service.natsConnection.Request(os.Getenv("GET_USER_BY_JMBG"), dataToSend, 5*time.Second)

	var user model.User
	err = json.Unmarshal(response.Data, &user)
	if err != nil {
		log.Println("Error in Unmarshalling json")
		return nil, err
	}

	vaccination, err := service.repository.GetVaccinationByID(id)
	if err != nil {
		log.Println("Error in finding Vaccination By ID")
		return nil, err
	}

	vaccination.User = &user

	err = service.repository.SetVaccination(vaccination)
	if err != nil {
		log.Println("Error in Updating Vaccination")
		return nil, err
	}

	return vaccination, nil
}

func (service *HealthcareService) DeleteVaccinationByID(id primitive.ObjectID) error {
	return service.repository.DeleteVaccinationByID(id)
}

func (service *HealthcareService) GetAllZdravstvenoStanje() ([]*model.ZdravstvenoStanje, error) {
	return service.repository.GetAllZdravstvenoStanje()
}

func (service *HealthcareService) GetZdravstvenoStanjeByID(id primitive.ObjectID) (*model.ZdravstvenoStanje, error) {
	return service.repository.GetZdravstvenoStanjeByID(id)
}

func (service *HealthcareService) GetZdravstvenoStanjeByJMBG(jmbg string) (*model.ZdravstvenoStanje, error) {
	return service.repository.GetZdravstvenoStanjeByJMBG(jmbg)
}

func (service *HealthcareService) GetMyZdravstvenoStanje(jmbg string) (*model.ZdravstvenoStanje, error) {
	return service.repository.GetZdravstvenoStanjeByJMBG(jmbg)
}

func (service *HealthcareService) CreateNewZdravstvenoStanje(zdravstvenoStanje *model.ZdravstvenoStanje) error {
	zdravstvenoStanje.ID = primitive.NewObjectID()
	err := service.repository.CreateNewZdravstvenoStanje(zdravstvenoStanje)
	if err != nil {
		log.Println("Error in trying to save Zdravstveno Stanje")
		return err
	}

	return nil
}

func (service *HealthcareService) DeleteZdravstvenoStanjeByJMBG(jmbg string) error {
	return service.repository.DeleteZdravstvenoStanjeByJMBG(jmbg)
}

func (service *HealthcareService) AddPersonToRegistry(user *model.User) (*model.User, int) {
	user.ID = primitive.NewObjectID()

	dataToSend, err := json.Marshal(user)
	if err != nil {
		log.Print("Error in Marshaling JSON")
		return nil, 0
	}

	response, err := service.natsConnection.Request(os.Getenv("CREATE_USER"), dataToSend, 5*time.Second)
	if err != nil {
		log.Println(err)
		return nil, 0
	}

	err = json.Unmarshal(response.Data, &user)
	if err != nil {
		log.Print("Error in Unmarshal JSON")
		return nil, 0
	}

	if user.ID == primitive.NilObjectID {
		return nil, 1
	}

	return user, 0
}

func (service *HealthcareService) SubscribeToNats(natsConnection *nats.Conn) {

	_, err := natsConnection.QueueSubscribe(os.Getenv("GET_STANJE_BY_JMBG"), "queue-healthcare-group", func(message *nats.Msg) {
		var jmbg string
		err := json.Unmarshal(message.Data, &jmbg)
		if err != nil {
			log.Println("Error in unmarshal JSON")
			return
		}

		zdravstvenoStanje, err := service.GetZdravstvenoStanjeByJMBG(jmbg)
		if err != nil {
			log.Println(err)
			return
		}

		dataToSend, err := json.Marshal(zdravstvenoStanje)
		if err != nil {
			log.Println("Error in marshaling JSON")
			return
		}
		reply := dataToSend
		err = natsConnection.Publish(message.Reply, reply)
		if err != nil {
			log.Println("Error in publishing response: %s", err.Error())
			return
		}
	})
	if err != nil {
		log.Print("Error in receiving message: %s", err.Error())
	}

	log.Printf("Subscribed to channel: %s", os.Getenv("GET_STANJE_BY_JMBG"))

}
