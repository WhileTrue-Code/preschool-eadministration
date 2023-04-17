package service

import "registrar_service/repository"

type AuthService struct {
	store repository.AuthRepository
}

func NewAuthService(store repository.AuthRepository) *AuthService {
	return &AuthService{
		store: store,
	}
}
