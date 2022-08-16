package service

import (
	"github.com/zhangshuai268/spg-go-pkg/pkg/logger"
	"os"
	"spg-demo/internal/api/store"
	"spg-demo/internal/config"
	"testing"
)

var serviceTest Service

func TestMain(m *testing.M) {
	//初始化日志
	_, err := logger.InitLogger(false)
	if err != nil {
		panic("日志初始化失败" + err.Error())
	}
	//初始化代码配置
	_, err = config.InitConfig("./../../../../config_local.json")
	if err != nil {
		panic("配置初始化失败" + err.Error())
	}
	//初始化数据层
	factory, err := store.GetFactory()
	if err != nil {
		panic("数据层初始化失败" + err.Error())
	}
	serviceTest = NewService(factory)
	os.Exit(m.Run())
}
