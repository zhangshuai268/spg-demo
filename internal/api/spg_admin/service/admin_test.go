package service

import (
	"context"
	"github.com/zhangshuai268/spg-go-pkg/pkg/logger"
	"testing"
)

func TestAdmins_AdminSendCode(t *testing.T) {
	mobile := "17685752864"
	err := serviceTest.Admin().AdminSendCode(context.TODO(), mobile)
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	return
}

func TestAdminService_AdminLogin(t *testing.T) {
	mobile := "17685752864"
	code := "346500"
	token, err := serviceTest.Admin().AdminLogin(context.TODO(), mobile, code)
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	logger.Logger.Info(token)
	return
}
