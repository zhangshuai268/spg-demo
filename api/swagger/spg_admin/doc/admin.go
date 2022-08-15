package docs

//TODO: 短信发送
//swagger:route POST /spg_admin/admin/send_code 后台-短信发送 sendCode
//短信发送
//Security:
// api_key:
//Responses:
//  200: adminSendCodeResponse
//  default: errorResp

//swagger:parameters sendCode
type SendCodeRequest struct {
	//手机号
	//required: true
	//in: formData
	Mobile string `json:"mobile" form:"mobile" binding:"required"`
}

//swagger:response adminSendCodeResponse
type adminSendCodeResponse struct {
	//in: body
	Body struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    string `json:"data"`
	}
}
