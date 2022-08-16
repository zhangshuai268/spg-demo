package service

import (
	"context"
	"errors"
	wechat_login "github.com/zhangshuai268/spg-go-pkg/service/wechat/login"
	"spg-demo/internal/api/spg_user/auth"
	"spg-demo/internal/api/store"
	"spg-demo/internal/config"
	"spg-demo/internal/model"
	"spg-demo/internal/pkg/middle"
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
	user, flag, _ := u.factory.User().Get(ctx, &model.SpgUser{
		OpenId: oauth.OpenId,
	})
	var userId int
	if !flag {
		userId, err = u.factory.User().Create(ctx, &model.SpgUser{
			OpenId:     oauth.OpenId,
			SessionKey: oauth.AccessToken,
		})
	} else {
		userId = user.Id
	}
	token, err := middle.GetToken(config.Conf.Apiuser.Api_secret, &auth.UserClaim{
		UserId: userId,
	})
	if err != nil {
		return "", err
	}
	return token, nil
}

func NewUserService(s *service) UserService {
	return &userService{
		factory: s.factory,
	}
}
