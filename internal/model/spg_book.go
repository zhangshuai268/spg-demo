package model

import (
	"github.com/zhangshuai268/spg-go-pkg/pkg/myTime"
)

type SpgBook struct {
	Id         int           `json:"id" xorm:"not null pk autoincr INT(11)"`
	Title      string        `json:"title" xorm:"not null comment('书记名称') VARCHAR(100)"`
	Image      string        `json:"image" xorm:"not null comment('书封面') VARCHAR(500)"`
	Writer     string        `json:"writer" xorm:"not null comment('作者') VARCHAR(100)"`
	Press      string        `json:"press" xorm:"not null comment('出版社') VARCHAR(100)"`
	AddTime    myTime.MyTime `json:"add_time" xorm:"not null created DATETIME"`
	UpdateTime myTime.MyTime `json:"update_time" xorm:"not null updated DATETIME"`
	DelTime    myTime.MyTime `json:"del_time" xorm:"not null DATETIME"`
	Del        int           `json:"del" xorm:"not null TINYINT(4)"`
}
