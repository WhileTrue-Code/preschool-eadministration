package service

import "registrar_service/repository"

type RegistrarService struct {
	store repository.RegistrarRepository
}

func NewRegistrarService(store repository.RegistrarRepository) *RegistrarService {
	return &RegistrarService{
		store: store,
	}
}
