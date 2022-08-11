package models

type SpgUser struct {
	Id         int           `json:"id" xorm:"not null pk INT(11)"`
	Avatar     string        `json:"avatar" xorm:"not null comment('头像') VARCHAR(100)"`
	Title      string        `json:"title" xorm:"not null comment('用户名') VARCHAR(100)"`
	Mobile     string        `json:"mobile" xorm:"not null comment('手机号') VARCHAR(100)"`
	OpenId     string        `json:"open_id" xorm:"not null VARCHAR(100)"`
	SessionKey string        `json:"session_key" xorm:"not null VARCHAR(1000)"`
	AddTime    myTime.MyTime `json:"add_time" xorm:"not null created DATETIME"`
	UpdateTime myTime.MyTime `json:"update_time" xorm:"not null updated DATETIME"`
	DelTime    myTime.MyTime `json:"del_time" xorm:"not null DATETIME"`
	Del        int           `json:"del" xorm:"not null TINYINT(4)"`
}
