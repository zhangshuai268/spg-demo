package store

import "github.com/zhangshuai268/spg-go-pkg/pkg/xorm"

type AdminStore interface {
}

type admin struct {
	xorm *xorm.Engine
}

func NewAdminStore(ds *datastore) AdminStore {
	return &admin{
		xorm: ds.xorm,
	}
}
