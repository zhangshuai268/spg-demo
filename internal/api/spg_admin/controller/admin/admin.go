package admin

import (
	"spg-demo/internal/api/spg_admin/service"
	"spg-demo/internal/api/store"
)

type ControllerAdmin struct {
	srv     service.Service
	factory store.Factory
}

func NewControllerAdmin(factory store.Factory) *ControllerAdmin {
	return &ControllerAdmin{
		srv:     service.NewService(factory),
		factory: factory,
	}
}
