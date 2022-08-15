package model

import (
	"github.com/zhangshuai268/spg-go-pkg/pkg/myTime"
)

type SpgBookDetail struct {
	Id         int           `json:"id" xorm:"not null pk autoincr INT(11)"`
	BookId     int           `json:"book_id" xorm:"not null INT(11)"`
	BookNo     string        `json:"book_no" xorm:"not null comment('图书编号') VARCHAR(100)"`
	IsBorrow   int           `json:"is_borrow" xorm:"not null comment('借出状态 1借出 2未借出') TINYINT(4)"`
	AddTime    myTime.MyTime `json:"add_time" xorm:"not null created DATETIME"`
	UpdateTime myTime.MyTime `json:"update_time" xorm:"not null updated DATETIME"`
	DelTime    myTime.MyTime `json:"del_time" xorm:"not null DATETIME"`
	Del        int           `json:"del" xorm:"not null TINYINT(4)"`
}
