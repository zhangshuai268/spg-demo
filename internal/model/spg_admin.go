package models

import (
	"github.com/zhangshuai268/spg-go-pkg/pkg/myTime"
)

type SpgAdmin struct {
	Id         int           `json:"id" xorm:"not null pk INT(11)"`
	Mobile     string        `json:"mobile" xorm:"not null VARCHAR(100)"`
	Title      string        `json:"title" xorm:"not null VARCHAR(100)"`
	AddTime    myTime.MyTime `json:"add_time" xorm:"not null created DATETIME"`
	UpdateTime myTime.MyTime `json:"update_time" xorm:"not null updated DATETIME"`
	DelTime    myTime.MyTime `json:"del_time" xorm:"not null DATETIME"`
	Del        int           `json:"del" xorm:"not null TINYINT(4)"`
}
