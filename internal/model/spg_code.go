package model

import (
	"github.com/zhangshuai268/spg-go-pkg/pkg/myTime"
)

type SpgCode struct {
	Id         int           `json:"id" xorm:"not null pk autoincr INT(11)"`
	Code       string        `json:"code" xorm:"not null VARCHAR(100)"`
	AddTime    myTime.MyTime `json:"add_time" xorm:"not null created DATETIME"`
	UpdateTime myTime.MyTime `json:"update_time" xorm:"not null updated DATETIME"`
	DelTime    myTime.MyTime `json:"del_time" xorm:"not null DATETIME"`
	Del        int           `json:"del" xorm:"not null TINYINT(4)"`
	Mobile     string        `json:"mobile" xorm:"not null VARCHAR(100)"`
}
