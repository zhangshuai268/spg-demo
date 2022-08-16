package service

import (
	"context"
	"github.com/zhangshuai268/spg-go-pkg/pkg/logger"
	"testing"
)

func TestUserService_UserWxLogin(t *testing.T) {
	code := "sdasdasdasd"
	login, err := serviceTest.User().UserWxLogin(context.TODO(), code)
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	logger.Logger.Info(login)
	return
}
