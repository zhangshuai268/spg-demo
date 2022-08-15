package auth

import (
	"bytes"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/zhangshuai268/spg-go-pkg/pkg/logger"
	"io/ioutil"
	"spg-demo/internal/api/store"
	"spg-demo/internal/config"
	"spg-demo/internal/model"
	"spg-demo/internal/pkg/code"
	"spg-demo/internal/pkg/middle"
	"strings"
)

type AdminClaim struct {
	AdminId int `json:"user_id"`
	jwt.StandardClaims
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取请求方法
		reqMethod := ctx.Request.Method
		//获取请求路由
		reqUri := ctx.Request.RequestURI
		// 请求IP
		clientIP := ctx.ClientIP()
		// 状态码
		statusCode := ctx.Writer.Status()
		//获取请求参数
		data, err := ctx.GetRawData()
		if err != nil {
			logger.Logger.Error(err.Error())
			code.BuildReturn(ctx, 0, err.Error(), "")
			ctx.Abort()
			return
		}
		//很关键
		//把读过的字节流重新放到body
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		//生成接口调用日志
		if logger.Logger.OnFile {
			err = logger.Logger.SetApiFile(reqUri)
			if err != nil {
				logger.Logger.Error(err.Error())
				code.BuildReturn(ctx, 0, err.Error(), "")
				ctx.Abort()
				return
			}
		}
		body := strings.ReplaceAll(string(data), "&", " ")
		// 日志格式
		logger.Logger.Api(
			statusCode,
			clientIP,
			reqMethod,
			reqUri,
			body,
		)
		//解析token
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code.BuildReturn(ctx, 0, "未登录", "")
			ctx.Abort()
			return
		}
		auth := &AdminClaim{}
		err = middle.ParseToken(token, config.Conf.Apiadmin.Api_secret, auth)
		if err != nil {
			logger.Logger.Error(err)
			code.BuildReturn(ctx, 0, err.Error(), "")
			ctx.Abort()
			return
		}
		adminId := auth.AdminId
		admin := &model.SpgAdmin{
			Id: adminId,
		}
		get, flag, err := store.Client().Admin().Get(ctx, admin)
		if err != nil {
			logger.Logger.Error(err)
			code.BuildReturn(ctx, 0, err.Error(), "")
			ctx.Abort()
			return
		}
		if !flag {
			code.BuildReturn(ctx, 0, "解析失败", "")
			ctx.Abort()
			return
		}
		ctx.Set("admin_id", get.Id)
		ctx.Next()
	}
}
