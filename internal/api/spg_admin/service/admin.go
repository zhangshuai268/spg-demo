package service

import (
	"context"
	"errors"
	"github.com/zhangshuai268/spg-go-pkg/pkg/util"
	aliyun_sms "github.com/zhangshuai268/spg-go-pkg/service/aliyun/sms"
	"spg-demo/internal/api/spg_admin/auth"
	"spg-demo/internal/api/store"
	"spg-demo/internal/config"
	"spg-demo/internal/model"
	"spg-demo/internal/pkg/middle"
)

type AdminService interface {
	AdminSendCode(ctx context.Context, mobile string) error
	AdminLogin(ctx context.Context, mobile, code string) (string, error)
}

type adminService struct {
	factory store.Factory
}

func (a *adminService) AdminLogin(ctx context.Context, mobile, code string) (string, error) {
	find, _, err := a.factory.Code().Find(ctx, map[string]interface{}{
		"mobile": mobile,
	}, "add_time")
	if err != nil {
		return "", err
	}
	if len(find) == 0 {
		return "", errors.New("请发送验证码")
	}
	if code != find[0].Code {
		return "", errors.New("验证码错误")
	}
	//获取管理员信息
	get, flag, _ := a.factory.Admin().Get(ctx, &model.SpgAdmin{
		Mobile: mobile,
	})
	if !flag {
		return "", errors.New("无此管理员")
	}
	token, err := middle.GetToken(config.Conf.Apiadmin.Api_secret, &auth.AdminClaim{
		AdminId: get.Id,
	})
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *adminService) AdminSendCode(ctx context.Context, mobile string) error {
	client, err := aliyun_sms.NewSmsClient(config.Conf.Alisms.Access_key_id, config.Conf.Alisms.Access_key_secret, config.Conf.Alisms.Region_id, config.Conf.Alisms.Template_code, config.Conf.Alisms.Sign_name)
	if err != nil {
		return err
	}
	code := util.GetRandNum(6, util.NumChar)
	content := `{"code":"` + code + `"}`
	send, err := client.SmsSend(mobile, content)
	if err != nil {
		return err
	}
	if send.Code != "OK" {
		return errors.New(send.Message)
	}
	_, err = a.factory.Code().Create(ctx, &model.SpgCode{
		Code:   code,
		Mobile: mobile,
	})
	if err != nil {
		return err
	}
	return nil
}

func NewAdminService(s *service) AdminService {
	return &adminService{
		factory: s.factory,
	}
}
