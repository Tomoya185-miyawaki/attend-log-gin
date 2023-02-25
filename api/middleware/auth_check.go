/*
認証チェック用のミドルウェア
*/
package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/Tomoya185-miyawaki/attend-log-gin/helper"
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
	response "github.com/Tomoya185-miyawaki/attend-log-gin/response/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	dproxy "github.com/koron/go-dproxy"
)

func AdminLoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		adminLoginUserJson, err := dproxy.New(session.Get("loginAdminUser")).String()
		unauthorizedMessage := "権限のないユーザーです"

		if err != nil {
			response := &response.AuthCheckResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    unauthorizedMessage,
			}
			helper.Response(c, nil, response, response.StatusCode)
			c.Abort()
		} else {
			var adminLoginUser dto.Admin
			err := json.Unmarshal([]byte(adminLoginUserJson), &adminLoginUser)
			response := &response.AuthCheckResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    unauthorizedMessage,
			}
			if err != nil {
				helper.Response(c, nil, response, response.StatusCode)
				c.Abort()
			} else {
				c.Next()
			}
		}
	}
}
