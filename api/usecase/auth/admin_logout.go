/*
ログアウト用のユースケース
*/
package auth

import (
	"net/http"

	response "github.com/Tomoya185-miyawaki/attend-log-gin/response/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AdminLogoutUseCase struct{}

func (adminLogoutUseCase *AdminLogoutUseCase) Exec(
	c *gin.Context,
) *response.LogoutResponse {
	session := sessions.Default(c)
	session.Clear()
	session.Options(sessions.Options{Path: "/", MaxAge: -1})
	session.Save()
	return &response.LogoutResponse{StatusCode: http.StatusOK, Message: "ログアウトに成功しました"}
}
