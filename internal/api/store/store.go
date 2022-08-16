package store

import (
	"github.com/zhangshuai268/spg-go-pkg/pkg/gorm"
	"github.com/zhangshuai268/spg-go-pkg/pkg/mongo"
	"github.com/zhangshuai268/spg-go-pkg/pkg/redis"
	"github.com/zhangshuai268/spg-go-pkg/pkg/xorm"
	"spg-demo/internal/config"
)

type datastore struct {
	xorm  *xorm.Engine
	redis *redis.Engine
	mongo *mongo.Engine
	gorm  *gorm.Engine
}

func (d *datastore) User() UserStore {
	return NewUserStore(d)
}

func (d *datastore) Code() CodeStore {
	return NewCodeStore(d)
}

func (d *datastore) Admin() AdminStore {
	return NewAdminStore(d)
}

var DataStore Factory

func GetFactory() (Factory, error) {

	mysqlEngine, err := getMysqlEngine()
	if err != nil {
		return nil, err
	}

	redisEngine, err := getRedisEngine()
	if err != nil {
		return nil, err
	}

	mongoEngine, err := getMongoEngine()
	if err != nil {
		return nil, err
	}

	gormEngine, err := getGormEngine()
	if err != nil {
		return nil, err
	}
	DataStore = &datastore{
		xorm:  mysqlEngine,
		redis: redisEngine,
		mongo: mongoEngine,
		gorm:  gormEngine,
	}

	return DataStore, nil
}

func getMongoEngine() (*mongo.Engine, error) {
	mongoC := config.Conf.Mongodb
	var uri string
	if mongoC.User != "" {
		uri = mongoC.Driver + "://" + mongoC.User + ":" + mongoC.Pass_word + "@" + mongoC.Host
	} else {
		uri = mongoC.Driver + "://" + mongoC.Host
	}
	mongoEngin, err := mongo.InitMongoEngin(uri, mongoC.Db_name)
	if err != nil {
		return nil, err
	}
	return mongoEngin, nil
}

func getMysqlEngine() (*xorm.Engine, error) {
	//获取配置信息
	mysql := config.Conf.Mysql
	xormEngine, err := xorm.InitXormEngine("mysql", mysql.User+":"+mysql.Pass_word+"@("+mysql.Host+":"+mysql.Port+")/"+mysql.Db_name+
		"?charset="+mysql.Charset+"&"+"&parseTime=true&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci")
	if err != nil {
		return nil, err
	}
	return xormEngine, nil
}

func getRedisEngine() (*redis.Engine, error) {
	redisC := config.Conf.Redis
	redisEngine, err := redis.InitRedisEngine(redisC.Host+":"+redisC.Port, redisC.Pass_word, int(redisC.Db))
	if err != nil {
		return nil, err
	}
	return redisEngine, nil
}

func getGormEngine() (*gorm.Engine, error) {
	//获取配置信息
	mysql := config.Conf.Mysql
	gormEngine, err := gorm.InitGormEngine(mysql.User + ":" + mysql.Pass_word + "@tcp(" + mysql.Host + ":" + mysql.Port + ")/" + mysql.Db_name + "?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	return gormEngine, nil
}
