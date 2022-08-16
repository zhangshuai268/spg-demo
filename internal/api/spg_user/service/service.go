package service

import (
	"spg-demo/internal/api/store"
)

type Service interface {
	User() UserService
}

type service struct {
	factory store.Factory
}

func (s *service) User() UserService {
	return NewUserService(s)
}

func NewService(factory store.Factory) Service {
	return &service{
		factory: factory,
	}
}
