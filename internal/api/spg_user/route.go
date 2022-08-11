package spg_user

import (
	"github.com/gin-gonic/gin"
	"spg-demo/internal/api/store"
	"spg-demo/internal/pkg/middle"
)

func RouterInit(factory store.Factory) (*gin.Engine, error) {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//全局异常监控
	router.Use(middle.ErrHandler())
	router.Use(middle.CORS())

	store.SetClient(factory)
	return router, nil
}