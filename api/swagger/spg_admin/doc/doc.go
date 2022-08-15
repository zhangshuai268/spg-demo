//Package docs spg_admin
//
// spg_admin接口文档
// Schemes: http, https
// Version: 1.0.0
// BasePath: /
// Host: localhost:8081
// Consumes:
// - application/x-www-form-urlencoded
// Produces:
// - application/json
// Security:
// - api_key:
// SecurityDefinitions:
//  api_key:
//   type: apiKey
//   in: header
//   name: Authorization
// swagger:meta
package docs

//swagger:response errorResp
type errorResp struct {
	// in: body
	Body struct {
		//default: 0
		Status int `json:"status"`
		//default: error
		Message string `json:"message"`
		Data    string `json:"data"`
	}
}
