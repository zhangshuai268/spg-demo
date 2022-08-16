package store

import (
	"context"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
	"github.com/zhangshuai268/spg-go-pkg/pkg/xorm"
	"spg-demo/internal/model"
)

type BookStore interface {
	Create(ctx context.Context, book *model.SpgBook) (int, error)
	Get(ctx context.Context, book *model.SpgBook) (*model.SpgBook, bool, error)
	Update(ctx context.Context, book *model.SpgBook) (int, error)
	Find(ctx context.Context, book map[string]interface{}) ([]*model.SpgBook, int, error)
}

type bookStore struct {
	xorm *xorm.Engine
}

func (b *bookStore) Find(ctx context.Context, book map[string]interface{}) ([]*model.SpgBook, int, error) {
	type search struct {
		Page   int    `json:"page"`
		Limit  int    `json:"limit"`
		Title  string `json:"title"`
		Writer string `json:"writer"`
		Press  string `json:"press"`
	}
	var s search
	err := util.StructTo(&book, &s)
	if err != nil {
		return nil, 0, err
	}
	res := make([]*model.SpgBook, 0)
	session := b.xorm.Orm.Where("del = ?", 0)
	if s.Title != "" {
		session.Where("title like ?", "%"+s.Title+"%")
	}
	if s.Writer != "" {
		session.Where("writer like ?", "%"+s.Writer+"%")
	}
	if s.Press != "" {
		session.Where("press like ?", "%"+s.Press+"%")
	}
	if s.Page != 0 && s.Limit != 0 {
		session.Limit(s.Limit, (s.Page-1)*s.Limit)
	}
	count, err := session.FindAndCount(&res)
	if err != nil {
		return nil, 0, err
	}
	return res, int(count), nil
}

func (b *bookStore) Update(ctx context.Context, book *model.SpgBook) (int, error) {
	_, err := b.xorm.Orm.Id(book.Id).Update(book)
	if err != nil {
		return nil, err
	}
	return book.Id, nil
}

func (b *bookStore) Get(ctx context.Context, book *model.SpgBook) (*model.SpgBook, bool, error) {
	get, err := b.xorm.Orm.Where("del = ?", 0).Get(book)
	if err != nil {
		return nil, false, err
	}
	return book, get, nil
}

func (b *bookStore) Create(ctx context.Context, book *model.SpgBook) (int, error) {
	_, err := b.xorm.Orm.Insert(book)
	if err != nil {
		return 0, err
	}
	return book.Id, nil
}

func NewBookStore(ds *datastore) BookStore {
	return &bookStore{
		xorm: ds.xorm,
	}
}
