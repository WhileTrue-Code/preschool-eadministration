package repository

import domain "registrar_service/model/entity"

type RegistrarRepository interface {
	CreateNewBirthCertificate(user domain.User) error
	IsUserExist(jmbg string) bool
	FindOneUser(jmbg string) *domain.User
	CreateNewMarriage(marriage domain.ExcerptFromTheMarriageRegister)
	UpdateCertificate(user domain.User) error
	GetChildren(jmbg string, pol domain.Pol) []domain.User
}
