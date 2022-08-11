package config

type Apiuser struct{
	Api_port string
	Api_secret string
}
type Mongodb struct{
	Driver string
	User string
	Pass_word string
	Host string
	Db_name string
}
type Wxpay struct{
	App_id string
	Mch_id string
	App_key string
	Notify_url string
}
type Alipay struct{
	App_id string
	Private_key string
	Public_key string
	Notify_url string
	Return_url string
	Port string
}
type Alisms struct{
	Template_code string
	Sign_name string
	Access_key_id string
	Access_key_secret string
	Region_id string
}
type Apiadmin struct{
	Api_secret string
	Api_port string
}
type Mysql struct{
	Pass_word string
	Port string
	Charset string
	Show_sql string
	Loc string
	Driver string
	User string
	Host string
	Db_name string
	Parsetime string
}
type Redis struct{
	Host string
	Port string
	Pass_word string
	Db float64
}
type Wxconfig struct{
	App_id string
	App_secret string
}
type Config struct{
	Apiuser Apiuser
	Mongodb Mongodb
	Wxpay Wxpay
	Alipay Alipay
	Alisms Alisms
	Apiadmin Apiadmin
	Mysql Mysql
	Redis Redis
	Wxconfig Wxconfig
}