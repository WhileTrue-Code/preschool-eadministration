package service

import (
	domain "auth_service/model/entity"
	"auth_service/repository"
	"github.com/cristalhq/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

type AuthService struct {
	store repository.AuthRepository
}

func NewAuthService(store repository.AuthRepository) *AuthService {
	return &AuthService{
		store: store,
	}
}

func (service *AuthService) IsJMBGUnique(jmbg string) bool {
	return service.store.IsJMBGUnique(jmbg)
}

func (service *AuthService) SignUp(credentials domain.Credentials) (int, error) {

	isExists := service.IsJMBGUnique(credentials.JMBG)
	if isExists == true {
		return -1, nil
	}

	credentials.ID = primitive.NewObjectID()                                                       //creating unique UUID for MongoDB
	password, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost) //hashing password
	credentials.Password = string(password)
	if err != nil {
		return 0, err
	}
	service.store.SignUp(credentials)

	return 0, nil
}

func (service *AuthService) Login(jmbg string, password string) (string, int) {

	credentials, err := service.store.GetCredentials(jmbg)
	if err != nil {
		log.Println(err)
		return "", 1
	}

	err = bcrypt.CompareHashAndPassword([]byte(credentials.Password), []byte(password))
	if err != nil {
		log.Println(err)
		return "", 2
	}

	tokenString, err := GenerateJWT(credentials)
	if err != nil {
		return "", 3
	}

	return tokenString, 0
}

func GenerateJWT(credentials *domain.Credentials) (string, error) {
	key := []byte(os.Getenv("SECRET_KEY"))
	signer, err := jwt.NewSignerHS(jwt.HS256, key)
	if err != nil {
		log.Println(err)
	}

	builder := jwt.NewBuilder(signer)

	claims := &domain.Claims{
		UserID:    credentials.ID,
		JMBG:      credentials.JMBG,
		Role:      credentials.UserType,
		ExpiresAt: time.Now().Add(time.Minute * 60),
	}

	token, err := builder.Build(claims)
	if err != nil {
		log.Println(err)
	}

	return token.String(), nil
}
