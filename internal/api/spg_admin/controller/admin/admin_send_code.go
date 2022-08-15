package admin

import (
	"github.com/gin-gonic/gin"
	docs "spg-demo/api/swagger/spg_admin/doc"
	"spg-demo/internal/pkg/code"
)

func (c *ControllerAdmin) AdminSendCode(ctx *gin.Context) {
	var request docs.SendCodeRequest
	err := ctx.Bind(&request)
	if err != nil {
		code.BuildReturn(ctx, 0, err.Error(), "")
		return
	}
	err = c.srv.Admin().AdminSendCode(ctx, request.Mobile)
	if err != nil {
		code.BuildReturn(ctx, 0, err.Error(), "")
		return
	}
}
