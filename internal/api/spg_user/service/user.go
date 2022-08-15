package service

import (
	"context"
	"errors"
	wechat_login "github.com/zhangshuai268/spg-go-pkg/service/wechat/login"
	"spg-demo/internal/api/store"
	"spg-demo/internal/config"
)

type UserService interface {
	UserWxLogin(ctx context.Context, code string) (string, error)
}

type userService struct {
	factory store.Factory
}

func (u *userService) UserWxLogin(ctx context.Context, code string) (string, error) {
	client, err := wechat_login.NewLoginClient(config.Conf.Wxconfig.App_id, config.Conf.Wxconfig.App_secret)
	if err != nil {
		return "", err
	}
	oauth, err := client.Oauth(code)
	if oauth.ErrCode != 0 {
		return "", errors.New(oauth.ErrMessage)
	}
	return "", nil
}

func NewUserService(s *service) UserService {
	return &userService{
		factory: s.factory,
	}
}
