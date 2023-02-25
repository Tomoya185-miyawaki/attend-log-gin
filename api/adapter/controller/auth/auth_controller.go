/*
認証用コントローラ
*/
package auth

import (
	"net/http"

	"github.com/Tomoya185-miyawaki/attend-log-gin/helper"
	"github.com/Tomoya185-miyawaki/attend-log-gin/request/auth"
	usecase "github.com/Tomoya185-miyawaki/attend-log-gin/usecase/auth"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (ac *AuthController) Login(c *gin.Context) {
	var request auth.LoginRequest
	var adminLoginUseCase usecase.AdminLoginUseCase
	response := adminLoginUseCase.Exec(c, &request)
	if response.StatusCode != http.StatusOK {
		helper.Response(c, nil, response, response.StatusCode)
	} else {
		helper.Response(c, response, nil, response.StatusCode)
	}
}

func (ac *AuthController) Logout(c *gin.Context) {
	var adminLogoutUseCase usecase.AdminLogoutUseCase
	response := adminLogoutUseCase.Exec(c)
	helper.Response(c, response, nil, response.StatusCode)
}
