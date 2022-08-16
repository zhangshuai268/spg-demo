package store

import (
	"context"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
	"github.com/zhangshuai268/spg-go-pkg/pkg/xorm"
	"spg-demo/internal/model"
)

type AdminStore interface {
	Create(ctx context.Context, admin *model.SpgAdmin) (int, error)
	Get(ctx context.Context, admin *model.SpgAdmin) (*model.SpgAdmin, bool, error)
	Update(ctx context.Context, admin *model.SpgAdmin) (int, error)
	Find(ctx context.Context, admin map[string]interface{}) ([]*model.SpgAdmin, int, error)
}

type admins struct {
	xorm *xorm.Engine
}

func (a *admins) Create(ctx context.Context, admin *model.SpgAdmin) (int, error) {
	_, err := a.xorm.Orm.Insert(admin)
	if err != nil {
		return 0, err
	}
	return admin.Id, nil
}

func (a *admins) Get(ctx context.Context, admin *model.SpgAdmin) (*model.SpgAdmin, bool, error) {
	get, err := a.xorm.Orm.Where("del = ?", 0).Get(admin)
	if err != nil {
		return nil, false, err
	}
	return admin, get, nil
}

func (a *admins) Update(ctx context.Context, admin *model.SpgAdmin) (int, error) {
	_, err := a.xorm.Orm.Id(admin.Id).Update(admin)
	if err != nil {
		return 0, err
	}
	return admin.Id, nil
}

func (a *admins) Find(ctx context.Context, admin map[string]interface{}) ([]*model.SpgAdmin, int, error) {
	type search struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
	}
	var s search
	err := util.StructTo(&admin, &s)
	if err != nil {
		return nil, 0, err
	}
	res := make([]*model.SpgAdmin, 0)

	session := a.xorm.Orm.Where("del = ?", 0)
	if s.Page != 0 && s.Limit != 0 {
		session.Limit(s.Limit, (s.Page-1)*s.Limit)
	}

	count, err := session.FindAndCount(&res)
	if err != nil {
		return nil, 0, err
	}
	return res, int(count), nil
}

func NewAdminStore(ds *datastore) AdminStore {
	return &admins{
		xorm: ds.xorm,
	}
}
