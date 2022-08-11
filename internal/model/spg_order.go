package models

import "github.com/zhangshuai268/spg-go-pkg/pkg/myTime"

type SpgOrder struct {
	Id           int           `json:"id" xorm:"not null pk INT(11)"`
	UserId       int           `json:"user_id" xorm:"not null INT(11)"`
	OrderNo      string        `json:"order_no" xorm:"not null comment('自定义订单号') VARCHAR(100)"`
	OrderTradeNo string        `json:"order_trade_no" xorm:"not null comment('微信或支付宝订单号') VARCHAR(100)"`
	Price        float64       `json:"price" xorm:"not null comment('支付价格') DECIMAL(10,2)"`
	BorrowId     int           `json:"borrow_id" xorm:"not null comment('对应借书id') INT(11)"`
	PayTime      myTime.MyTime `json:"pay_time" xorm:"not null comment('支付时间') DATETIME"`
	AddTime      myTime.MyTime `json:"add_time" xorm:"not null created DATETIME"`
	UpdateTime   myTime.MyTime `json:"update_time" xorm:"not null updated DATETIME"`
	DelTime      myTime.MyTime `json:"del_time" xorm:"not null DATETIME"`
	Del          int           `json:"del" xorm:"not null TINYINT(4)"`
}
