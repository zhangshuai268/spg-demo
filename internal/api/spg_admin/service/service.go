package service

import (
	"spg-demo/internal/api/store"
)

type Service interface {
	Admin() AdminService
}

type service struct {
	factory store.Factory
}

func (s *service) Admin() AdminService {
	return NewAdminService(s)
}

func NewService(factory store.Factory) Service {
	return &service{
		factory: factory,
	}
}
