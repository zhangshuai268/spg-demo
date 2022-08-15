package main

import (
	"github.com/zhangshuai268/spg-go-pkg/pkg/logger"
	"spg-demo/internal/api/spg_admin"
	"spg-demo/internal/api/store"
	"spg-demo/internal/config"
)

func main() {
	//初始化日志
	_, err := logger.InitLogger(false)
	if err != nil {
		panic("日志初始化失败" + err.Error())
	}
	//初始化代码配置
	_, err = config.InitConfig()
	if err != nil {
		panic("配置初始化失败" + err.Error())
	}
	port := config.Conf.Apiadmin.Api_port
	//初始化数据层
	factory, err := store.GetFactory()
	if err != nil {
		panic("数据层初始化失败" + err.Error())
	}
	//初始化路由
	router, err := spg_admin.RouterInit(factory)
	err = router.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
