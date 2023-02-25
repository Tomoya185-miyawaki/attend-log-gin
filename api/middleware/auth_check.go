/*
認証チェック用のミドルウェア
*/
package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/Tomoya185-miyawaki/attend-log-gin/helper"
	"github.com/Tomoya185-miyawaki/attend-log-gin/infrastructure/dto"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	dproxy "github.com/koron/go-dproxy"
)

func AdminLoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		adminLoginUserJson, err := dproxy.New(session.Get("loginAdminUser")).String()

		if err != nil {
			helper.Response(c, nil, "権限のないユーザーです", http.StatusUnauthorized)
			c.Abort()
		} else {
			var adminLoginUser dto.Admin
			err := json.Unmarshal([]byte(adminLoginUserJson), &adminLoginUser)
			if err != nil {
				helper.Response(c, nil, "権限のないユーザーです", http.StatusUnauthorized)
				c.Abort()
			} else {
				c.Next()
			}
		}
	}
}
