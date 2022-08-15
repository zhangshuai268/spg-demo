package store

import (
	"context"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
	"github.com/zhangshuai268/spg-go-pkg/pkg/xorm"
	"spg-demo/internal/model"
)

type CodeStore interface {
	Create(ctx context.Context, code *model.SpgCode) (int, error)
	Get(ctx context.Context, code *model.SpgCode) (*model.SpgCode, bool, error)
	Find(ctx context.Context, code map[string]interface{}, desc string) ([]*model.SpgCode, int, error)
}

type codeStore struct {
	xorm *xorm.Engine
}

func (c *codeStore) Find(ctx context.Context, code map[string]interface{}, desc string) ([]*model.SpgCode, int, error) {
	type search struct {
		Page   int    `json:"page"`
		Limit  int    `json:"limit"`
		Mobile string `json:"mobile"`
	}
	var s search
	err := util.StructTo(&code, &s)
	if err != nil {
		return nil, 0, err
	}
	res := make([]*model.SpgCode, 0)

	session := c.xorm.Orm.Where("del = ?", 0)
	if s.Page != 0 && s.Limit != 0 {
		session.Limit(s.Limit, (s.Page-1)*s.Limit)
	}
	if desc != "" {
		session.Desc(desc)
	}
	count, err := session.FindAndCount(&res)
	if err != nil {
		return nil, 0, err
	}
	return res, int(count), nil
}

func (c *codeStore) Get(ctx context.Context, code *model.SpgCode) (*model.SpgCode, bool, error) {
	get, err := c.xorm.Orm.Where("del = ?", 0).Get(code)
	if err != nil {
		return nil, false, err
	}
	return code, get, nil
}

func (c *codeStore) Create(ctx context.Context, code *model.SpgCode) (int, error) {
	_, err := c.xorm.Orm.Insert(code)
	if err != nil {
		return 0, err
	}
	return code.Id, nil
}

func NewCodeStore(ds *datastore) CodeStore {
	return &codeStore{
		xorm: ds.xorm,
	}
}
