package repository

import (
	domain "auth_service/model/entity"
)

type AuthRepository interface {
	IsJMBGUnique(jmbg string) bool
	SignUp(credentials domain.Credentials)
	GetCredentials(jmbg string) (*domain.Credentials, error)
}
