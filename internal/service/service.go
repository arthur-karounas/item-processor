package service

import "item-processor/internal/repository"

type Service struct {
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
