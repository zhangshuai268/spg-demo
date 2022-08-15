package model

import (
	"github.com/zhangshuai268/spg-go-pkg/pkg/myTime"
)

type SpgBorrow struct {
	Id         int           `json:"id" xorm:"not null pk autoincr INT(11)"`
	BookId     int           `json:"book_id" xorm:"not null INT(11)"`
	UserId     int           `json:"user_id" xorm:"not null INT(11)"`
	BorrowTime myTime.MyTime `json:"borrow_time" xorm:"not null comment('借书时间') DATETIME"`
	ReturnTime myTime.MyTime `json:"return_time" xorm:"not null comment('还书时间') DATETIME"`
	Status     int           `json:"status" xorm:"not null comment('1借出 2归还') TINYINT(4)"`
	AddTime    myTime.MyTime `json:"add_time" xorm:"not null created DATETIME"`
	UpdateTime myTime.MyTime `json:"update_time" xorm:"not null updated DATETIME"`
	DelTime    myTime.MyTime `json:"del_time" xorm:"not null DATETIME"`
	Del        int           `json:"del" xorm:"not null TINYINT(4)"`
}
