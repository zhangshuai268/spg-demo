package store

import (
	"context"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
	"github.com/zhangshuai268/spg-go-pkg/pkg/xorm"
	"spg-demo/internal/model"
)

type UserStore interface {
	Create(ctx context.Context, user *model.SpgUser) (int, error)
	Get(ctx context.Context, user *model.SpgUser) (*model.SpgUser, bool, error)
	Update(ctx context.Context, user *model.SpgUser) (int, error)
	Find(ctx context.Context, user map[string]interface{}) ([]*model.SpgUser, int, error)
}

type users struct {
	xorm *xorm.Engine
}

func (u *users) Create(ctx context.Context, user *model.SpgUser) (int, error) {
	_, err := u.xorm.Orm.Insert(user)
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (u *users) Get(ctx context.Context, user *model.SpgUser) (*model.SpgUser, bool, error) {
	get, err := u.xorm.Orm.Where("del = ?", 0).Get(user)
	if err != nil {
		return nil, false, err
	}
	return user, get, nil
}

func (u *users) Update(ctx context.Context, user *model.SpgUser) (int, error) {
	_, err := u.xorm.Orm.Update(user)
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (u *users) Find(ctx context.Context, user map[string]interface{}) ([]*model.SpgUser, int, error) {
	type search struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
	}
	var s search
	err := util.StructTo(&user, &s)
	if err != nil {
		return nil, 0, err
	}
	res := make([]*model.SpgUser, 0)

	session := u.xorm.Orm.Where("del = ?", 0)
	if s.Page != 0 && s.Limit != 0 {
		session.Limit(s.Limit, (s.Page-1)*s.Limit)
	}

	count, err := session.FindAndCount(&res)
	if err != nil {
		return nil, 0, err
	}
	return res, int(count), nil
}

func NewUserStore(ds *datastore) UserStore {
	return &users{
		xorm: ds.xorm,
	}
}
